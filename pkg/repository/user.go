package repository

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string         `gorm:"unique_index"`
	Password      string
	FollowCount   int64
	FollowerCount int64
}
