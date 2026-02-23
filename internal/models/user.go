package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username" gorm:"uniqueIndex;column:username;type:varchar(20);not null" validate:"required"`
	Email       string    `json:"email" gorm:"uniqueIndex;column:email;type:varchar(100);not null" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" gorm:"uniqueIndex;column:phone_number;type:varchar(15);not null" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100);not null" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password,omitempty" gorm:"column:password;type:varchar(255);not null" validate:"required"`
	Role        string    `json:"role,omitempty" gorm:"column:role;type:varchar(20);not null" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func (*User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  int       `gorm:"primarykey"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:text;" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:text;" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
