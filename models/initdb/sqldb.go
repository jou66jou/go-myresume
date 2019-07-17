package initdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	host          = "127.0.0.1"
	port          = "3307"
	user          = "test"
	password      = "test"
	dbname        = "db"
	DbMaxConnect  = 10
	DbIdleConnect = 10
)

func initMysql() (*gorm.DB, error) {
	dbConStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)

	mydb, err := gorm.Open("mysql", dbConStr)
	if err != nil {
		return nil, err
	}

	mydb.DB().SetMaxOpenConns(DbMaxConnect)
	mydb.DB().SetMaxIdleConns(DbIdleConnect)

	return mydb, nil

}
