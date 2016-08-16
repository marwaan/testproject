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

type RewardController struct {

	interceptors.JWTAuthorization
	*revel.Controller
}
func (c RewardController) Index()revel.Result  {
	var rewards []models.Reward
	var limitQuery =c.Request.URL.Query().Get("limit")

	if limitQuery == "" {
		limitQuery = "0"
	}
	var offsetQuery =c.Request.URL.Query().Get("offset")
	var id int
	c.Params.Bind(&id, "id")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&rewards).RowsAffected; founded < 1{
		return c.RenderJson(util.ResponseError("not founded"))
	}

	return c.RenderJson(rewards);
}
func (c RewardController) Create()revel.Result  {
	var reward = encoders.EncodeReward(c.Request.Body);
	if reward.Name == ""{
		return c.RenderJson(util.ResponseError("Reward Information Not Founded"))
	}
	//reward.UserID,_ = strconv.ParseInt(c.Session["id"], 10, 0)
	//c.Params.Bind(&reward.EmployeeID, "id")
	if err := app.Db.Create(&reward).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("Reward Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(reward));
}
func (c RewardController) Update() revel.Result  {
	var update = encoders.EncodeReward(c.Request.Body);

	var reward models.Reward
	var id int
	c.Params.Bind(&id, "id")

	if rowcount := app.Db.First(&reward, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("Reward Information Not Founded"))
	}
	if err := app.Db.Model(&reward).Update(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("Reward Update Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update));
}
func (c RewardController) Delete() revel.Result  {
	var (
		id int
		Reward models.Reward
	)
	c.Params.Bind(&id, "id")
	if rowcount := app.Db.First(&Reward, id).RowsAffected; rowcount < 1{

		return c.RenderJson(util.ResponseError("Reward Information Not Founded"))
	}

	if err := app.Db.Delete(&Reward).Error; err != nil{
		return c.RenderJson(util.ResponseError("Reward Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(Reward));
}



