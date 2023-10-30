package response

type PageData struct {
	Items any   `json:"items"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}
