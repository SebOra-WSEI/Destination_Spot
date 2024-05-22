package message

import (
	"errors"
)

var (
	ErrEmptyFields              = errors.New("Fields can not be empty")
	ErrUserAlreadyExists        = errors.New("User already exists")
	ErrUserNotFound             = errors.New("User not found")
	ErrWhileUpdatingUser        = errors.New("Error while updating user")
	ErrInvalidEmailFormat       = errors.New("Invalid email format")
	ErrInvalidDomain            = errors.New("Email should be a part of correct domain")
	ErrMinCharacterLength       = errors.New("Password must contains at least 8 characters")
	ErrUppercaseCharacter       = errors.New("Password must contain at least 1 uppercase character")
	ErrMissingNumber            = errors.New("Password must contain at least 1 number")
	ErrMissingSpecialCharacter  = errors.New("Password must contain at least 1 special character")
	ErrPasswordNotTheSame       = errors.New("Provided passwords are not the same")
	ErrPasswordTheSame          = errors.New("New password cannot be the same as an actual one")
	ErrInvalidCurrentPassword   = errors.New("Invalid current password")
	ErrInternalServer           = errors.New("Internal server error")
	ErrProblemWhileRegistration = errors.New("Problem while registration")
	ErrInvalidLoginOrPassword   = errors.New("Invalid login or password")
	ErrWhileCreatingToken       = errors.New("Error while creating token")
)

const (
	UserCreatedMsg     string = "User created successfully!"
	PasswordChangedMsg        = "Password changed successfully!"
)
