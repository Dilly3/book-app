package util

type ErrorResponse struct {
	Code    int         `json:"code"`
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}
type SuccessResponse struct {
	Code    int         `json:"code"`
	Object  interface{} `json:"object"`
	Message interface{} `json:"message"`
}
