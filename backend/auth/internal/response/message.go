package response

// Error messages
const (
	EmptyFieldsErrMsg              string = "Fields can not be empty"
	UserAlreadyExistsErrMsg               = "User already exists"
	UserNotFoundErrMsg                    = "User not found"
	ErrorWhileCreatingUSerErrMsg          = "Error while creating user"
	InvalidEmailFormatErrMsg              = "Invalid email format"
	InvalidDomainErrMsg                   = "Email should be a part of correct domain"
	MinCharacterLengthErrMsg              = "Password must contains at least 8 characters"
	UppercaseCharacterErrMsg              = "Password must contain at least 1 uppercase character"
	MissingNumberErrMsg                   = "Password must contain at least 1 number"
	MissingSpecialCharacterErrMsg         = "Password must contain at least 1 special character"
	PasswordNotTheSameErrMsg              = "Provided passwords are not the same"
	PasswordTheSameErrMsg                 = "New password cannot be the same as an actual one"
	InvalidCurrentPasswordErrMsg          = "Invalid current password"
	InternalServerErrMsg                  = "Internal server error"
	ProblemWhileRegistrationErrMsg        = "Problem while registration"
	InvalidLoginOrPasswordErrMsg          = "Invalid login or password"
	ErrorWhileCreatingTokenErrMsg         = "Error while creating token"
	AuthTokenNotFoundErrMsg               = "Authorization token not found"
	AuthTokenIncorrectFormatErrMsg        = "Incorrect authorization token format"
	TokenExpiredErrMsg                    = "Expired token. Please log in again"
)

const (
	UserCreatedMsg     string = "User created successfully!"
	PasswordChangedMsg        = "Password changed successfully!"
)
