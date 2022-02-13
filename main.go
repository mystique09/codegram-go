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

	/* Routes */
	user_rt := routes.UserRoute{
		Client: client,
	}
	post_rt := routes.PostRoute{
		Client: client,
	}

	app := echo.New()
	app.Use(routes.CustomLogger())
	app.Use(middleware.Recover())

	app.GET("/api", homeRoute)
	user_group := app.Group("/api/users")
	{
		user_group.GET("", user_rt.GetAllUser)
		user_group.GET("/:id", user_rt.GetUserById)
		user_group.POST("", user_rt.CreateUser)
		app.POST("/auth", user_rt.LoginUser)
	}

	jwt_group := app.Group("/api/private", routes.AuthMiddleware())

	// Private users routes.
	priv_user_group := jwt_group.Group("/users")
	{
		priv_user_group.PUT("/:id", user_rt.UpdateUser)
		priv_user_group.DELETE("/:id", user_rt.DeleteUser)
	}
	// Private posts routes.
	priv_post_group := jwt_group.Group("/posts")
	{
		priv_post_group.GET("", post_rt.GetAllPost)
		priv_post_group.GET("/:id", post_rt.GetPostById)
		priv_post_group.POST("", post_rt.CreatePost)
		priv_post_group.PUT("/:id", post_rt.UpdatePost)
		priv_post_group.DELETE("/:id", post_rt.DeletePost)
	}

	app.Logger.Fatal(app.Start(":3000"))
}

func homeRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Codegram!")
}
