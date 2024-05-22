package message

import "errors"

var (
	 ErrUserNotFound             = errors.New("User not found")
	 ErrInternalServer           = errors.New("Internal server error")
	 ErrAuthTokenNotFound        = errors.New("Authorization token not found")
	 ErrAuthTokenIncorrectFormat = errors.New("Incorrect authorization token format")
	 ErrTokenExpired             = errors.New("Expired token. Please log in again")
	 ErrActionNotPermitted       = errors.New("Action not permitted")
)
