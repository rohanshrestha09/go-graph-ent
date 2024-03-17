package ports

import (
	"context"

	"github.com/rohanshrestha09/go-graph-ent/common"
	. "github.com/rohanshrestha09/go-graph-ent/core/domains"
)

type BlogRepository interface {
	FindBlog(context.Context, Blog) (*Blog, error)
	FindBlogs(context.Context, Blog, common.Query) ([]*Blog, int, error)
	CreateBlog(context.Context, *Blog) (*Blog, error)
	UpdateBlog(context.Context, Blog, *Blog) (*Blog, error)
}
