package models

import (
	"time"
)

type Post struct {
	Name      string
	Text      string
	CreatedAt time.Time
}
