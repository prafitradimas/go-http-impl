package types

type ApiResponse[T any] struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     T      `json:"result,omitempty"`
}
