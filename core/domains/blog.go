package domains

import (
	"time"

	"github.com/google/uuid"
	"github.com/rohanshrestha09/go-graph-ent/common/enums"
)

type Blog struct {
	ID        int              `json:"id"`
	Slug      string           `json:"slug"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	Image     string           `json:"image"`
	Status    enums.BlogStatus `json:"status"`
	UserID    uuid.UUID        `json:"userId"`
	User      *User            `json:"user"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}
