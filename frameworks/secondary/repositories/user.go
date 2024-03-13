package repositories

import (
	"context"

	"github.com/rohanshrestha09/go-graph-ent/common"
	. "github.com/rohanshrestha09/go-graph-ent/core/domains"
	"github.com/rohanshrestha09/go-graph-ent/ent"
	"github.com/rohanshrestha09/go-graph-ent/ent/predicate"
	"github.com/rohanshrestha09/go-graph-ent/ent/user"
)

type UserRepository struct {
	UserClient *ent.UserClient
}

func toDomain(user *ent.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserRepository(uc *ent.UserClient) *UserRepository {
	return &UserRepository{
		UserClient: uc,
	}
}

func (ur *UserRepository) FindUser(ctx context.Context, u User) (User, error) {
	user, err := ur.
		UserClient.
		Query().
		Where(user.Or(user.ID(u.ID))).
		Only(ctx)

	if err != nil {
		return User{}, err
	}

	return toDomain(user), err
}

func (ur *UserRepository) FindUsers(ctx context.Context, u User, pagination common.Pagination) ([]User, int, error) {
	query := []predicate.User{}

	count, err := ur.
		UserClient.
		Query().
		Where(query...).
		Count(ctx)

	if err != nil {
		return []User{}, count, err
	}

	order := ent.Asc(pagination.Sort)

	if pagination.Order == common.DESC {
		order = ent.Desc(pagination.Sort)
	}

	users, err := ur.
		UserClient.
		Query().
		Where(query...).
		Offset((pagination.Page - 1) * pagination.Size).
		Limit(pagination.Size).
		Order(order).
		All(context.Background())

	if err != nil {
		return []User{}, count, err
	}

	data := []User{}

	for _, user := range users {
		data = append(data, toDomain(user))
	}

	return data, count, err
}

func (ur *UserRepository) CreateUser(ctx context.Context, u *User) (User, error) {
	user, err := ur.
		UserClient.
		Create().
		SetName(u.Name).
		SetAge(u.Age).
		Save(ctx)

	if err != nil {
		return User{}, err
	}

	return toDomain(user), err
}

func (ur *UserRepository) UpdateUser(ctx context.Context, u *User) (User, error) {
	user, err := ur.
		UserClient.
		UpdateOneID(u.ID).
		SetName(u.Name).
		SetAge(u.Age).
		Save(ctx)

	return toDomain(user), err
}
