package user

import "time"

func Register_user(user *User) error {
	q := db.Where(id_q, user.Id).Find()

	user.SignDate = time.Now().Format("2006-01-02 15:04:05")
	if _hash {
		err := user.encodePw()
		if err != nil {
			return err
		}
	}
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
