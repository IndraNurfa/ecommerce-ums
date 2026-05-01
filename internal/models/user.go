package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID          int       `json:"-"`
	Username    string    `json:"username" gorm:"uniqueIndex;column:username;type:varchar(20);not null" validate:"required"`
	Email       string    `json:"email" gorm:"uniqueIndex;column:email;type:varchar(100);not null" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" gorm:"uniqueIndex;column:phone_number;type:varchar(15);not null" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100);not null" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password,omitempty" gorm:"column:password;type:varchar(255);not null" validate:"required"`
	Role        string    `json:"role" gorm:"column:role;type:varchar(20);not null"`
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
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt           time.Time `gorm:"index"`
	UpdatedAt           time.Time
	UserID              int       `gorm:"not null;index"`
	Token               string    `gorm:"type:char(64);not null;uniqueIndex"`
	RefreshToken        string    `gorm:"type:char(64);not null;uniqueIndex"`
	TokenExpired        time.Time `gorm:"index"`
	RefreshTokenExpired time.Time `gorm:"index"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
