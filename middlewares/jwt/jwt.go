package jwt

import (
	"github.com/gin-gonic/gin"
	app "telrobot/util/common"
)

var jwtSecret = app.Env("OC_JWT_SECRET", "test")
var requestCountLimit = app.Env("OC_API_Request_Limit", "500")

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		//status := http.StatusOK
		//code := 0
		//
		//// read token from query
		//token := c.Query("token")
		//if token == "" {
		//	// read token from header
		//	auth := c.GetHeader("Authorization")
		//	if len(auth) > 0 {
		//		token = strings.TrimPrefix(auth, "Bearer ")
		//	}
		//}
		//
		//if token == "" {
		//	// read token from cookie
		//	cookie, err := c.Cookie("token")
		//	if err == nil && len(cookie) > 0 {
		//		token = cookie
		//	}
		//}
		//
		//if token == "" {
		//	status = http.StatusBadRequest
		//	code = e.ErrorAuthTokenRequired
		//} else {
		//	//claims, err := util.ParseToken(token, jwtSecret)
		//	//if err != nil {
		//	//	status = http.StatusUnauthorized
		//	//	code = e.ErrorAuthCheckTokenFail
		//	//} else if time.Now().Unix() > claims.ExpiresAt {
		//	//	status = http.StatusUnauthorized
		//	//	code = e.ErrorAuthCheckTokenTimeout
		//	//} else {
		//	//	c.Set("auth.id", claims.ID)
		//	//	c.Set("auth.username", claims.Username)
		//	//
		//	//	password := app.GetHashStringValue("userInfo", claims.Username)
		//	//	if claims.Password != password {
		//	//		status = http.StatusUnauthorized
		//	//		code = e.ErrorAuthCheckTokenFail
		//	//	}
		//	//}
		//}
		//
		//if status != http.StatusOK {
		//	//fmt.Println(code)
		//	c.JSON(status, e.Error(code))
		//
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}
