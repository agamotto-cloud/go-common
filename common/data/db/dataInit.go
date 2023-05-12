package data

import (
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/agamotto-cloud/go-common/common/logger"
	logger2 "gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GlobalDB *gorm.DB

func init() {
	dbConfig := config.GetConfig("db", config.MysqlConfig{})
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Url + ")/" + dbConfig.Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         dbLog{},
	})

	if err != nil {
		//	log.Fatal(err)
		return
	}

	GlobalDB = db
	//log.Println("Gorm client initialized successfully.")
}

type dbLog struct {
	logger.C
}

func (d dbLog) LogMode(level logger2.LogLevel) logger2.Interface {
	return d
}
