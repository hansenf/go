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

type MahasiswaHandler interface {
	All(ctx *gin.Context)
	CreateMahasiswa(ctx *gin.Context)
	UpdateMahasiswa(ctx *gin.Context)
	DeleteMahasiswa(ctx *gin.Context)
	FindOneMahasiswaByID(ctx *gin.Context)
}

type mahasiswaHandler struct {
	mahasiswaService service.MahasiswaService
	jwtService       service.JWTService
}

func NewMahasiswaHandler(mahasiswaService service.MahasiswaService, jwtService service.JWTService) MahasiswaHandler {
	return &mahasiswaHandler{
		mahasiswaService: mahasiswaService,
		jwtService:       jwtService,
	}
}

func (c *mahasiswaHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	mahasiswas, err := c.mahasiswaService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", mahasiswas)
	ctx.JSON(http.StatusOK, response)
}

func (c *mahasiswaHandler) CreateMahasiswa(ctx *gin.Context) {
	var createMahasiswaReq dto.CreateMahasiswaRequest
	err := ctx.ShouldBind(&createMahasiswaReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.mahasiswaService.CreateMahasiswa(createMahasiswaReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *mahasiswaHandler) FindOneMahasiswaByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.mahasiswaService.FindOneMahasiswaByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *mahasiswaHandler) DeleteMahasiswa(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.mahasiswaService.DeleteMahasiswa(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *mahasiswaHandler) UpdateMahasiswa(ctx *gin.Context) {
	updateMahasiswaRequest := dto.UpdateMahasiswaRequest{}
	err := ctx.ShouldBind(&updateMahasiswaRequest)

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
	updateMahasiswaRequest.ID = id
	mahasiswa, err := c.mahasiswaService.UpdateMahasiswa(updateMahasiswaRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", mahasiswa)
	ctx.JSON(http.StatusOK, response)

}
