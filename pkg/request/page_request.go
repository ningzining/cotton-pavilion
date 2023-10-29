package request

type PageParam struct {
	CurrentPage int `json:"current_page" form:"current_page"`
	PageSize    int `json:"page_size" form:"page_size"`
}
