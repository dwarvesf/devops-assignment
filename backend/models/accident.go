package models

import (
	"time"
)

type Accident struct {
	ID      string    `json:"id"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
