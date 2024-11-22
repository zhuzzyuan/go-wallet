package models

type Response struct {
	Status string `json:"status"`

	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`

	Data interface{} `json:"data,omitempty"`
}
