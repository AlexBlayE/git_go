package errors

import "github.com/gin-gonic/gin"

var NeedAuthHeaderBody gin.H = gin.H{
	"error": "Need an bearer auth header",
}

var InvalidJWT gin.H = gin.H{
	"error": "invalid json web token",
}

var UserNotCreated gin.H = gin.H{
	"error": "user not created",
}

var CantGetUserTokn gin.H = gin.H{
	"error": "can't get token",
}
