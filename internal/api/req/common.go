package req

type PageInfo struct {
	Name       string `form:"name"`     // 分页查询的name
	Page       int    `form:"page"`     // 分页查询的页码
	PageSize   int    `form:"pageSize"` // 分页查询的页容量
	Type       int    `form:"type"`
	CategoryId uint64 `form:"categoryId"`
	Status     int    `form:"status"`
}
