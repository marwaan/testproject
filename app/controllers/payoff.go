package controllers

import (
	"github.com/revel/revel"
	"github.com/marwaan/testproject/app/interceptors"
	"github.com/marwaan/testproject/app/encoders"
	"github.com/marwaan/testproject/app/util"
	//"strconv"
	"log"
	"github.com/marwaan/testproject/app"
	"github.com/marwaan/testproject/app/models"
)

type PayoffController struct {

	interceptors.JWTAuthorization
	*revel.Controller
}
func (c PayoffController) Index()revel.Result  {
	var payoffs []models.Payoff
	var limitQuery =c.Request.URL.Query().Get("limit")

	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset")
	var id int
	c.Params.Bind(&id, "id")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&payoffs).RowsAffected; founded < 1{
		return c.RenderJson(util.ResponseError("not founded"))
	}
	return c.RenderJson(payoffs);
}
func (c PayoffController) Create()revel.Result  {
	var payoff = encoders.EncodePayoff(c.Request.Body);
	if payoff.Emp_firstName == ""{
		return c.RenderJson(util.ResponseError("payoff Information Not Founded"))
	}
	//payoff.EmployeeID,_ = strconv.ParseInt(c.Session["id"], 10, 0)

	//c.Params.Bind(&payoff.EmployeeID, "id")
	log.Println("sessionid:",payoff);
	if err := app.Db.Create(&payoff).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("payoff Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(payoff));
}
func (c PayoffController) Update() revel.Result  {
	var update = encoders.EncodePayoff(c.Request.Body);

	var payoff models.Payoff
	var id int
	c.Params.Bind(&id, "id")

	if rowcount := app.Db.First(&payoff, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("payoff Information Not Founded"))
	}
	if err := app.Db.Model(&payoff).Update(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("payoff Update Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update));
}
func (c PayoffController) Delete() revel.Result  {
	var (
		id int
		payoff models.Payoff
	)
	c.Params.Bind(&id, "id")
	if rowcount := app.Db.First(&payoff, id).RowsAffected; rowcount < 1{

		return c.RenderJson(util.ResponseError("payoff Information Not Founded"))
	}

	if err := app.Db.Delete(&payoff).Error; err != nil{
		return c.RenderJson(util.ResponseError("payoff Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(payoff));
}



