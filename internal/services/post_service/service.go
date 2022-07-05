package post_service

import (
	"github.com/idanieldrew/blog-golang/internal/domain/post"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
)

var (
	PostService postInterfaceService = &postService{}
)

type (
	postService struct{}

	postInterfaceService interface {
		Get(slug string) (*post.Post, *restError.RestError)
	}
)

func (p *postService) Get(slug string) (*post.Post, *restError.RestError) {
	res := &post.Post{
		Slug: slug,
	}

	if err := res.Get(); err != nil {
		return nil, err
	}

	return res, nil
}
