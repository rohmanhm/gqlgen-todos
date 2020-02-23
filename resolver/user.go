package resolver

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/auth"
	"github.com/rohmanhm/gqlgen-todos/firebase"
	"github.com/rohmanhm/gqlgen-todos/graph/model"
	"google.golang.org/api/iterator"
)

// CreateUser Mutation.
func (r *Mutation) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	params := (&auth.UserToCreate{}).
		Email(input.Email).
		EmailVerified(false).
		Password(input.Password).
		DisplayName(input.Name).
		Disabled(false)
	u, err := firebase.Client.CreateUser(ctx, params)
	if err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}
	log.Printf("Successfully created user: %v\n", u)
	return &model.User{
		Email: u.Email,
		ID:    u.UID,
		Name:  u.DisplayName,
	}, nil
}

// Login resolver.
func (r *Mutation) Login(ctx context.Context, input model.LoginUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users query resolver.
func (r *Query) Users(ctx context.Context) ([]*model.User, error) {
	uIter := firebase.Client.Users(context.Background(), "")
	for {
		u, err := uIter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}

		r.users = append(r.users, &model.User{
			Email: u.Email,
			ID:    u.UID,
			Name:  u.DisplayName,
		})
	}

	return r.users, nil
}

// User query resolver.
func (r *Query) User(ctx context.Context, input *model.UserInput) (*model.User, error) {
	u, err := firebase.Client.GetUserByEmail(context.Background(), input.Email)
	if err != nil {
		log.Printf("could not get user: %v", err)
		return nil, err
	}

	return &model.User{
		Email: u.Email,
		ID:    u.UID,
		Name:  u.DisplayName,
	}, nil
}
