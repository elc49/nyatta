package model

import "time"

type User struct {
	ID         string      `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Phone      string      `json:"phone"`
	Onboarding bool        `json:"onboarding"`
	Email      string      `json:"email"`
	Avatar     string      `json:"avatar"`
	Properties []*Property `json:"properties"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignIn struct {
	Onboarding *bool  `json:"onboarding"`
	Token      string `json:"token"`
}
