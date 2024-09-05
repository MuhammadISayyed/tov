package models

import "time"

type Guide struct {
	ID int
	Title string
	Content string
	CreatedAt time.Time
}