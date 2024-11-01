package filter

// queryParams represents the query parameters for filtering and pagination.
type queryParams struct {
	Search         string   `form:"search"`
	Filter         []string `form:"filter"`
	Page           int      `form:"page,default=1"`
	PageSize       int      `form:"page_size,default=10"`
	OrderBy        string   `form:"order_by,default=id"`
	OrderDirection string   `form:"order_direction,default=desc,oneof=desc asc"`
}
