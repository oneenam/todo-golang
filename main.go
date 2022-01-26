package main

import (
	"TodoGolang/database"
	"TodoGolang/routers"
	"log"
)

func main() {
	database.Setup()
	r := routers.Setup()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

//test connection refused
//https://stackoverflow.com/questions/58222386/github-actions-using-mysql-service-throws-access-denied-for-user-rootlocalh
