package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"
	"gqlgen/graph/model"
	"location"
	"log"
)

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input model.NewLocation) (*model.Location, error) {
	locationService := location.Location(r.DB)
	location, err := locationService.CreateLocation(input)
	if err != nil {
		err := fmt.Errorf("fail create location, %w", err)
		log.Panicf("%v\n", err)
		return nil, err
	}
	return location, nil
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	locationService := location.Location(r.DB)
	location, err := locationService.GetLocations()
	if err != nil {
		err := fmt.Errorf("fail get location, %w", err)
		log.Panicf("%v\n", err)
		return nil, err
	}
	return location, nil
}

// Years is the resolver for the years field.
func (r *queryResolver) Years(ctx context.Context) ([]int, error) {
	panic(fmt.Errorf("not implemented: Years - years"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
