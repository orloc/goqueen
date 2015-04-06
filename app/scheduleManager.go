package app

import (
	_ "database/sql"
	_ "fmt"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/orloc/goqueen/model"
	_ "log"
	"net/http"
	_ "os"
)

type ScheduleManager struct{}

func (manager ScheduleManager) DoPost(c *echo.Context) {
	schedule := new(model.Schedule)

	if err := c.Bind(schedule); err == nil {
		sch := *schedule
		c.JSON(200, sch)
	} else {
		http.Error(c.Response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (manager ScheduleManager) DoGet(c *echo.Context) {
	sch := model.Schedule{}
	c.JSON(200, sch)
}
