package db

import (
	"codegram/ent"
	"codegram/ent/post"
	"context"

	"github.com/google/uuid"
)

func QueryPosts(ctx context.Context, client *ent.Client) ([]*ent.Post, error) {
	p, err := client.
		Post.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return p, nil
}

func QueryPost(ctx context.Context, client *ent.Client, puid uuid.UUID) (*ent.Post, error) {
	p, err := client.
		Post.
		Query().
		Where(post.ID(puid)).
		Only(ctx)

	if err != nil {
		return nil, err
	}
	return p, nil
}

func CreatePost(ctx context.Context, client *ent.Client, payload CPost) (*ent.Post, error) {
	p, err := client.
		Post.
		Create().
		SetAuthorID(payload.Author).
		SetTitle(payload.Title).
		SetDescription(payload.Description).
		SetImage(payload.Image).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return p, nil
}

func UpdatePost(ctx context.Context, client *ent.Client, puid uuid.UUID, payload UPost) (int, error) {
	p, err := client.
		Post.
		Update().
		Where(post.ID(puid)).
		SetTitle(payload.Title).
		SetDescription(payload.Description).
		SetImage(payload.Image).
		Save(ctx)

	if err != nil {
		return -1, err
	}
	return p, nil
}

func DeletetePost(ctx context.Context, client *ent.Client, puid uuid.UUID) (int, error) {
	p, err := client.
		Post.
		Delete().
		Where(post.ID(puid)).
		Exec(ctx)

	if err != nil {
		return -1, err
	}
	return p, nil
}

type (
	CPost struct {
		Author      uuid.UUID `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Image       string    `json:"image"`
	}

	UPost struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}
)
