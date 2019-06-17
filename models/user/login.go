package user

import (
	"errors"
)

var (
	id_q = "id = ?"
)

// Loing 登入方法
func Login_user(user *User) (*User, error) {

	// //route引用
	// user := new(User)
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	ResponseWithJson(w, http.StatusBadRequest,
	// 		Response{Code: http.StatusBadRequest, Msg: "bad params"})
	// 	return
	// }
	ru := new(User)
	q := db.Where(id_q, user.Id).Find(ru)
	if q.Error != nil {
		return nil, errors.New("user login fail : can't find user id")
	}
	if ru.decodePw(user.Pw) != nil {
		return nil, errors.New("user login fail : can't find user password")
	}
	return ru, nil
}
