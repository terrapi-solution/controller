package api

// ListResponseDto struct defines list response structure
type ListResponseDto[T any] struct {
	Data []T `json:"data"`
}
