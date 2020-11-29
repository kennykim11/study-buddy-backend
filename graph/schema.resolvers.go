package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/kennykim11/study-buddy-backend/graph/generated"
	"github.com/kennykim11/study-buddy-backend/graph/model"
	"github.com/kennykim11/study-buddy-backend/repository"
)

var repo repository.Repository = repository.New()
var truevar = true

func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	user := &model.User{
		GoogleID:       newUser.GoogleID,
		CreatedAt:      int(time.Now().Unix()),
		LastModifiedAt: int(time.Now().Unix()),
		GivenName:      newUser.GivenName,
		FamilyName:     newUser.FamilyName,
		Contacts:       []*model.Contact{},
		Taking:         []int{},
	}
	repo.SaveUser(ctx, user)
	return user, nil
}

func (r *mutationResolver) EnrollInSection(ctx context.Context, userID int, sectionID int) (*bool, error) {
	repo.EnrollInSection(ctx, userID, sectionID)
	return &truevar, nil
}

func (r *mutationResolver) RemoveFromSection(ctx context.Context, userID int, sectionID int) (*bool, error) {
	repo.RemoveFromSection(ctx, userID, sectionID)
	return &truevar, nil
}

func (r *mutationResolver) RegisterContactInfo(ctx context.Context, userID int, contact model.ContactInput) (*bool, error) {
	repo.RegisterContactInfo(ctx, userID, contact)
	return &truevar, nil
}

func (r *mutationResolver) SendContactMessage(ctx context.Context, fromUserID int, toUserID int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
	return &truevar, nil
}

func (r *queryResolver) User(ctx context.Context, googleID int) (*model.User, error) {
	return repo.FindUser(ctx, googleID), nil
}

func (r *queryResolver) Section(ctx context.Context, sectionID int) (*model.Section, error) {
	return repo.FindSection(ctx, sectionID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
