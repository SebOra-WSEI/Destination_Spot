package response

import (
	 "errors"
	 "github.com/gin-gonic/gin"
)

func createResponseObject(field string, res interface{}) gin.H {
	 return gin.H{field: res}
}

func CreateError(err error) gin.H {
	 return createResponseObject("error", err.Error())
}

func Create(res interface{}) gin.H {
	 return createResponseObject("response", res)
}

var (
	 ErrUserNotFound                       = errors.New("User not found")
	 ErrSpotNotFound                       = errors.New("Spot not found")
	 ErrReservationNotFound                = errors.New("Reservation not found")
	 ErrWhileUpdatingUser                  = errors.New("Error while updating user")
	 ErrWhileUpdatingReservation           = errors.New("Error while updating reservation")
	 ErrInternalServer                     = errors.New("Internal server error")
	 ErrAuthTokenNotFound                  = errors.New("Authorization token not found")
	 ErrAuthTokenIncorrectFormat           = errors.New("Incorrect authorization token format")
	 ErrTokenExpired                       = errors.New("Expired token. Please log in again")
	 ErrActionNotPermitted                 = errors.New("Action not permitted")
	 ErrEmptyFields                        = errors.New("Fields can not be empty")
	 ErrUserAlreadyExists                  = errors.New("User already exists")
	 ErrSpotAlreadyExists                  = errors.New("Spot already exists")
	 ErrInvalidEmailFormat                 = errors.New("Invalid email format")
	 ErrInvalidDomain                      = errors.New("Email should be a part of correct domain")
	 ErrMinCharacterLength                 = errors.New("Password must contains at least 8 characters")
	 ErrUppercaseCharacter                 = errors.New("Password must contain at least 1 uppercase character")
	 ErrMissingNumber                      = errors.New("Password must contain at least 1 number")
	 ErrMissingSpecialCharacter            = errors.New("Password must contain at least 1 special character")
	 ErrPasswordNotTheSame                 = errors.New("Provided passwords are not the same")
	 ErrPasswordTheSame                    = errors.New("New password cannot be the same as an actual one")
	 ErrInvalidCurrentPassword             = errors.New("Invalid current password")
	 ErrProblemWhileRegistration           = errors.New("Problem while registration")
	 ErrProblemWhileCreatingNewSpot        = errors.New("Problem while creating a new spot")
	 ErrProblemWhileCreatingNewReservation = errors.New("Problem while creating a new reservation")
	 ErrProblemWhileRemovingSpot           = errors.New("Problem while removing the spot")
	 ErrProblemWhileRemovingUser           = errors.New("Problem while removing the user")
	 ErrProblemWhileRemovingReservation    = errors.New("Problem while removing the reservation")
	 ErrInvalidLoginOrPassword             = errors.New("Invalid login or password")
	 ErrWhileCreatingToken                 = errors.New("Error while creating token")
	 ErrRequestNotExecuted                 = errors.New("Request can not be executed")
	 ErrSpotAlreadyReservedMsg             = errors.New("Spot is already reserved for selected day")
	 ErrParsingMsg                         = errors.New("Error while parsing message")
	 ErrReadingBody                        = errors.New("Error while reading body")
)

const (
	 UserCreatedMsg        string = "User created successfully!"
	 ReservationCreatedMsg        = "Reservation created successfully!"
	 ReservationUpdatedMsg        = "Reservation updated successfully!"
	 SpotCreatedMsg               = "Spot created successfully!"
	 SpotRemoveMsg                = "Spot removed successfully!"
	 ReservationRemoveMsg         = "Reservation removed successfully!"
	 UserRemoveMsg                = "User and related reservations removed successfully!"
	 PasswordChangedMsg           = "Password changed successfully!"
)
