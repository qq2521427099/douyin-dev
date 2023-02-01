package repository

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int64 `gorm:"AUTO_INCREMENT"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Name          string `gorm:"INDEX"`
	Password      string
	FollowCount   int64
	FollowerCount int64
}
