package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  UserRole = "manager"
	Attendee UserRole = "attendee"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"type:text; not null"`
	Email     string    `json:"email" gorm:"text; not null"`
	Role      UserRole  `json:"role" gorm:"text; default:attendee"`
	Password  string    `json:"-"`
	CreateAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Manager)
	}
	return
}
