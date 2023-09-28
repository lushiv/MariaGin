package auth

import "time"

// TblUsers represents the 'users' table.
type TblUsers struct {
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Type         string `json:"type"`
	Password     string `json:"password"`
	ProfilePic   string `json:"profile_pic"`
	Gender       int    `json:"gender"`
	Status       int    `json:"status"`
	Deleted      int    `json:"deleted"`
	TwoFaEnabled int    `json:"two_fa_enabled"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    int    `json:"created_by"`
	UpdatedAt    string `json:"updated_at"`
	UpdatedBy    int    `json:"updated_by"`
	DeletedAt    string `json:"deleted_at"`
	DeletedBy    int    `json:"deleted_by"`
}

// TblLoginSession represents the 'login_session' table.
type TblLoginSession struct {
	ID                     int    `json:"id"`
	UUID                   string `json:"uuid"`
	Token                  string `json:"token"`
	UserID                 int    `json:"user_id"`
	SessionExpiryTimestamp string `json:"session_expiry_timestamp"`
	CreatedAt              string `json:"created_at"`
	IP                     string `json:"ip"`
	Deleted                int    `json:"deleted"`
}

// TblValidateToken represents the 'validate_token' table.
type TblValidateToken struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	UserID     int       `json:"user_id"`
	Token      string    `json:"token"`
	Used       int       `json:"used"`
	ExpiryTime time.Time `json:"expiry_time"`
	CreatedAt  string    `json:"created_at"`
	DeletedAt  string    `json:"deleted_at"`
	Deleted    int       `json:"deleted"`
}
