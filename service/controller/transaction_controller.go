package controller

import (
	"github.com/gin-gonic/gin"
	"mnc-test/model"
	"mnc-test/model/request"
	"mnc-test/service/usecase"
	"net/http"
)

type topupController struct {
	topupUsecase usecase.TransactionUsecase
}

func NewTopupController(topup usecase.TransactionUsecase) topupController {
	return topupController{topup}
}

func (t *topupController) TopUp(ctx *gin.Context) {
	var input request.TopUp

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

	resTopUp, err := t.topupUsecase.CreateTopUp(&input, users)
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
		"result": resTopUp,
	}

	ctx.JSON(http.StatusOK, response)

}

func (t *topupController) Payment(ctx *gin.Context) {
	var input request.Payment

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

	resTopUp, err := t.topupUsecase.CreatePayment(&input, users)
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
		"result": resTopUp,
	}

	ctx.JSON(http.StatusOK, response)

}

func (t *topupController) Transfer(ctx *gin.Context) {
	var input request.Transfer

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

	resTopUp, err := t.topupUsecase.Transfer(&input, users)
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
		"result": resTopUp,
	}

	ctx.JSON(http.StatusOK, response)

}

func (t *topupController) Transaction(ctx *gin.Context) {

	users, found := ctx.MustGet("currentUser").(*model.User)
	if !found {
		response := gin.H{
			"status": "FAILED",
			"result": "user not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	resTopUp, err := t.topupUsecase.TransactionReport(users.UserId)
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
		"result": resTopUp,
	}

	ctx.JSON(http.StatusOK, response)

}
