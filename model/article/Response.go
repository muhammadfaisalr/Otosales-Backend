package article

import (
	"time"
)

type Response struct {
	Id         float64    `json:"id"`
	Title      string     `json:"title,omitempty"`
	Content    string     `json:"content,omitempty"`
	AuthorId   float64    `json:"author_id,omitempty"`
	AuthorName string     `json:"author_name,omitempty"`
	PostDate   time.Time  `json:"post_date"`
	UpdateDate *time.Time `json:"update_date"`
}
