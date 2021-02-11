package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haru0017/go-clean-architecture/infrastructure/db"
	"github.com/haru0017/go-clean-architecture/interface/controllers"
)

// Router はrouterを外部から使うために使用
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewController(db.NewSQLHandler())

	router.POST("/users", func(c *gin.Context) { userController.CreateUser(c) })
	router.GET("/users", func(c *gin.Context) { userController.GetUsers(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.GetUser(c) })
	router.PUT("/users/:id", func(c *gin.Context) { userController.UpdateUser(c) })
	router.DELETE("/users/:id", func(c *gin.Context) { userController.DeleteUser(c) })

	Router = router
}
