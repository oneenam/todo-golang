package database

import (
	"TodoGolang/models"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*DB is connected database object*/
var DB *gorm.DB

// SetupDB : initializing mysql database
func Setup() {

	/*USER := "root"
	PASS := "Enam:@7137"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "todo"*/

	//URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	//db, err := gorm.Open("mysql", URL)

	db, err := gorm.Open("mysql", "root:Enam:@7137@tcp(docker.for.mac.localhost:3306)/todo?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	/**
	If you set db.LogMode to true,
	you can see the SQL queries that you've written,
	whilst db.AutoMigrate create tables named "users"  & "tasks" from the struct
	*/
	db.LogMode(false)

	db.AutoMigrate([]models.User{})
	db.AutoMigrate([]models.Task{})

	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

func ClearTable() {
	DB.Exec("DELETE FROM users")
	DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}
