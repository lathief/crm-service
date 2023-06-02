package database

import (
	"fmt"
	"github.com/lathief/crm-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

func StartDB() (*gorm.DB, error) {
	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	//dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?%s`, user, password, host, dbPort, dbName, val.Encode())
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?%s`, config.Config.Database.DBUser,
		config.Config.Database.DBPass, config.Config.Database.DBHost, config.Config.Database.DBPort,
		config.Config.Database.DBName, val.Encode())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
		return nil, err
	}
	fmt.Println("Connection database success")
	return db.Debug(), nil
}
