package models

import (
	"time"
)

type Post struct {
	Id int64 
	Subject string
	Content string
	Created time.Time
}
