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

func (UserRepository) toDomain(user *ent.User) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Active:    user.Active,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserRepository(uc *ent.UserClient) *UserRepository {
	return &UserRepository{
		UserClient: uc,
	}
}

func (ur *UserRepository) FindUser(ctx context.Context, u User) (*User, error) {
	user, err := ur.
		UserClient.
		Query().
		Where(
			user.Or(
				user.ID(u.ID),
				user.Email(u.Email),
			),
		).
		Only(ctx)

	if err != nil {
		return &User{}, err
	}

	return ur.toDomain(user), err
}

func (ur *UserRepository) FindUsers(ctx context.Context, u User, q common.Query) ([]*User, int, error) {
	query := []predicate.User{}

	if q.Search != "" {
		query = append(query, user.NameContainsFold(q.Search))
	}

	count, err := ur.
		UserClient.
		Query().
		Where(query...).
		Count(ctx)

	if err != nil {
		return []*User{}, count, err
	}

	order := ent.Asc(q.Sort)

	if q.Order == common.Desc {
		order = ent.Desc(q.Sort)
	}

	users, err := ur.
		UserClient.
		Query().
		Where(query...).
		Offset((q.Page - 1) * q.Size).
		Limit(q.Size).
		Order(order).
		All(ctx)

	if err != nil {
		return []*User{}, count, err
	}

	data := []*User{}

	for _, user := range users {
		data = append(data, ur.toDomain(user))
	}

	return data, count, err
}

func (ur *UserRepository) CreateUser(ctx context.Context, u *User) (*User, error) {
	user, err := ur.
		UserClient.
		Create().
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetImage(u.Image).
		SetActive(u.Active).
		Save(ctx)

	if err != nil {
		return &User{}, err
	}

	return ur.toDomain(user), err
}

func (ur *UserRepository) UpdateUser(ctx context.Context, condition User, u *User) (*User, error) {
	executeUpdate := ur.
		UserClient.
		UpdateOneID(condition.ID)

	if u.Name != "" {
		executeUpdate.SetName(u.Name)
	}

	if u.Password != "" {
		executeUpdate.SetPassword(u.Password)
	}

	if u.Image != "" {
		executeUpdate.SetImage(u.Image)
	}

	user, err := executeUpdate.
		SetActive(u.Active).
		Save(ctx)

	if err != nil {
		return &User{}, err
	}

	return ur.toDomain(user), err
}
