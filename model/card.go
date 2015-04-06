package model

type Card struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Pin       string `json:"pin"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
