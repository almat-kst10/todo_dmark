package models

import "time"

type TodoItem struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	TimeAt      time.Time `json:"time_at"`
	IsDone     	bool      `json:"is_done"`
}
