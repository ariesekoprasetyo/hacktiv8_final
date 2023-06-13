package repository

import "time"

type (
	User struct {
		ID        uint   `gorm:"primaryKey"`
		Username  string `gorm:"uniqueIndex"`
		Email     string `gorm:"uniqueIndex"`
		Password  string
		Age       int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	Photo struct {
		ID        uint   `gorm:"primaryKey"`
		Title     string `gorm:"not null;type:varchar(191)"`
		Caption   string `gorm:"not null;type:varchar(191)"`
		PhotoUrl  string `gorm:"not null;type:varchar(191)"`
		UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		User      User   `json:"user"`
		Comment   []Comment
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	SocialMedia struct {
		ID        uint   `gorm:"primaryKey"`
		Name      string `gorm:"not null;type:varchar(191)"`
		SosmedUrl string `gorm:"not null;type:varchar(191)"`
		UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	Comment struct {
		ID        uint   `gorm:"primaryKey"`
		UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		PhotoID   uint   `gorm:"not null;type:varchar(191)"`
		Message   string `gorm:"not null;type:varchar(191)"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
