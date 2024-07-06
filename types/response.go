package types

type Response struct {
	Error       bool   `json:"error"`
	Message     string `json:"message"`
	ErrorFields any    `json:"error_fields"`
	Data        any    `json:"data"`
}
