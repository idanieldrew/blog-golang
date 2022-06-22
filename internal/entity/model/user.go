package model

import (
	"github.com/idanieldrew/blog-golang/internal/entity/enum"
	"time"
)

type User struct {
	Id              uint      `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Password        string    `json:"password"`
	Type            enum.Type `json:"type"`
	EmailVerifiedAt string    `json:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}