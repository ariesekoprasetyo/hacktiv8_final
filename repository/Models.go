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
		User      User
		Comment   []CommentRespon
		CreatedAt time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:timestamp(0);default:null"`
	}
	SocialMedia struct {
		ID        uint      `gorm:"primaryKey"`
		Name      string    `gorm:"not null;type:varchar(191)"`
		SosmedUrl string    `gorm:"not null;type:varchar(191)"`
		UserID    uint      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:timestamp(0);default:null"`
	}
	Comment struct {
		ID        uint `gorm:"primaryKey"`
		UserID    uint `gorm:"foreignKey:UserID;constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
		User      UserRespon
		PhotoID   uint `gorm:"foreignKey:PhotoID;constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
		Photo     Photo
		Message   string    `gorm:"not null;type:varchar(191)"`
		CreatedAt time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:timestamp(0);default:null"`
	}
)

type CommentRespon struct {
	PhotoID uint
	Message string
	UserID  uint
	User    UserRespon
}

func (CommentRespon) TableName() string {
	return "comments"
}

type UserRespon struct {
	ID       uint
	Username string
}

func (UserRespon) TableName() string {
	return "users"
}
