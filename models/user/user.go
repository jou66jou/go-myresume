package user

import (
	"errors"
	"fmt"
	"time"

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

	// sql query
	id_q = "id = ?"
)

func InitDB(odb *gorm.DB) {
	db = odb
	dbl = true
}

func (u *User) FindUser() (User, error) {
	var dbU = User{}
	q := db.Where(id_q, u.Id).Find(&dbU)
	if q.RecordNotFound() {
		return dbU, errors.New("user login fail : can't find user id")
	} else if q.Error != nil {
		return dbU, q.Error
	}
	fmt.Println(dbU)
	return dbU, nil
}

func (u *User) CreateUserTable() error {

	if db.HasTable(&User{}) {
		return errors.New("db can't create user table : has user table in db")
	}
	db.CreateTable(&User{})
	return nil
}

func (u *User) CheckPw(pw string) error {
	fmt.Println("u.Pw : ", (u.Pw))
	// fmt.Println("u.Pw (byte): ", []byte(u.Pw))
	fmt.Println("Pw (byte): ", []byte(pw))

	err := bcrypt.CompareHashAndPassword([]byte(u.Pw), []byte(pw))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) EncodePw() error {
	fmt.Println("u.Pw (byte): ", []byte(u.Pw))
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Pw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Pw = string(hash)
	return nil
}

func (u *User) Register() error {
	if _, err := u.FindUser(); err == nil {
		return errors.New("user can't register : has same id in table")
	}
	u.SignDate = time.Now().Format("2006-01-02 15:04:05")
	if _hash {
		err := u.EncodePw()
		if err != nil {
			return err
		}
	}
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
