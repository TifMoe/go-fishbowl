package errors

// ErrorInternal contains internal definition of an error
type ErrorInternal struct {
	Status int    `json:"status,omitempty"`
	Error  *Error `json:"error,omitempty"`
}

// IsEmpty will return true if Error struct does not contain something other than default
func (e ErrorInternal) IsEmpty() bool {
	return e.Status == 0
}

// Error holds info on the error response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// IsEmpty will return true if Error struct does not contain something other than default
func (e Error) IsEmpty() bool {
	return e.Message == ""
}
