package post_service

import (
	"github.com/idanieldrew/blog-golang/internal/domain/post"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"time"
)

var (
	PostService postInterfaceService = &postService{}
)

type (
	postService struct{}

	postInterfaceService interface {
		Get(string) (*post.Post, *restError.RestError)
		Store(post.Post) (*post.Post, *restError.RestError)
	}
)

// Get service
func (p *postService) Get(slug string) (*post.Post, *restError.RestError) {
	res := &post.Post{
		Slug: slug,
	}

	if err := res.Get(); err != nil {
		return nil, err
	}
	return res, nil
}

// Store service

func (p *postService) Store(post post.Post) (*post.Post, *restError.RestError) {
	post.UserId = 1
	post.CreatedAt, post.UpdatedAt = time.Now(), time.Now()

	if savePostErr := post.Save(); savePostErr != nil {
		return nil, savePostErr
	}

	return &post, nil
}
