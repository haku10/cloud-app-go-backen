package service

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
  DBMS := "mysql"

	CONNECT := os.Getenv("MYSQL_ENDPOINT") + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	log.Print(CONNECT)
	connection, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return connection
}
