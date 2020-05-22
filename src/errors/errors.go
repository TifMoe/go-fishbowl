package errors

import (
	"net/http"
)

var (
	// ErrNewGame is thrown when service fails to instantiate a new game
	ErrNewGame = &ErrorInternal{
		Status: http.StatusInternalServerError,
		Error:  &Error{Code: 1000, Message: "Failed to instantiate new game session, please try again later"},
	}

	// ErrInvalidInput is thrown when input is invalid
	ErrInvalidInput = &ErrorInternal{
		Status: http.StatusBadRequest,
		Error:  &Error{Code: 1001, Message: "Invalid input"},
	}

	// ErrInternalError is thrown when an unknown error in the service or repository layer is thrown
	ErrInternalError = &ErrorInternal{
		Status: http.StatusInternalServerError,
		Error:  &Error{Code: 5000, Message: "Internal server error"},
	}
)
