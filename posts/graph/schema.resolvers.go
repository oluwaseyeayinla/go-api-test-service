package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/posts/graph/generated"
	"github.com/zalbiraw/go-api-test-service/posts/graph/helpers"
	"github.com/zalbiraw/go-api-test-service/posts/graph/model"
)

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	postId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	posts, err := helpers.GetPosts()
	if nil != err {
		return nil, err
	}

	return &((*posts)[postId]), nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	postsArray, err := helpers.GetPosts()
	if nil != err {
		return nil, err
	}

	posts := make([]*model.Post, len(*postsArray))
	for i := 0; i < len(*postsArray); i++ {
		posts[i] = &((*postsArray)[i])
	}

	return posts, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
