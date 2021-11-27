package model

import (
	"time"
)

// User userテーブルの情報
type User struct {
	Id     int
	Name   string
	Age    int
	Gender int
	CreatedAt time.Time
	UpdatedAt time.Time
}
