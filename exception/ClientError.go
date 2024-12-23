package exception

import "fmt"

type ClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (clientError *ClientError) Error() string {
	return fmt.Sprintf("Error %d: %s", clientError.Code, clientError.Message)
}

func NewClientError(code int, message string) *ClientError {
	return &ClientError{
		Code:    code,
		Message: message,
	}
}
