package models

import "time"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID        int        `db:"id" json:"id"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	Role      Role       `db:"role" json:"role"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	Profile   Profile    `db:"-" json:"profile"`
}

type Profile struct {
	UserID      int     `db:"user_id" json:"user_id" example:"1"`
	FirstName   *string `db:"firstname" json:"firstname" example:"Farid"`
	LastName    *string `db:"lastname" json:"lastname" example:"Darari"`
	PhoneNumber *string `db:"phone_number" json:"phone_number" example:"089876543210"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user1@gmail.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type RegisterRequest struct {
	Email       string  `json:"email" binding:"required,email" example:"newuser@mail.com"`
	Password    string  `json:"password" binding:"required" example:"mypassword"`
	Role        string  `json:"role" example:"user"`
	FirstName   *string `json:"firstname" example:"Farid"`
	LastName    *string `json:"lastname" example:"Rhamadhan"`
	PhoneNumber *string `json:"phone_number" example:"08123456789"`
}
