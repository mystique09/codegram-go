package routes

import (
	"codegram/db"
	"context"
	"fmt"
	"net/http"

	"codegram/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	UserRoute struct {
		Client *ent.Client
	}

	Error struct {
		Message string `json:"error"`
	}
)

func (user_rt *UserRoute) GetAllUser(c echo.Context) error {
	res, err := db.QueryUsers(context.Background(),
		user_rt.Client)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (user_rt *UserRoute) CreateUser(c echo.Context) error {
	payload := new(db.CUser)

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}

	if payload.Username == "" ||
		payload.Password == "" ||
		payload.Email == "" {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: "Missing required fields."})
	}

	res, err := db.CreateUser(context.Background(),
		user_rt.Client, *payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, res)
}

func (user_rt *UserRoute) GetUserById(c echo.Context) error {
	uid, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}

	res, err := db.QueryUser(context.Background(),
		user_rt.Client, uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (user_rt *UserRoute) UpdateUser(c echo.Context) error {
	payload := new(db.UUser)
	uid, uuid_err := uuid.Parse(c.Param("id"))

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: uuid_err.Error()})
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, &Error{Message: err.Error()})
	}

	res, err := db.UpdateUserInfo(context.Background(),
		user_rt.Client, uid, *payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,
		fmt.Sprintf("User updated: %v", res))
}

func (user_rt *UserRoute) DeleteUser(c echo.Context) error {
	uid, uuid_err := uuid.Parse(c.Param("id"))

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: uuid_err.Error()})
	}

	res, err := db.DeleteUser(context.Background(),
		user_rt.Client, uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			&Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,
		fmt.Sprintf("User deleted: %v", res))
}
