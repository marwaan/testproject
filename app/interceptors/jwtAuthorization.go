package interceptors

import (
	"github.com/revel/revel"
	"github.com/dgrijalva/jwt-go"
	"github.com/marwaan/testproject/app/util"
	"strconv"
	"fmt"
)

type JWTAuthorization  struct {

    *revel.Controller

}
func (c JWTAuthorization) checkUser() revel.Result {
	var tokenString = c.Request.Header.Get("token");

	token, err:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		appSecret, _:= revel.Config.String("app.secret")
		return []byte(appSecret), nil
	})
	if  err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Session["email"]= claims["email"].(string);
			c.Session["id"]= strconv.Itoa(int(claims["id"].(float64)));
			return nil
		}

	}else {
		return c.RenderJson(util.ResponseError("Invalid token Key"))
	}
	return c.RenderJson(util.ResponseSuccess(token));
}
func init() {
	revel.InterceptMethod(JWTAuthorization.checkUser, revel.BEFORE);
}