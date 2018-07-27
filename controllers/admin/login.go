package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	PassWord string `form:"pass_word" json:"pass_word" binding:"required"`
}
func (login *Login)LoginInfo(context *gin.Context)  {
	var json Login
	if err := context.Bind(&json); err == nil {
		if json.UserName == "test" && json.PassWord == "123" {
			context.JSON(http.StatusOK, gin.H{"status":"login success"})
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"status":"unauthorized"})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
