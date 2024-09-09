package web

import "time"

// APIError example
type APIError struct {
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	CreatedAt    time.Time `json:"created_at"`
}
