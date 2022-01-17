package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"tmi-gin/common/obj"
	"tmi-gin/common/response"
	"tmi-gin/dto"
	"tmi-gin/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MaidHandler interface {
	All(ctx *gin.Context)
	CreateMaid(ctx *gin.Context)
	UpdateMaid(ctx *gin.Context)
	DeleteMaid(ctx *gin.Context)
	FindOneMaidByID(ctx *gin.Context)
}

type maidHandler struct {
	maidService service.MaidService
	jwtService  service.JWTService
}

func NewMaidHandler(maidService service.MaidService, jwtService service.JWTService) MaidHandler {
	return &maidHandler{
		maidService: maidService,
		jwtService:  jwtService,
	}
}

func (c *maidHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	maids, err := c.maidService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", maids)
	ctx.JSON(http.StatusOK, response)
}

func (c *maidHandler) CreateMaid(ctx *gin.Context) {
	var createMaidReq dto.CreateMaidRequest
	err := ctx.ShouldBind(&createMaidReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.maidService.CreateMaid(createMaidReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *maidHandler) FindOneMaidByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.maidService.FindOneMaidByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *maidHandler) DeleteMaid(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.maidService.DeleteMaid(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *maidHandler) UpdateMaid(ctx *gin.Context) {
	updateMaidRequest := dto.UpdateMaidRequest{}
	err := ctx.ShouldBind(&updateMaidRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateMaidRequest.ID = id
	maid, err := c.maidService.UpdateMaid(updateMaidRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", maid)
	ctx.JSON(http.StatusOK, response)

}
