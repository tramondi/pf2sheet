package dto

type Response struct {
	Error *string `json:"error,omitempty"`
	Data  any     `json:"data,omitempty"`
}
