package domain

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Birthday int    `json:"birthday"`
	Nickname string `json:"nickname"`
	Aboutme  string `json:"aboutme"`
	Role     string `json:"role"`
}

type SignInRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SMSRequest struct {
	Phone string `json:"phone"`
}

type SMSLogin struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
