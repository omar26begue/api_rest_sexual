package models

type HTTPResponse struct {
	Error   bool   `json:"error,omitempty"`
	Code    int    `json:"code,omitempty" example:"400"`
	Message string `json:"message,omitempty" example:"status bad request"`
}
