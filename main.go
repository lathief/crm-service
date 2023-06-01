package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/config"
	"github.com/lathief/crm-service/utils/database"
	"log"
)

func init() {
	config.SetupConfiguration()
}

func main() {
	router := gin.New()
	fmt.Println("Initial Project")
	db, err := database.StartDB()
	if err != nil {
		return
	}
	//check database
	checkdb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	//ping to database
	errconn := checkdb.Ping()
	if err != nil {
		log.Fatal(errconn)
	}
	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Printf(errRouter.Error())
		return
	}
}
