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
		UserID    uint   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
		User      UserRespon
		CommentID uint `gorm:"constraint:OnDelete:CASCADE;"`
		Comment   []CommentRespon
		CreatedAt time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:timestamp(0);default:null"`
	}
	SocialMedia struct {
		ID        uint   `gorm:"primaryKey"`
		Name      string `gorm:"not null;type:varchar(191)"`
		SosmedUrl string `gorm:"not null;type:varchar(191)"`
		UserID    uint   `gorm:"constraint:OnDelete:CASCADE;"`
		User      User
		CreatedAt time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:timestamp(0);default:null"`
	}
	Comment struct {
		ID        uint `gorm:"primaryKey"`
		UserID    uint `gorm:"constraint:OnDelete:CASCADE;"`
		User      UserRespon
		PhotoID   uint
		Message   string    `gorm:"not null;type:varchar(191)"`
		CreatedAt time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:timestamp(0);default:null"`
	}
)

type CommentRespon struct {
	ID      uint   `json:"id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

func (CommentRespon) TableName() string {
	return "comments"
}

type UserRespon struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (UserRespon) TableName() string {
	return "users"
}

type PhotoRespo struct {
	ID       uint   `json:"id"`
	PhotoUrl string `json:"photoUrl"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
}

func (PhotoRespo) TableName() string {
	return "photos"
}
