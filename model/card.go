package model

type Card struct {
	Id        int64    `valid:"required"`
	Name      string   `valid:"required"`
	Code      string   `valid:"required"`
	Pin       string   `valid:"required"`
	IsActive  bool     `valid:"required"`
	Scheudle  Schedule `valid:"required"`
	CreatedAt string   `valid:"required"`
	UpdatedAt string   `valid:"required"`
}
