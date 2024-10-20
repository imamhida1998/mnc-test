package controller

import (
	"github.com/gin-gonic/gin"
	"mnc-test/model"
	"mnc-test/model/request"
	"mnc-test/service/usecase"
	"net/http"
)

type userController struct {
	userService usecase.UserUsecase
	authService usecase.AuthService
}

func NewUserController(userService usecase.UserUsecase, authService usecase.AuthService) *userController {
	return &userController{userService, authService}
}

func (c *userController) Register(ctx *gin.Context) {
	var input request.Register

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	resUser, err := c.userService.Register(input)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := gin.H{
		"status": "SUCCESS",
		"result": resUser,
	}

	ctx.JSON(http.StatusCreated, response)
	return

}

func (c *userController) Login(ctx *gin.Context) {
	var input request.Login

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	resUser, err := c.userService.GetAcoountByPhoneNumber(input)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	accessToken, err := c.authService.AccessToken(resUser.PhoneNumber)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"access_token":  accessToken,
			"refresh_token": accessToken,
		},
	}

	ctx.JSON(http.StatusCreated, response)
	return

}

func (c *userController) UpdateProfile(ctx *gin.Context) {
	var input model.User

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, found := ctx.MustGet("currentUser").(*model.User)
	if !found {
		response := gin.H{
			"status": "FAILED",
			"result": "user not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	resUser, err := c.userService.UpdateProfiles(&input, users)
	if err != nil {
		response := gin.H{
			"status": "FAILED",
			"result": err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := gin.H{
		"status": "SUCCESS",
		"result": resUser,
	}

	ctx.JSON(http.StatusCreated, response)
	return

}
