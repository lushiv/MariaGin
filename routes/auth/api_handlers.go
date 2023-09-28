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
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
