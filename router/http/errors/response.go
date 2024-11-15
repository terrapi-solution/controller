package errors

// ErrorResponse is the error type for the application
type ErrorResponse struct {
	// The application error code.
	Code string `json:"code" bson:"code"`
	// A human-readable message to send back to the end user.
	Message string `json:"message" bson:"message"`
	// Defines what operation is currently being run.
	Operation string `json:"operation" bson:"op"`
	// The error that was returned from the caller.
	Err error `json:"error" bson:"error"`
}
