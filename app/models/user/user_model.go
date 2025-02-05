package models

import "time"

type UserModel struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"nome"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u UserModel) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Nick:      u.Nick,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
