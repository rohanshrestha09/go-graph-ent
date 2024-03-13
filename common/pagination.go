package common

type Order string

const (
	ASC  Order = "asc"
	DESC Order = "desc"
)

type Pagination struct {
	Page  int    `form:"page,default=1"`
	Size  int    `form:"size,default=10"`
	Sort  string `form:"sort,default=id"`
	Order Order  `form:"order,default=desc"`
}
