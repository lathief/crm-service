package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var (
	host       = "localhost"   //os.Getenv("DB_HOST")
	user       = "root"        //os.Getenv("DB_USER")
	password   = ""            //os.Getenv("DB_PASSWORD")
	dbPort     = "3306"        //os.Getenv("DB_PORT")
	dbName     = "crm_service" //os.Getenv("DB_NAME")
	DEBUG_MODE = true
	db         *gorm.DB
	err        error
)

func StartDB() (*gorm.DB, error) {
	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?%s`, user, password, host, dbPort, dbName, val.Encode())
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Connection database success")
	return db, nil
}
