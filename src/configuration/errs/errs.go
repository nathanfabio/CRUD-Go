package errs

import "net/http"

type Errs struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error returns the error message. When you need to pass an error as a parameter
func (e *Errs) Error() string {
	return e.Message
}

// NewErrs create a new error
func NewErrs(message, err string, code int, causes []Causes) *Errs {
	return &Errs {
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

//NewBadRequestErrs create a new bad request error
func NewBadRequestErrs(message string) *Errs {
	return &Errs {
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

//NewBadRequestValidationErrs create a new bad request validation error
func NewBadRequestValidationErrs(message string, causes []Causes) *Errs {
	return &Errs{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

//NewInternalServerErrs create a new internal server error
func NewInternalServerErrs(message string) *Errs {
	return &Errs {
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

//NewNotFoundErrs create a new not found error
func NewNotFoundErrs(message string) *Errs {
	return &Errs {
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

//NewForbiddenErrs create a new forbidden error
func NewForbiddenErrs(message string) *Errs {
	return &Errs {
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}