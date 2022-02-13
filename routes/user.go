package routes

import (
	"codegram/db"
	"context"
	"net/http"

	"codegram/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	UserRoute struct {
		Client *ent.Client
	}
)

func (user_rt *UserRoute) GetAllUser(c echo.Context) error {
	res, err := db.QueryUsers(context.Background(),
		user_rt.Client)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully queried all users.", res))
}

func (user_rt *UserRoute) CreateUser(c echo.Context) error {
	payload := new(db.CUser)

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(err.Error()))
	}

	if payload.Username == "" ||
		payload.Password == "" ||
		payload.Email == "" {
		return c.JSON(http.StatusBadRequest,
			NewError("Missing required fields."))
	}

	res, err := db.CreateUser(context.Background(),
		user_rt.Client, *payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewResponse(true, "Successfully created one user.", res))
}

func (user_rt *UserRoute) GetUserById(c echo.Context) error {
	uid, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(err.Error()))
	}

	res, err := db.QueryUser(context.Background(),
		user_rt.Client, uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully queried one user.", res))
}

func (user_rt *UserRoute) GetAllUsersPosts(c echo.Context) error {
	auid, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}

	p, err := db.QueryUserPosts(context.Background(), user_rt.Client, auid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully queried all users post.", p))
}

func (user_rt *UserRoute) UpdateUser(c echo.Context) error {
	payload := new(db.UUser)
	uid, uuid_err := uuid.Parse(c.Param("id"))

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(uuid_err.Error()))
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}

	res, err := db.UpdateUserInfo(context.Background(),
		user_rt.Client, uid, *payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully updated one user.", res))
}

func (user_rt *UserRoute) DeleteUser(c echo.Context) error {
	uid, uuid_err := uuid.Parse(c.Param("id"))

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest,
			NewError(uuid_err.Error()))
	}

	res, err := db.DeleteUser(context.Background(),
		user_rt.Client, uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully deleted one user.", res))
}
