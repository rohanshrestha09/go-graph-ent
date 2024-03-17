package enums

type BlogStatus string

const (
	Published   BlogStatus = "PUBLISHED"
	Unpublished BlogStatus = "UNPUBLISHED"
)

func (BlogStatus) Values() (kinds []string) {
	for _, s := range []BlogStatus{Published, Unpublished} {
		kinds = append(kinds, string(s))
	}
	return
}
