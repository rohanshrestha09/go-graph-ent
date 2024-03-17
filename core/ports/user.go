package ports

import (
	"context"

	"github.com/rohanshrestha09/go-graph-ent/common"
	. "github.com/rohanshrestha09/go-graph-ent/core/domains"
)

type UserRepository interface {
	FindUser(context.Context, User) (*User, error)
	FindUsers(context.Context, User, common.Query) ([]*User, int, error)
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, User, *User) (*User, error)
}
