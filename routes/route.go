package routes

import (
	control "cbnit/controller"
	m "cbnit/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", control.Signup)
	router.POST("/login", control.Login)
	router.POST("/logout", m.TokenAuthMiddleware(), control.Logout)
	router.GET("/getallusers", m.TokenAuthMiddleware(), control.GetAllUsers)
	return router
}
