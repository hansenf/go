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
	UpdataMahasiswa(ctx *gin.Context)
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

func (c *mahasiswaHandler) UpdataMahasiswa(ctx *gin.Context) {
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
	updateMahasiswaRequest.IDUser = int32(id)
	mahasiswa, err := c.mahasiswaService.UpdateMahasiswa(updateMahasiswaRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", mahasiswa)
	ctx.JSON(http.StatusOK, response)

}
