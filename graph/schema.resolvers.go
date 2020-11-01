package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/kennykim11/study-buddy-backend/graph/generated"
	"github.com/kennykim11/study-buddy-backend/graph/model"
	"github.com/kennykim11/study-buddy-backend/graph/repository"
)

var repo = repository.UserRepository = repository.New()

func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	user := &model.User{
		GoogleID:       newUser.GoogleID,
		CreatedAt:      int(time.Now().Unix()),
		LastModifiedAt: int(time.Now().Unix()),
		GivenName:      newUser.GivenName,
		FamilyName:     newUser.FamilyName,
		Contacts:       []*model.Contact{},
		Taking:         []string{},
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *mutationResolver) EnrollInSection(ctx context.Context, userID int, sectionID int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RegisterContactInfo(ctx context.Context, userID int, contact model.ContactInput) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendContactMessage(ctx context.Context, fromUserID int, toUserID int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) Sections(ctx context.Context) ([]*model.Section, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Classmates(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
