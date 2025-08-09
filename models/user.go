package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null" form:"username" json:"username" binding:"required"`
	Password string    `gorm:"not null" form:"password" json:"password" binding:"required"`
	Email    string    `gorm:"unique;not null" form:"email" json:"email" binding:"required,email"`
	Posts    []Post    `form:"-" json:"posts,omitempty"`
	Comments []Comment `form:"-" json:"comments,omitempty"`
}
