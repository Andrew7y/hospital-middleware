package query

type Pagination struct {
	Page     int
	PageSize int
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

type Sorting struct {
	SortBy    string
	SortOrder string
}
