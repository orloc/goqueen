package model

type CardLog struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	ValidPin  string `json:"valid_pin"`
	CreatedAt string `json:"created_at"`
}
