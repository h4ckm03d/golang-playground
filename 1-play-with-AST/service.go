package main

import "time"

type Author struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Book struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	AuthorID int64  `json:"author_id,omitempty"`
	Author   Author `json:"author,omitempty"`
	ISBN     string `json:"isbn,omitempty"`
}

type Service interface {
	Get() int
}
