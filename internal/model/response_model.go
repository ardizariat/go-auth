package model

type WebResponse[T any] struct {
	Data    T             `json:"data,omitempty"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  any           `json:"errors,omitempty"`
	Message string        `json:"message,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
