package model

type AuthBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordBody struct {
	CurrentPassword    string `json:"currentPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

type ActionControlBody struct {
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}
