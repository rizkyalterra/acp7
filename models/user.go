package models

import (
	"time"

	"gorm.io/gorm"
)

// db
type Users struct {
	ID        uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	Name      string         `json:"name", form:"name"`
	Email     string         `json:"email", form:"email"`
	Password  string         `json:"password", form:"password"`
	// SchoolID  uint           `json:"schoolId", form:"schoolId"`
	// School    School         `gorm:"foreignKey:ID;references:SchoolID" json:"-"`
}

type UsersDBResponse struct {
	ID        uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	Name      string         `json:"name", form:"name"`
	Email     string         `json:"email", form:"email"`
	Token     string         `json:"token", form:"token"`
	// SchoolID  uint           `json:"schoolId", form:"schoolId"`
	// School    School         `gorm:"foreignKey:ID;references:SchoolID" json:"-"`
}

type UserRequest struct {
	Name     string `json:"name", form:"name"`
	Email    string `json:"email", form:"email"`
	Password string `json:"password", form:"password"`
	// SchoolID int
}

type UserResponseSingle struct {
	Code    int             `json:"code", form:"code"`
	Message string          `json:"message", form:"message"`
	Data    UsersDBResponse `json:"data", form:"data"`
	Status  string          `json:"status", form:"status"`
}

type UserResponse struct {
	Code    int     `json:"code", form:"code"`
	Message string  `json:"message", form:"message"`
	Data    []Users `json:"data", form:"data"`
	Status  string  `json:"status", form:"status"`
}


