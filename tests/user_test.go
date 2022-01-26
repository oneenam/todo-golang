package tests

import (
	database "TodoGolang/database"
	"TodoGolang/handlers"
	"TodoGolang/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setGetUsersRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()

	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	r.GET("/user/all", api.GetUsers)
	req, err := http.NewRequest(http.MethodGet, "/user/all", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}

func insertTestUser(db *gorm.DB) (models.User, error) {
	b := models.User{
		Email:     "email@test.com",
		FirstName: "First name",
		LastName:  "Last name",
	}

	if err := db.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func setGetUserRouter(db *gorm.DB, id string) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	r.GET("/user/:id", api.GetUser)

	req, err := http.NewRequest(http.MethodGet, "/user/"+id, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}

func setCreateUserRouter(db *gorm.DB,
	body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	r.POST("/user/create", api.CreateUser)
	req, err := http.NewRequest(http.MethodPost, "/user/create", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

//EMPTY
func Test_GetUsers_EmptyResult(t *testing.T) {
	database.Setup()
	db := database.GetDB()
	req, w := setGetUsersRouter(db)
	defer db.Close()

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.User{}
	a.Equal(expected, actual)
	database.ClearTable()
}

//CREATE
func Test_CreateUser_OK(t *testing.T) {
	a := assert.New(t)
	database.Setup()
	db := database.GetDB()
	user := models.User{
		Email:     "email@test.com",
		FirstName: "First name",
		LastName:  "Last name",
	}

	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateUserRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := user
	a.Equal(expected, actual)
	database.ClearTable()
}

//GET
func Test_GetUser_OK(t *testing.T) {
	a := assert.New(t)
	database.Setup()
	db := database.GetDB()

	user, err := insertTestUser(db)
	if err != nil {
		a.Error(err)
	}

	req, w := setGetUserRouter(db, strconv.FormatUint(uint64(user.ID), 10))
	defer db.Close()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.User{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := user
	expected.Model = gorm.Model{}
	a.Equal(expected, actual)
	database.ClearTable()
}
