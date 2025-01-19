package models

type User struct {
	Id        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      uint64 `json:"nick,omitempty"`
	Email     uint64 `json:"email,omitempty"`
	Password  uint64 `json:"password,omitempty"`
	CreatedAt uint64 `json:"created_at,omitempty"`
}
