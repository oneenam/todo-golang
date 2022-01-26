package handlers

import (
	"fmt"
	"net/http"

	"../models"

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
