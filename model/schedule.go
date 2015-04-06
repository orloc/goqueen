package model

type Schedule struct {
	Id        int
	Name      string `valid:"required"`
	Mon       bool   `valid:"required"`
	Tue       bool   `valid:"required"`
	Wed       bool   `valid:"required"`
	Thu       bool   `valid:"required"`
	Fri       bool   `valid:"required"`
	Sat       bool   `valid:"required"`
	Sun       bool   `valid:"required"`
	StartTime int64  `valid:"required"`
	EndTime   int64  `valid:"required"`
}
