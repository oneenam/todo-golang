package routers

import (
	"TodoGolang/database"
	"TodoGolang/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	r.GET("/users/:id", api.GetUser)
	r.GET("/user/all", api.GetUsers)
	r.POST("/user/create", api.CreateUser)
	r.PUT("/user/update/:id", api.UpdateUser)
	r.DELETE("user/delete/:id", api.DeleteUser)

	//r.GET("/task/:id", api.GetTask)
	//r.GET("/task/all/:id", api.AllTasks)
	//r.GET("/task/create", api.CreateTask)
	//r.GET("/task/complete", api.CompleteTask)

	return r
}
