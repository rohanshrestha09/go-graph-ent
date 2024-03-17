package common

type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

type Query struct {
	Page   int    `form:"page,default=1"`
	Size   int    `form:"size,default=10"`
	Sort   string `form:"sort,default=id"`
	Order  Order  `form:"order,default=desc"`
	Search string `form:"search"`
}
