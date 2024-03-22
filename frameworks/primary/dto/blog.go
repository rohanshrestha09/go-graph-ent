package dto

import (
	"github.com/rohanshrestha09/go-graph-ent/common"
	"github.com/rohanshrestha09/go-graph-ent/common/enums"
)

type QueryBlogDto struct {
	*common.Query
	UserID string `form:"userId"`
}

type CreateBlogDto struct {
	Title   string           `json:"title" binding:"required" validate:"required"`
	Content string           `json:"content" binding:"required" validate:"required"`
	Image   string           `json:"image"`
	Status  enums.BlogStatus `json:"status" validate:"omitempty,oneof=PUBLISHED UNPUBLISHED"`
	UserID  string           `json:"userId" binding:"required" validate:"required,uuid"`
}

type UpdateBlogDto struct {
	Title   string           `json:"title"`
	Content string           `json:"content"`
	Image   string           `json:"image"`
	Status  enums.BlogStatus `json:"status" validate:"omitempty,oneof=PUBLISHED UNPUBLISHED"`
}
