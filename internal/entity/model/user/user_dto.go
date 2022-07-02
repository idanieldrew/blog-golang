package user

import "time"

type (
	User struct {
		Id              int64     `json:"id"`
		Name            string    `json:"name"`
		Email           string    `json:"email"`
		Phone           string    `json:"phone"`
		Password        string    `json:"password"`
		Type            int64     `json:"type"`
		EmailVerifiedAt time.Time `json:"email_verified_at"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		DeletedAt       time.Time `json:"deleted_at"`
	}

	Users []User
)
