package auth

import (
	"fmt"
	"gopkg-spider/helper"
	"gopkg-spider/models"
	"gopkg-spider/pkg/setting"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretWord string = "ajoidf123"

func init() {
	sec, _ := setting.Cfg.GetSection("JWT")
	secretWord = sec.Key("SECRET").String()
	fmt.Println("Please enter JWT secret word (default " + secretWord + "): ")
	fmt.Scanln(&secretWord)
}

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // expiration after 2 hours
	})

	return token.SignedString([]byte(secretWord))
}

// JWT decode
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			helper.ResponseWithJson(w, http.StatusUnauthorized,
				helper.Response{Code: http.StatusUnauthorized, Msg: "authorized is \"\""})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					helper.ResponseWithJson(w, http.StatusUnauthorized,
						helper.Response{Code: http.StatusUnauthorized, Msg: "not ok authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte(secretWord), nil
			})
			if !token.Valid {
				helper.ResponseWithJson(w, http.StatusUnauthorized,
					helper.Response{Code: http.StatusUnauthorized, Msg: "!token.Valid authorized"})
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
