package db

import (
	"context"

	"codegram/ent"
	"codegram/ent/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func QueryUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, id uuid.UUID) (*ent.User, error) {
	u, err := client.User.Query().Where(user.ID(id)).Only(ctx)
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
		SetPassword(payload.Password).
		SetHashedPassword(string(hashed_pass)).
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

func AddUserFollower(ctx context.Context,
	client *ent.Client,
	uid uuid.UUID, fuids []uuid.UUID) (int, error) {
	u, err := client.
		User.
		Update().
		Where(user.ID(uid)).
		AddFollowerIDs(fuids...).
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
		Save(ctx)

	if err != nil {
		return -1, err
	}
	return u, nil
}

type CUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
