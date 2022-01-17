package main

import (
	"tmi-gin/config"
	v1 "tmi-gin/handler/v1"
	"tmi-gin/middleware"
	"tmi-gin/repo"
	"tmi-gin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB            = config.SetupDatabaseConnection()
	userRepo    repo.UserRepository = repo.NewUserRepo(db)
	maidRepo    repo.MaidRepository = repo.NewMaidRepo(db)
	authService service.AuthService = service.NewAuthService(userRepo)
	jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepo)
	maidService service.MaidService = service.NewMaidService(maidRepo)
	authHandler v1.AuthHandler      = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler v1.UserHandler      = v1.NewUserHandler(userService, jwtService)
	maidHandler v1.MaidHandler      = v1.NewMaidHandler(maidService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
	}

	maidRoutes := server.Group("api/maid", middleware.AuthorizeJWT(jwtService))
	{
		maidRoutes.GET("/", maidHandler.All)
		maidRoutes.POST("/", maidHandler.CreateMaid)
		maidRoutes.GET("/:id", maidHandler.FindOneMaidByID)
		maidRoutes.PUT("/:id", maidHandler.UpdateMaid)
		maidRoutes.DELETE("/:id", maidHandler.DeleteMaid)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
