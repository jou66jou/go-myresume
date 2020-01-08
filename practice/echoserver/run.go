package main

import (
	_ "github.com/jou66jou/go-myresume/practice/echos/statik"
	"github.com/labstack/echo"
	"github.com/rakyll/statik/fs" 
)

func main() {

	e := echo.New()
	statikFS, err := fs.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Static("/", statikFS)
	e.GET("/", func(c echo.Context) error {
		return c.File("public/views/index.html")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
