package app

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/orloc/goqueen/model"
	_ "log"
	"strconv"
)

type ScheduleManager struct {
	DbName    string
	TableName string
	Options   []string
}

func (manager ScheduleManager) SetupDB(truncate bool) {

	db := manager.getHandle()
	defer db.Close()

	var schedule *model.Schedule = &model.Schedule{}

	if truncate {
		db.DropTable(schedule)
	}

	db.CreateTable(schedule)
}

func (manager ScheduleManager) Save(modelMap *ModelMap) int64 {
	schedule := modelMap.Schedule

	db := manager.getHandle()
	defer db.Close()

	db.Create(schedule)

	return modelMap.Schedule.Id

}

func (manager ScheduleManager) Update(modelMap *ModelMap, scheduleId int64) {
	schedule := modelMap.Schedule

	db := manager.getHandle()
	defer db.Close()

	db.Save(schedule)
}

func (manager ScheduleManager) GetById(id string) (entity *ModelMap) {
	sch := new(model.Schedule)

	db := manager.getHandle()
	defer db.Close()

	intId, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)

	db.First(sch, intId)
	entity = manager.getResponse(*sch)

	return
}

func (manager ScheduleManager) GetAll() (results []*ModelMap) {

	db := manager.getHandle()
	defer db.Close()

	var schedules []model.Schedule

	db.Find(&schedules)

	for _, sch := range schedules {
		r := manager.getResponse(sch)
		results = append(results, r)
	}

	return
}

func (manager ScheduleManager) getHandle() (db gorm.DB) {
	db, err := gorm.Open("sqlite3", "./goqueen.db?_busy_timeout=600")
	CheckErr(err)

	db.LogMode(true)

	return
}

func (manager ScheduleManager) getResponse(sch model.Schedule) *ModelMap {
	response := new(ModelMap)
	response.Schedule = &sch
	return response
}
