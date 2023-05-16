package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"urils/application/constants"
	. "urils/application/services"
	. "urils/application/utils"
)

/**
用户认证登陆
*/

func UserAuthenticate(ctx *gin.Context) {

	// 用户登陆认证业务
	user, err := UserLogin(ctx)

	if err != nil || user.ID < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeNoSuchUser,
			"message": constants.NoSuchUser + err.Error(),
		})
		return
	}

	// 生成token
	newJwt := NewJWT()
	publicClaims := PublicClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	token, err := newJwt.AccessToken(publicClaims)
	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

/**
创建用户
*/

func UserCreate(ctx *gin.Context) {

	id, err := CreateUser(ctx)

	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFail,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.CreateUserSuccess,
		"data":    id,
	})
}
