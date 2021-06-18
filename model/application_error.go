package model

// ApplicationError --
type ApplicationError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}
