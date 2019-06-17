package user

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Index    int    `sql:"AUTO_INCREMENT"`
	Name     string `bson:"username" json:"username"`
	Id       string `bson:"id" json:"id" sql:"primary_key";"not null;unique;type:varchar(32);unique_index`
	Pw       string `bson:"pw" json:"pw" sql:"not null"`
	Email    string `bson:"email" bson:"email sql:"not null;unique;type:varchar(120);unique_index"`
	SignDate string `bson:"-" json:"-"`
}

const _TimeFormat = "2006-01-02 15:04:05"
const _hash = true

var (
	db  *gorm.DB
	dbl = false
)

func InitDB(odb *gorm.DB) {
	db = odb
	dbl = true
}

func (u *User) CreateUserTable() error {

	if db.HasTable(&User{}) {
		return errors.New("db can't create user table : has user table in db")
	}
	db.CreateTable(&User{})
	return nil
}

func (u *User) decodePw(pw string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Pw), []byte(pw))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *User) encodePw() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Pw), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return err
	}
	u.Pw = string(hash)
	return nil
}
