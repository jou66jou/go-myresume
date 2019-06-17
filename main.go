package main

import (
	"fmt"
	"gotest/models/minit"
	"gotest/models/user"
)

//test code
func main() {
	db := minit.Mysql_init()
	user.InitDB(db)

	u := user.User{Id: "test", Pw: "test"}
	err := u.CreateUserTable()
	if err != nil {
		fmt.Println(err)
	}
	user.Register_user(&u)
	user.Login_user(&u)
}
