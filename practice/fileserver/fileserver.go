package main

import (
	"github.com/jou66jou/go-myresume/fileserver/errs"
	"github.com/labstack/echo"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Fields map[string]interface{}

func main() {
	e := echo.New()
	// setting
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Debug = false
	// cover all api error response
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				logFields := Fields{}

				// get request data
				req := c.Request()
				{
					logFields["requestMethod"] = req.Method
					logFields["requestURL"] = req.URL.String()
				}

				// get response data
				resp := c.Response()
				{
					logFields["responseStatus"] = resp.Status
				}
				log.Printf("%+v, error message : %+v\n", logFields, err)
			}
			return err
		}
	})

	// route
	e.POST("/upload", upload)
	// start
	e.Logger.Fatal(e.StartServer(s))
}

func upload(c echo.Context) error {
	// Read form fields
	// cardNum := c.FormValue("cardNum")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return errs.NewWithMessage(errs.ErrInternalError, "bind request file fail, err:"+err)
	}
	src, err := file.Open()
	if err != nil {
		return errs.NewWithMessage(errs.ErrInternalError, "file open fail, err:"+err)
	}
	defer src.Close()

	// Create directory if needed.
	nowPath, _ := os.Getwd()
	if os.MkdirAll(nowPath, 0666) != nil {
		panic("Unable to create directory for tagfile!")
	}
	dirName := time.Now().Format("20060102")
	// Destination
	dst, err := os.Create(strings.Join([]string{nowPath, dirName, file.Filename}, "/"))
	if err != nil {
		return errs.NewWithMessage(errs.ErrInternalError, "os create fail, err:"+err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return errs.NewWithMessage(errs.ErrInternalError, "io copy fail, err:"+err)
	}

	return c.NoContent(http.StatusNoContent)
}
