package routes

import (
	"codegram/db"
	"codegram/ent"
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	PostRoute struct {
		Client *ent.Client
	}
)

func (post_rt *PostRoute) GetAllPost(c echo.Context) error {
	p, err := db.QueryPosts(context.Background(), post_rt.Client)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully queried all post.", p))
}

func (post_rt *PostRoute) GetPostById(c echo.Context) error {
	puid, uuid_err := uuid.Parse(c.Param("id"))

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest, NewError(uuid_err.Error()))
	}

	p, err := db.QueryPost(context.Background(), post_rt.Client, puid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusOK, NewResponse(true, "Successfully queried one post.", p))
}

func (post_rt *PostRoute) CreatePost(c echo.Context) error {
	payload := new(db.CPost)

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}

	if payload.Author == uuid.Nil || payload.Title == "" {
		return c.JSON(http.StatusBadRequest, NewError("Missing required fields."))
	}

	p, err := db.CreatePost(context.Background(), post_rt.Client, *payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewResponse(true, "Successfully created one post.", p))
}

func (post_rt *PostRoute) UpdatePost(c echo.Context) error {
	puid, uuid_err := uuid.Parse(c.Param("id"))
	payload := new(db.UPost)

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest, NewError(uuid_err.Error()))
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}

	p, err := db.UpdatePost(context.Background(), post_rt.Client, puid, *payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewResponse(true, "Successfully updated one post.", p))
}

func (post_rt *PostRoute) DeletePost(c echo.Context) error {
	puid, uuid_err := uuid.Parse(c.Param("id"))

	if uuid_err != nil {
		return c.JSON(http.StatusBadRequest, NewError(uuid_err.Error()))
	}

	p, err := db.DeletetePost(context.Background(), post_rt.Client, puid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewResponse(true, "Successfully deleted one post.", p))
}
