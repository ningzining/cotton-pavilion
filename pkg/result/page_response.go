package result

type PageData struct {
	Items       any   `json:"items"`
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	Total       int64 `json:"total"`
}
