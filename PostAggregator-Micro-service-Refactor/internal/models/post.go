package models

import (
	"sync"
	"time"
)

type Post struct {
	PostId    int
	UserId    int
	TimeStamp time.Time
	Content   string
	mu        sync.RWMutex
}
