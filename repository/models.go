package repository

import "time"

type (
	User struct {
		ID          uint   `gorm:"primaryKey"`
		Username    string `gorm:"uniqueIndex"`
		Email       string `gorm:"uniqueIndex"`
		Password    string
		Age         int
		CreatedAt   time.Time
		UpdatedAt   time.Time
		SocialMedia SocialMedia `gorm:"foreignKey:ID"`
		Comment     []Comment   `gorm:"foreignKey:ID"`
	}
	Photo struct {
		ID        uint `gorm:"primaryKey"`
		Title     string
		Caption   string
		PhotoUrl  string
		UserId    uint `gorm:"foreignKey:ID"`
		User      []User
		CreatedAt time.Time
		UpdatedAt time.Time
		Comment   []Comment `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE"`
	}
	SocialMedia struct {
		ID             uint `gorm:"primaryKey"`
		Name           string
		SocialMediaUrl string
		UserId         uint `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE"`
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	Comment struct {
		ID        uint `gorm:"prima"`
		UserId    uint `gorm:"foreignKey:ID"`
		User      User
		PhotoID   uint `gorm:"foreignKey:ID"`
		Photo     Photo
		Message   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
