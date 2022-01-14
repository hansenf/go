package v1

import (
	"fmt"
	"net/http"

	"tmi-gin/common/obj"
	"tmi-gin/common/response"
	"tmi-gin/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Profile(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserHandler(
	userService service.UserService,
	jwtService service.JWTService,
) UserHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userHandler) getUserIDByHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := response.BuildErrorResponse("Error", "Failed to validate token", obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *userHandler) Profile(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := response.BuildErrorResponse("Error", "Failed to validate token", obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.userService.FindUserByID(id)

	if err != nil {
		response := response.BuildErrorResponse("Error", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	res := response.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}
