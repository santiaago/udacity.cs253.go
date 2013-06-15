package models

import (
	"time"
)

type Post struct {
	Subject string
	Content string
	Created time.Time
}
