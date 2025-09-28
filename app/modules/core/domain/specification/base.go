// Package specification defines the Specification interface for implementing the Specification pattern.
package specification

type Specification interface {
	IsSatisfiedBy(any) bool
	ToSQL() (string, []any)
}

type Pagination struct {
	Offset  int32
	Limit   int32
	OrderBy string
	SortDir string // "ASC" or "DESC"
}

func (p Pagination) GetOffset() int32 {
	if p.Offset <= 0 {
		return 0
	}
	return (p.Offset - 1) * p.Limit
}

func (p Pagination) GetLimit() int32 {
	return p.Limit
}
