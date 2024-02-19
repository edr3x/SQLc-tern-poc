package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/edr3x/tern-sqlc-poc/internal/db/connection"
	"github.com/edr3x/tern-sqlc-poc/internal/db/query"
)

type userService struct {
	query *query.Queries
}

func UserService() *userService {
	return &userService{
		query: query.New(connection.PgPool),
	}
}

func (u *userService) CreateUser(ctx context.Context, input UserCreateInput) (query.User, error) {
	createUser := query.CreateUserParams{
		FirstName: pgtype.Text{String: input.FirstName, Valid: true},
		LastName:  pgtype.Text{String: input.LastName, Valid: true},
		Email:     pgtype.Text{String: input.Email, Valid: true},
		Password:  pgtype.Text{String: input.Password, Valid: true},
	}

	usr, err := u.query.CreateUser(ctx, createUser)
	if err != nil {
		return query.User{}, err
	}
	return usr, nil
}

func (u *userService) GetAllUsers(ctx context.Context) ([]query.User, error) {
	users, err := u.query.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) GetUserById(ctx context.Context, id string) (query.User, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return query.User{}, err
	}

	user, err := u.query.GetUser(ctx, pgtype.UUID{Bytes: parsedUUID, Valid: true})
	if err != nil {
		return query.User{}, err
	}

	return user, nil
}

func (u *userService) DeleteUser(ctx context.Context, id string) error {
	userId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = u.query.DeleteUser(ctx, pgtype.UUID{Bytes: userId, Valid: true})
	if err != nil {
		return err
	}

	return nil
}
