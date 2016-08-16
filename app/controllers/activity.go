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

type ActivityController struct {

	interceptors.JWTAuthorization
	*revel.Controller
}
func (c ActivityController) Index()revel.Result  {
	var activities []models.Activity
	var limitQuery =c.Request.URL.Query().Get("limit")

	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset")
	var id int
	c.Params.Bind(&id, "id")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&activities).RowsAffected; founded < 1{
		return c.RenderJson(util.ResponseError("not founded"))
	}

	return c.RenderJson(activities);
}
func (c ActivityController) Create()revel.Result  {
	var Activity = encoders.EncodeActivity(c.Request.Body);
	if Activity.Discription == ""{
		return c.RenderJson(util.ResponseError("Activity Information Not Founded"))
	}

	if err := app.Db.Create(&Activity).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Activity Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(Activity));
}
func (c ActivityController) Update() revel.Result  {
	var update = encoders.EncodeActivity(c.Request.Body);

	var Activity models.Activity
	var id int
	c.Params.Bind(&id, "id")

	if rowcount := app.Db.First(&Activity, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("Activity Information Not Founded"))
	}
	if err := app.Db.Model(&Activity).Update(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Activity Update Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update));
}
func (c ActivityController) Delete() revel.Result  {
	var (
		id int
		Activity models.Activity
	)
	c.Params.Bind(&id, "id")
	if rowcount := app.Db.First(&Activity, id).RowsAffected; rowcount < 1{

		return c.RenderJson(util.ResponseError("Activity Information Not Founded"))
	}

	if err := app.Db.Delete(&Activity).Error; err != nil{
		return c.RenderJson(util.ResponseError("Activity Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(Activity));
}



