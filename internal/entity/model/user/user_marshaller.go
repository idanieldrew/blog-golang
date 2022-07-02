package user

import (
	"encoding/json"
	"time"
)

type (
	PublicUser struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	PrivateUser struct {
		Id              int64     `json:"id"`
		Name            string    `json:"name"`
		Email           string    `json:"email"`
		Type            int64     `json:"type"`
		UpdatedAt       time.Time `json:"updated_at"`
		EmailVerifiedAt time.Time `json:"email_verified_at"`
	}
)

func (u Users) Marshal(public bool) []interface{} {
	res := make([]interface{}, len(u))
	for i, v := range u {
		res[i] = v.Marshal(public)
	}
	return res
}

func (u *User) Marshal(public bool) interface{} {
	if !public {
		return PublicUser{
			Id:    u.Id,
			Name:  u.Name,
			Email: u.Email,
		}
	}

	var privateUser PrivateUser
	userJson, _ := json.Marshal(u)
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return err
	}
	return privateUser
}
