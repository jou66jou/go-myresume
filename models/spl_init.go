package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	Mysql_init()
}

func Mysql_init() {
	var (
		err                                  error
		dbType, dbName, user, password, host = "mysql", "db", "test", "test", "127.0.0.1:3307"
	)

	for {
		ScanDB_config(&dbType, &dbName, &user, &password, &host)

		fmt.Println("Models: db ready to connect...")
		db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))

		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	fmt.Println("Models: db connect success")

}

func CloseDB() {
	defer db.Close()
}

func ScanDB_config(dbType, dbName, user, password, host *string) {

	fmt.Println("Open up database connection, please sets config(no input as default value) ")
	fmt.Println("db type (default " + *dbType + "): ")
	SwitchScanf(dbType)
	fmt.Println("db name(default " + *dbName + "): ")
	SwitchScanf(dbName)
	fmt.Println("db user(default " + *user + "): ")
	SwitchScanf(user)
	fmt.Println("db password(default " + *password + "): ")
	SwitchScanf(password)
	fmt.Println("db host(default " + *host + "): ")
	SwitchScanf(host)

}

func SwitchScanf(v *string) {
	var s string
	fmt.Scanln(&s)
	if s != "" {
		*v = s
	}
}
