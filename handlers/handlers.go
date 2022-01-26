package handlers

import (
	"fmt"
	"net/http"

	database "TodoGolang/database"
	"TodoGolang/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *APIEnv) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, exists, err := database.GetUserByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *APIEnv) GetUsers(c *gin.Context) {
	users, err := database.GetUsers(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)

}

func (a *APIEnv) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetUserByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	err = database.DeleteUser(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}

func (a *APIEnv) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetUserByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	updatedUser := models.User{}
	err = c.BindJSON(&updatedUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.UpdateUser(a.DB, &updatedUser); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.GetUser(c)
}
