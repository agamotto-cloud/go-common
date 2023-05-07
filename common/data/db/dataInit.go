package data

import (
	"github.com/agamotto-cloud/go-common/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var GlobalDB *gorm.DB

func init() {
	dbConfig := config.GetConfig("db", config.MysqlConfig{})
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Url + ")/" + dbConfig.Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatal(err)
		return
	}

	GlobalDB = db
	log.Println("Gorm client initialized successfully.")
}
