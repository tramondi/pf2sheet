package dto

type Response[T any] struct {
	Error *string `json:"error,omitempty"`
	Data  *T      `json:"data,omitempty"`
}
