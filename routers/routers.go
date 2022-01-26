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

	//r.GET("", api.CreateUser)
	//r.GET("/:id", api.GetBook)
	r.POST("", api.CreateUser)
	//r.PUT("/:id", api.UpdateBook)
	//r.DELETE("/:id", api.DeleteBook)

	return r
}
