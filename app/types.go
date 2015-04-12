package app

import (
	model "github.com/orloc/goqueen/model"
)

type ModelMap struct {
	Schedule *model.Schedule
	Card     *model.Card
}

type ModelManager interface {
	SetupDB(bool)
	Save(*ModelMap) int64
	Update(*ModelMap, int64)
	GetAll() (results []*ModelMap)
	GetById(string) (model *ModelMap)
}
