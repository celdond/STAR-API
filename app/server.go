package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type (
	Excuse struct {
		Error string `json:"error"`
		Id    string `json:"id"`
		Quote string `json:"quote"`
	}
)

func main() {
	e := echo.New()
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	e.GET("/check", func(c echo.Context) error {
		_, err := db.Query("SELECT 1")

		if err != nil {
			fmt.Println(err)
			response := Excuse{Id: "", Error: "true", Quote: ""}
			return c.JSON(http.StatusInternalServerError, response)
		}
		return c.String(http.StatusOK, "Server is running.")
	})
	e.Logger.Fatal(e.Start(":5167"))
}
