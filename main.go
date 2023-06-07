package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/config"
	"github.com/lathief/crm-service/middleware"
	"github.com/lathief/crm-service/modules/actor"
	"github.com/lathief/crm-service/modules/customer"
	"github.com/lathief/crm-service/utils/database"
	"log"
)

func init() {
	config.SetupConfiguration()
}

func main() {
	router := gin.New()
	db, err := database.StartDB()
	if err != nil {
		log.Fatal(err)
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
	validation := middleware.NewValidation()
	jsonAuth := middleware.NewSecurity()
	customerHandler := customer.NewRouter(db, jsonAuth, validation)
	customerHandler.Handle(router)
	actorHandler := actor.NewRouter(db, jsonAuth, validation)
	actorHandler.Handle(router)
	errRouter := router.Run(fmt.Sprintf(":%s", config.Config.Server.Port))
	if errRouter != nil {
		log.Fatal(errRouter)
	}
}
