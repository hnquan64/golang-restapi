package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hnquan64/golang-restapi/config"
	"github.com/hnquan64/golang-restapi/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetUpDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnecttion(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)

	}
	r.Run()
}
