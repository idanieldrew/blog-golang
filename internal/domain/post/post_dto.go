package post

import "time"

type Post struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Details     string    `json:"details"`
	Description string    `json:"description"`
	BlueTick    bool      `json:"blue_tick"`
	UserId      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}