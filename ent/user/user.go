// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowing holds the string denoting the following edge name in mutations.
	EdgeFollowing = "following"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "user_posts"
	// FollowersTable is the table that holds the followers relation/edge. The primary key declared below.
	FollowersTable = "user_followers"
	// FollowingTable is the table that holds the following relation/edge. The primary key declared below.
	FollowingTable = "user_following"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldEmail,
	FieldRole,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// FollowersPrimaryKey and FollowersColumn2 are the table columns denoting the
	// primary key for the followers relation (M2M).
	FollowersPrimaryKey = []string{"user_id", "follower_id"}
	// FollowingPrimaryKey and FollowingColumn2 are the table columns denoting the
	// primary key for the following relation (M2M).
	FollowingPrimaryKey = []string{"user_id", "following_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// RoleNormal is the default value of the Role enum.
const DefaultRole = RoleNormal

// Role values.
const (
	RoleNormal Role = "Normal"
	RoleAdmin  Role = "Admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleNormal, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusOffline is the default value of the Status enum.
const DefaultStatus = StatusOffline

// Status values.
const (
	StatusOffline Status = "offline"
	StatusOnline  Status = "online"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOffline, StatusOnline:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}
