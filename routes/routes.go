package routes

import (
	"github.com/MikeFors0/ecommerce-go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.AddProductAdmin())
	incomingRoutes.POST("/users/viewproduct", controllers.ViewProduct())
	incomingRoutes.POST("/users/search", controllers.Search())
}