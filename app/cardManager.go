package app

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/orloc/goqueen/model"
	_ "log"
	"strconv"
)

type CardManager struct {
	DbName    string
	TableName string
	Options   []string
}

func (manager CardManager) SetupDB(truncate bool) {

	db := manager.getHandle()
	defer db.Close()

	var card *model.Card = &model.Card{}

	if truncate {
		db.DropTable(card)
	}

	db.CreateTable(card)

}

func (manager CardManager) Save(modelMap *ModelMap) int64 {
	card := modelMap.Card

	db := manager.getHandle()
	defer db.Close()

	db.Create(card)

	return modelMap.Card.Id

}

func (manager CardManager) Update(modelMap *ModelMap, scheduleId int64) {
	card := modelMap.Card

	db := manager.getHandle()
	defer db.Close()

	db.Save(card)
}

func (manager CardManager) GetById(id string) (entity *ModelMap) {
	card := new(model.Card)

	db := manager.getHandle()
	defer db.Close()

	intId, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)

	db.First(card, intId)
	entity = manager.getResponse(*card)

	return
}

func (manager CardManager) GetAll() (results []*ModelMap) {

	db := manager.getHandle()
	defer db.Close()

	var cards []model.Card

	db.Find(&cards)

	for _, card := range cards {
		r := manager.getResponse(card)
		results = append(results, r)
	}

	return
}

func (manager CardManager) getHandle() (db gorm.DB) {
	db, err := gorm.Open("sqlite3", "./goqueen.db?_busy_timeout=600")
	CheckErr(err)

	db.LogMode(true)

	return
}
func (manager CardManager) getResponse(card model.Card) *ModelMap {
	response := new(ModelMap)
	response.Card = &card
	return response
}
