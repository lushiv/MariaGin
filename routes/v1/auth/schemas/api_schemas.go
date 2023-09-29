package auth

// RegisterUserReq represents the request payload for user registration.
type RegisterUserReq struct {
	Email      string `json:"email" binding:"required"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password" binding:"required"`
	Phone      string `json:"phone"`
	ProfilePic string `json:"profile_pic"`
	Gender     int    `json:"gender"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginUserReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

type CommonResponse struct {
	Message string `json:"message"`
}
