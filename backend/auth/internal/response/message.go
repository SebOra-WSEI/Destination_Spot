package response

const (
	EmptyFieldsErrorMsg            string = "Fields can not be empty"
	UserAlreadyExistsErrorMsg             = "User already exists"
	InvalidEmailFormatErrorMsg            = "Invalid email format"
	InvalidDomainErrorMsg                 = "Email should be a part of correct domain"
	MinCharacterLengthErrorMsg            = "Password must contains at least 8 characters"
	UppercaseCharacterErrorMsg            = "Password must contain at least 1 uppercase character"
	MissingNumberErrorMsg                 = "Password must contain at least 1 number"
	MissingSpecialCharacterErrMsg         = "Password must contain at least 1 special character"
	PasswordNotTheSameErrorMsg            = "Provided passwords are not the same"
	InternalServerErrorMsg                = "Internal server error"
	ProblemWhileRegistrationErrMsg        = "Problem while registration"
)
