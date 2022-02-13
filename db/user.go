package db

import (
	"context"
	"time"

	"codegram/ent"
	"codegram/ent/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func QueryUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.
		User.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, id uuid.UUID) (*ent.User, error) {
	u, err := client.
		User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func QueryUserByUname(ctx context.Context,
	client *ent.Client, payload LUser) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Username(payload.Ussername)).Only(ctx)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func CreateUser(ctx context.Context,
	client *ent.Client, payload CUser) (*ent.User, error) {

	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err := client.
		User.
		Create().
		SetUsername(payload.Username).
		SetPassword(string(hashed_pass)).
		SetEmail(payload.Email).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateUserInfo(ctx context.Context,
	client *ent.Client,
	uid uuid.UUID,
	payload UUser) (int, error) {
	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		SetUsername(payload.Username).
		SetPassword(payload.Password).
		SetEmail(payload.Email).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return u, nil
}

func DeleteUser(ctx context.Context, client *ent.Client, uid uuid.UUID) (int, error) {
	u, err := client.User.
		Delete().
		Where(user.ID(uid)).
		Exec(ctx)
	if err != nil {
		return -1, err
	}

	return u, nil
}

func QueryUserPosts(ctx context.Context,
	client *ent.Client, auid uuid.UUID) ([]*ent.Post, error) {
	up, err := client.User.
		Query().
		Where(user.ID(auid)).
		QueryPosts().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return up, nil
}

func AddUserFollower(ctx context.Context,
	client *ent.Client,
	uid uuid.UUID, fuids []uuid.UUID) (int, error) {
	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		AddFollowerIDs(fuids...).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return u, nil
}

func AddUserFollowing(ctx context.Context,
	client *ent.Client, uid uuid.UUID, fuid uuid.UUID) (int, error) {
	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		AddFollowingIDs(fuid).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return u, nil
}

func RemoveUserFollowing(ctx context.Context,
	client *ent.Client,
	uid uuid.UUID,
	fuid uuid.UUID) (int, error) {
	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		RemoveFollowingIDs(fuid).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return -1, nil
	}

	return u, nil
}

func UpdateUserRole(ctx context.Context,
	client *ent.Client,
	uid uuid.UUID,
	new_role string) (int, error) {
	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		SetRole(user.Role(new_role)).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return -1, err
	}
	return u, nil
}

func UpdateUserStatus(ctx context.Context,
	client *ent.Client,
	uid uuid.UUID,
	new_status string) (int, error) {

	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		SetStatus(user.Status(new_status)).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return -1, err
	}
	return u, nil
}

type (
	CUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	LUser struct {
		Ussername string `json:"username"`
		Password  string `json:"password"`
	}
)
