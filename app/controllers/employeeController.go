package controllers

import (
	"github.com/revel/revel"
	"github.com/marwaan/testproject/app/interceptors"
	"github.com/marwaan/testproject/app/encoders"
	"github.com/marwaan/testproject/app/util"
	"strconv"
	"log"
	"github.com/marwaan/testproject/app"
	"github.com/marwaan/testproject/app/models"
)

type EmployeeController struct {

	interceptors.JWTAuthorization
	*revel.Controller
}
func (c EmployeeController) Find()revel.Result  {
	var employee models.Employee
	var id int
	c.Params.Bind(&id, "id")

	if founded := app.Db.Find(&employee, id).RowsAffected; founded < 1{
		return c.RenderJson(util.ResponseError("not founded"))
	}
	app.Db.First(&employee.User , employee.UserID)
	employee.User.Password = ""
	return c.RenderJson(employee);
}

func (c EmployeeController) Index()revel.Result  {
	var employees []models.Employee
	var limitQuery =c.Request.URL.Query().Get("limit")

	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset")

	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&employees).RowsAffected; founded < 1{
		return c.RenderJson(util.ResponseError("not founded"))
	}
	for i, employee := range employees {
		app.Db.First(&employees[i].User , employee.UserID)
		employees[i].User.Password = ""

	}
	return c.RenderJson(employees);
}

func (c EmployeeController) Create()revel.Result  {
	var employee = encoders.EncodeEmployee(c.Request.Body);
	employee.UserID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	if employee.First_name == ""||employee.Last_name == ""{
		return c.RenderJson(util.ResponseError("Employee Information Not Founded"))
	}

	if err := app.Db.Create(&employee).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Employee Registration Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(employee));
	//
	//var employee = encoders.EncodeEmployee(c.Request.Body);
	//if employee.First_name == ""||employee.Last_name == "" {
	//	return c.RenderJson(util.ResponseError("Employee Information Not Founded"))
	//}
	////var id int
	//employee.UserID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	//c.Params.Bind(&employee.UserID, "id")
	////log.Println("===============", id)
	//if err := app.Db.Create(&employee).Error; err != nil {
	//	log.Println(err);
	//	return c.RenderJson(util.ResponseError("Employee Creation Fialed"))
	//}
	//return c.RenderJson(util.ResponseSuccess(employee));
}

func (c EmployeeController) Update() revel.Result  {
	var update = encoders.EncodeEmployee(c.Request.Body);
	var id int
	var employee models.Employee
	c.Params.Bind(&id, "id")

	if rowcount := app.Db.First(&employee, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("Employee Information Not Founded"))
	}
	if err := app.Db.Model(&employee).Update(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Employee Update Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update));
}

func (c EmployeeController) Delete() revel.Result  {
	var (
		id int
		employee models.Employee
	)
	c.Params.Bind(&id, "id")
	log.Println("++++++++++++++++++++",id)
	if rowcount := app.Db.First(&employee, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("Employee Information Not Founded"))
	}

	if err := app.Db.Delete(&employee).Error; err != nil{
		return c.RenderJson(util.ResponseError("Employee Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(employee));
}

