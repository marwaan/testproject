package controllers

import (
	"github.com/revel/revel"
	"log"
	"github.com/marwaan/testproject/app/util"
	"github.com/marwaan/testproject/app"
	"github.com/marwaan/testproject/app/encoders"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/marwaan/testproject/app/models"
)

type Usercontroller struct  {
	*revel.Controller

}

func (c Usercontroller) Create() revel.Result  {
	//var user encoders.EncodeSingleUsers(c.Request.Body)
	var user = encoders.EncodeSingleUsers(c.Request.Body);
	if user.Email == ""||user.Password == "" {
		return c.RenderJson(util.ResponseError("User information empty"))
	}
	if err := app.Db.Create(&user).Error; err != nil {
		log.Println(err);
		return c.RenderJson(util.ResponseError("User Creation Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(user));
}
func (c Usercontroller) Login() revel.Result  {
	var user = encoders.EncodeSingleUsers(c.Request.Body);

	if user.Email == ""||user.Password == "" {
		return c.RenderJson(util.ResponseError("User information empty"))
	}

	if founded := app.Db.Where(&user).First(&user).RowsAffected; founded < 1{
		return c.RenderJson(util.ResponseError("User not founnd"))
	}
	//return c.RenderJson(util.ResponseSuccess());

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":user.ID,
		"email":user.Email,
		"exp":time.Now().Add(time.Hour * 24),
	})
	appSecret, _:= revel.Config.String("app.secret")
	tokestring, err := token.SignedString([]byte(appSecret));

	if err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("key generation Fialed"))
	}
	var tokenmodel models.Token
	tokenmodel.Email = user.Email
	tokenmodel.Name = user.Name
	tokenmodel.Token = tokestring

	 return c.RenderJson(util.ResponseSuccess(tokenmodel));
	//return c.RenderJson(tokenmodel);
}
func (c Usercontroller) Update() revel.Result  {
	var (
		update = encoders.EncodeSingleUsers(c.Request.Body)
		 id int
		 user models.User
	)
	c.Params.Bind(&id, "id")

	if rowcount := app.Db.First(&user, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("User Information Not Founded"))
	}
	if err := app.Db.Model(&user).Update(&update).Error; err != nil{
		return c.RenderJson(util.ResponseError("User Update Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(update));
}
func (c Usercontroller) Delete() revel.Result  {
	var (
		id int
		user models.User
	)
	c.Params.Bind(&id, "id")
	if rowcount := app.Db.First(&user, id).RowsAffected; rowcount < 1{
		return c.RenderJson(util.ResponseError("User Information Not Founded"))
	}

	if err := app.Db.Delete(&user).Error; err != nil{
		return c.RenderJson(util.ResponseError("User Delete Fialed"))
	}
	return c.RenderJson(util.ResponseSuccess(user));
}

