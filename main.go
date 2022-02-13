package main

import (
	"codegram/db"
	"codegram/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client := db.SetupDb()

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", homeRoute)

	user_group := app.Group("/user")
	{
		user_rt := routes.UserRoute{
			Client: client,
		}
		user_group.GET("", user_rt.GetAllUser)
		user_group.GET("/:id", user_rt.GetUserById)
		user_group.POST("", user_rt.CreateUser)
		user_group.PUT("/:id", user_rt.UpdateUser)
		user_group.DELETE("/:id", user_rt.DeleteUser)
	}

	post_group := app.Group("/post")
	{
		post_rt := routes.PostRoute{
			Client: client,
		}
		post_group.GET("", post_rt.GetAllPost)
		post_group.GET("/:id", post_rt.GetPostById)
		post_group.POST("", post_rt.CreatePost)
		post_group.PUT("/:id", post_rt.UpdatePost)
		post_group.DELETE("/:id", post_rt.DeletePost)
	}

	app.Logger.Fatal(app.Start(":3000"))
}

func homeRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Codegram!")
}
