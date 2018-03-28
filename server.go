package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// Reading staic file, like when you want to render static html file, you can use File method.
	e.File("/", "static/index.html")

	// You can route with REST METHODS, like e.GET or e.POST, e.PUT, e.DELETE.
	// For example, we route getUser method with GET method.
	e.GET("/users/:id", getUser)

	// If you need query parameters in url, example is in show function.
	e.GET("/show", show)

	// If you are posting values from form, example is in save function.
	e.POST("/save", save)
	e.POST("/save/avatar", saveAvatar)

	// Add new data using with structure.
	e.POST("/users", addUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// Get name and email from form value
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveAvatar(c echo.Context) error {
	// Get name from form value
	name := c.FormValue("name")

	// Get avatar things
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	// Set Close function to src, because src is file object and it should be closed when func is returned
	defer src.Close()

	// Destination
	dest, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	// Set Close function to dest, same reason as src should be closed.
	defer src.Close()

	// Copy
	if _, err = io.Copy(dest, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func addUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
