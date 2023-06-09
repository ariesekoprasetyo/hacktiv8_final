package repository

import "time"

type (
	User struct {
		ID          uint   `gorm:"primaryKey"`
		Username    string `gorm:"uniqueIndex"`
		Email       string `gorm:"uniqueIndex"`
		Password    string
		Age         uint
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Photo       []Photo     `gorm:"foreignKey:ID"`
		SocialMedia SocialMedia `gorm:"foreignKey:ID"`
		Comment     []Comment   `gorm:"foreignKey:ID"`
	}
	Photo struct {
		ID        uint `gorm:"primaryKey"`
		Title     string
		Caption   string
		PhotoUrl  string
		UserId    uint `gorm:"foreignKey:ID"`
		CreatedAt time.Time
		UpdatedAt time.Time
		Comment   []Comment `gorm:"foreignKey:ID"`
	}
	SocialMedia struct {
		ID             uint `gorm:"primaryKey"`
		Name           string
		SocialMediaUrl string
		UserId         uint `gorm:"foreignKey:ID"`
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	Comment struct {
		ID        uint `gorm:"prima"`
		UserId    uint `gorm:"foreignKey:ID"`
		PhotoID   uint `gorm:"foreignKey:ID"`
		Message   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
