package helper

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

func Register(routes []Route, method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}

func AddLoginRouters(router mux.Router, routes []Route) mux.Router {
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil { // JWT valid
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router

}

// JWT decode
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			ResponseWithJson(w, http.StatusUnauthorized,
				Response{Code: http.StatusUnauthorized, Msg: "authorized is \"\""})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					ResponseWithJson(w, http.StatusUnauthorized,
						Response{Code: http.StatusUnauthorized, Msg: "not ok authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte(secretWord), nil
			})
			if !token.Valid {
				ResponseWithJson(w, http.StatusUnauthorized,
					Response{Code: http.StatusUnauthorized, Msg: "!token.Valid authorized"})
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
