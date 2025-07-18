package models

// ErrorResponse contain error detailled data.
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

// BadRequest is the BadRequest (400) HTTP status code
type BadRequest int

// NotFound is the NotFound (404) HTTP status code
type NotFound int

// UnprocessableEntity is the UnprocessableEntity (422) HTTP status code
type UnprocessableEntity int

// Unauthorized is the Unauthorized (401) HTTP status code
type Unauthorized int

// GetError return BadRequest wrapped error
func (BadRequest *BadRequest) GetError(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 400,
		ErrorCode:  "BadRequest",
		Message:    message,
	}
}

// GetError return Unauthorized wrapped error
func (unauthorized *Unauthorized) GetError(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 401,
		ErrorCode:  "Unauthorized",
		Message:    message,
	}
}

// GetError return NotFound wrapped error
func (NotFound *NotFound) GetError(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 404,
		ErrorCode:  "NotFound",
		Message:    message,
	}
}

// GetError return UnprocessableEntity error
func (UnprocessableEntity *UnprocessableEntity) GetError(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 422,
		ErrorCode:  "UnprocessableEntity",
		Message:    message,
	}
}
