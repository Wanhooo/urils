package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urils/application/constants"
	. "urils/application/services"
)


/**
用户认证登陆
*/

func UserAuthenticate(ctx *gin.Context){
	user, err := UserLogin(ctx)

	if err != nil || user.ID < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeNoSuchUser,
			"message": constants.NoSuchUser+err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": constants.CodeSuccess,
		"message": constants.Success,
		"data": user,
	})
}


/**
 创建用户
*/

func UserCreate(ctx *gin.Context){

	id, err := CreateUser(ctx)

	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFail,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": constants.CodeSuccess,
		"message": constants.CreateUserSuccess,
		"data": id,
	})
}
