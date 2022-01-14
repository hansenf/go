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
	db               *gorm.DB                 = config.SetupDatabaseConnection()
	userRepo         repo.UserRepository      = repo.NewUserRepo(db)
	mahasiswaRepo    repo.MahasiswaRepository = repo.NewMahasiswaRepo(db)
	authService      service.AuthService      = service.NewAuthService(userRepo)
	jwtService       service.JWTService       = service.NewJWTService()
	userService      service.UserService      = service.NewUserService(userRepo)
	mahasiswaService service.MahasiswaService = service.NewMahasiswaService(mahasiswaRepo)
	authHandler      v1.AuthHandler           = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler      v1.UserHandler           = v1.NewUserHandler(userService, jwtService)
	mahasiswaHandler v1.MahasiswaHandler      = v1.NewMahasiswaHandler(mahasiswaService, jwtService)
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

	mahasiswaRoutes := server.Group("api/mahasiswa", middleware.AuthorizeJWT(jwtService))
	{
		mahasiswaRoutes.GET("/", mahasiswaHandler.All)
		mahasiswaRoutes.POST("/", mahasiswaHandler.CreateMahasiswa)
		mahasiswaRoutes.GET("/:id", mahasiswaHandler.FindOneMahasiswaByID)
		mahasiswaRoutes.PUT("/:id", mahasiswaHandler.UpdateMahasiswa)
		mahasiswaRoutes.DELETE("/:id", mahasiswaHandler.DeleteMahasiswa)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
