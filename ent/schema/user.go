package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("username").
			NotEmpty().
			Unique(),
		field.String("password").
			NotEmpty().
			MinLen(8).
			MaxLen(12).
			Sensitive(),
		field.String("hashed_password").
			NotEmpty().
			Unique(),
		field.String("email").
			NotEmpty().
			Unique(),
		field.Enum("role").
			Values("Normal", "Admin").
			Default("Normal"),
		field.Enum("status").
			Values("offline", "online").
			Default("offline"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("followers", User.Type),
		edge.To("following", User.Type),
	}
}
