package dto

type Register struct {
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	Email            string `json:"email" binding:"required"`
	UserName         string `json:"user_name" binding:"required"`
	Password         string `json:"password" binding:"required"`
	PasswordConfirm  string `json:"password_confirm" binding:"required"`
	VerificationCode string `json:"verification_code,omitempty"`
}

type Login struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"required"`
}

type ResetPassword struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordConfirm struct {
	VerificationCode string `json:"verification_code,omitempty"`
	Password         string `json:"password" binding:"required"`
	PasswordConfirm  string `json:"password_confirm" binding:"required"`
}
