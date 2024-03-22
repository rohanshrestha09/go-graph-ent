package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rohanshrestha09/go-graph-ent/common"
	. "github.com/rohanshrestha09/go-graph-ent/core/domains"
	"github.com/rohanshrestha09/go-graph-ent/ent"
	"github.com/rohanshrestha09/go-graph-ent/ent/blog"
	"github.com/rohanshrestha09/go-graph-ent/ent/predicate"
)

type BlogRepository struct {
	BlogClient *ent.BlogClient
}

func toBlogDomain(blog *ent.Blog) *Blog {
	b := &Blog{
		ID:        blog.ID,
		Title:     blog.Title,
		Slug:      blog.Slug,
		Content:   blog.Content,
		Image:     blog.Image,
		Status:    blog.Status,
		UserID:    blog.UserID,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}

	if blog.Edges.User != nil {
		b.User = toUserDomain(blog.Edges.User)
	}

	return b
}

func NewBlogRepository(bc *ent.BlogClient) *BlogRepository {
	return &BlogRepository{
		BlogClient: bc,
	}
}

func (br *BlogRepository) FindBlog(ctx context.Context, b Blog) (*Blog, error) {
	blog, err := br.
		BlogClient.
		Query().
		Where(
			blog.Or(
				blog.ID(b.ID),
				blog.Slug(b.Slug),
			)).
		WithUser().
		Only(ctx)

	if err != nil {
		return &Blog{}, err
	}

	return toBlogDomain(blog), err
}

func (br *BlogRepository) FindBlogs(ctx context.Context, b Blog, q common.Query) ([]*Blog, int, error) {
	query := []predicate.Blog{}

	if b.UserID != uuid.Nil {
		query = append(query, blog.UserID(b.UserID))
	}

	if q.Search != "" {
		query = append(query, blog.TitleContainsFold(q.Search))
	}

	count, err := br.
		BlogClient.
		Query().
		Where(query...).
		Count(ctx)

	if err != nil {
		return []*Blog{}, count, err
	}

	order := ent.Asc(q.Sort)

	if q.Order == common.Desc {
		order = ent.Desc(q.Sort)
	}

	blogs, err := br.
		BlogClient.
		Query().
		Where(query...).
		WithUser().
		Offset(
			(q.Page - 1) * q.Size,
		).
		Limit(q.Size).
		Order(order).
		All(ctx)

	if err != nil {
		return []*Blog{}, count, err
	}

	data := []*Blog{}

	for _, blog := range blogs {
		data = append(data, toBlogDomain(blog))
	}

	return data, count, err
}

func (br *BlogRepository) CreateBlog(ctx context.Context, b *Blog) (*Blog, error) {
	blog, err := br.
		BlogClient.
		Create().
		SetTitle(b.Title).
		SetSlug(b.Slug).
		SetContent(b.Content).
		SetImage(b.Image).
		SetStatus(b.Status).
		SetUserID(b.UserID).
		Save(ctx)

	if err != nil {
		return &Blog{}, err
	}

	return toBlogDomain(blog), err
}

func (br *BlogRepository) UpdateBlog(ctx context.Context, condition Blog, b *Blog) (*Blog, error) {
	blogId, err := br.
		BlogClient.
		Query().
		Where(
			blog.Or(
				blog.ID(condition.ID),
				blog.Slug(condition.Slug),
			)).
		FirstID(ctx)

	executeUpdate := br.
		BlogClient.
		UpdateOneID(blogId)

	if b.Title != "" {
		executeUpdate.SetTitle(b.Title)
	}

	if b.Content != "" {
		executeUpdate.SetContent(b.Content)
	}

	if b.Image != "" {
		executeUpdate.SetImage(b.Image)
	}

	if b.Status != "" {
		executeUpdate.SetStatus(b.Status)
	}

	blog, err := executeUpdate.
		Where(blog.Slug(b.Slug)).
		Save(ctx)

	if err != nil {
		return &Blog{}, err
	}

	return toBlogDomain(blog), err
}
