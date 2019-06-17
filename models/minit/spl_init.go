package minit

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Mysql_init() *gorm.DB {
	var (
		err                                  error
		dbType, dbName, user, password, host = "mysql", "db", "test", "test", "127.0.0.1:3307"
	)

	for {
		ScanDB_config(&dbType, &dbName, &user, &password, &host)

		fmt.Println("Models: db ready to connect...", dbType)
		DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
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

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	fmt.Println("Models: db connect success")
	return DB
}

func CloseDB() {
	defer DB.Close()
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
