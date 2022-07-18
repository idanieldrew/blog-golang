package post

import (
	"encoding/json"
	"log"
	"time"
)

type (
	PublicPost struct {
		Id          int       `json:"id"`
		Title       string    `json:"title"`
		Slug        string    `json:"slug"`
		Details     string    `json:"details"`
		Description string    `json:"description"`
		BlueTick    bool      `json:"blue_tick"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	PrivatePost struct {
		Id          int       `json:"id"`
		Title       string    `json:"title"`
		Slug        string    `json:"slug"`
		Details     string    `json:"details"`
		Description string    `json:"description"`
		BlueTick    bool      `json:"blue_tick"`
		UserId      int       `json:"user_id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func (p Posts) Marshal(public bool) interface{} {
	res := make([]interface{}, len(p))
	for i, v := range p {
		res[i] = v.Marshal(public)
	}
	return res
}

func (p *Post) Marshal(public bool) interface{} {
	if !public {
		var publicPost PublicPost
		b, me := json.Marshal(p)
		if me != nil {
			log.Print(me)
		}

		if ue := json.Unmarshal(b, &publicPost); ue != nil {
			log.Print(ue)
		}
		return publicPost
	}

	return PrivatePost{
		Id:          p.Id,
		Title:       p.Title,
		Slug:        p.Slug,
		Details:     p.Details,
		Description: p.Description,
		BlueTick:    p.BlueTick,
		UpdatedAt:   p.UpdatedAt,
		UserId:      p.UserId,
		CreatedAt:   p.CreatedAt,
	}
}
