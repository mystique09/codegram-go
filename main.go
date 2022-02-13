package main

import (
	"codegram/db"
	"codegram/routes"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client := db.SetupDb()

	JWT_SECRET := os.Getenv("JWT_SECRET")

	/* Routes */
	user_rt := routes.UserRoute{
		Client: client,
	}
	post_rt := routes.PostRoute{
		Client: client,
	}

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	app.Use(middleware.Recover())

	app.GET("/", homeRoute)
	user_group := app.Group("/users")
	{
		user_group.GET("", user_rt.GetAllUser)
		user_group.GET("/:id", user_rt.GetUserById)
		user_group.POST("", user_rt.CreateUser)
		user_group.POST("/auth", user_rt.LoginUser)
	}

	jwt_group := app.Group("/", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(JWT_SECRET),
	}))
	{
		post_group := jwt_group.Group("posts")
		// Post routes that need authorization.
		{
			post_group.POST("", post_rt.CreatePost)
			post_group.PUT("/:id", post_rt.UpdatePost)
			post_group.DELETE("/:id", post_rt.DeletePost)
		}
		// User routes that need authorization.
		user_jwt_group := jwt_group.Group("users")
		{
			user_jwt_group.GET("/:id/posts", user_rt.GetAllUsersPosts)
			user_jwt_group.PUT("/:id", user_rt.UpdateUser)
			user_jwt_group.DELETE("/:id", user_rt.DeleteUser)
		}
	}
	app.Logger.Fatal(app.Start(":3000"))
}

func homeRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Codegram!")
}
