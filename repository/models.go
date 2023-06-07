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
		Photo       []Photo       `gorm:"foreignKey:ID"`
		SocialMedia []SocialMedia `gorm:"foreignKey:ID"`
		Comment     []Comment     `gorm:"foreignKey:ID"`
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

//
//type Orders struct {
//	OrderId      uint      `gorm:"primaryKey" json:"order_id"`
//	CustomerName string    `json:"customer_name"`
//	Ordered_at   time.Time `json:"ordered_at"`
//	Created_at   time.Time `gorm:"autoCreateTime" json:"created_at"`
//	Updated_at   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
//	Items        []Items   `gorm:"foreignKey:OrderId" json:"order"`
//}
//
//type Items struct {
//	ItemId      uint      `gorm:"primaryKey" json:"item_id"`
//	ItemCode    string    `json:"item_code"`
//	Description string    `json:"description"`
//	Quantity    int       `json:"quantity"`
//	OrderId     uint      `gorm:"foreignKey:OrderId" json:"order_id"`
//	Created_at  time.Time `gorm:"autoCreateTime" json:"created_at"`
//	Updated_at  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
//}
