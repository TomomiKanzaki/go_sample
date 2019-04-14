package main


import (
	"net/http"
	"github.com/labstack/echo"
)


func main(){
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		user := &User{
			ID: 100,
			Name: "sample user",
			Email: "sample@test.com",
		}
		return context.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	ID int `json:"id"`
	Name string `json: "name"`
	Email string `json; "email"`
}
