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
	result, err := u.FindUser()
	if err != nil {
		fmt.Println(err)
		if err := u.Register(); err != nil {
			fmt.Println(err)
			return
		}
		result = u
	}
	testU := user.User{Id: "test", Pw: "test"}
	err = result.CheckPw(testU.Pw)
	if err != nil {
		fmt.Println("ERROR:CheckPW :", err)
	}
}
