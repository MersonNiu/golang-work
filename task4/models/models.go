package models

import (
	"time"
)

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Username string    `gorm:"type:varchar(50):unique;not null"`
	Password string    `gorm:"type:varchar(255);not null" `
	Email    string    `gorm:"type:varchar(100);unique;not null"`
	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(200);not null"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null:index"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null:index"`
	PostID    uint   `gorm:"not null:index"`
	User      User   `gorm:"constraint:OnDelete:SET NULL"`
	Post      Post   `gorm:"constraint:OnDelete:CASACADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
