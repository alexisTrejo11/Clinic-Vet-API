// Package specification defines the Specification interface for implementing the Specification pattern.
package specification

type Specification interface {
	IsSatisfiedBy(any) bool
	ToSQL() (string, []any)
}

type Pagination struct {
	Page     int
	PageSize int
	OrderBy  string
	SortDir  string // "ASC" or "DESC"
}

func (p Pagination) GetOffset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

func (p Pagination) GetLimit() int {
	return p.PageSize
}
