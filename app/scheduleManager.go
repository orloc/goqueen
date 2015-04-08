package app

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/orloc/goqueen/model"
	"log"
	_ "os"
	"strconv"
)

type ScheduleManager struct {
	DbName    string
	TableName string
}

func (manager ScheduleManager) SetupDB(trucate bool) {

	db, err := sql.Open("sqlite3", "./goqueen.db?_busy_timeout=600")
	CheckErr(err)
	defer db.Close()

	sqlStmt := `
		create table if not exists schedules (
			id integer not null primary key,
			name text not null,
			mon integer not null,
			tue integer not null,
			wed integer not null,
			thu integer not null,
			fri integer not null,
			sat integer not null,
			sun integer not null,
			startTime integer not null,
			endTime integer not null
		);
		`
	if trucate {
		sqlStmt = sqlStmt + `delete from schedules;`
	}

	if _, stmtErr := db.Exec(sqlStmt); stmtErr != nil {
		log.Printf("%q: %s\n", stmtErr, sqlStmt)
		panic(stmtErr)
	}
}

func (manager ScheduleManager) DoPost(schedule *model.Schedule) {
	db, err := sql.Open("sqlite3", "./goqueen.db?_busy_timeout=600")
	CheckErr(err)
	defer db.Close()

	tx, err := db.Begin()
	CheckErr(err)
	stmt, err := tx.Prepare(fmt.Sprintf("insert into %s(name, mon, tue, wed, thu, fri, sat, sun, startTime, endTime) values(?,?,?,?,?,?,?,?,?,?)", manager.TableName))
	CheckErr(err)

	defer stmt.Close()
	_, err = stmt.Exec(schedule.Name, schedule.Mon, schedule.Tue, schedule.Wed, schedule.Thu, schedule.Fri, schedule.Sat, schedule.Sun, schedule.StartTime, schedule.EndTime)

	CheckErr(err)

	tx.Commit()

}

func (manager ScheduleManager) DoGetById(id string) model.Schedule {
	var sch model.Schedule = model.Schedule{}

	db, err := sql.Open("sqlite3", "./goqueen.db?_busy_timeout=600")
	CheckErr(err)
	defer db.Close()

	intId, err := strconv.ParseInt(id, 10, 64)
	CheckErr(err)

	rows, err := db.Query(fmt.Sprintf("select * from %s where id = %d limit 1", manager.TableName, intId))
	CheckErr(err)
	defer rows.Close()

	rows.Next()
	rows.Scan(&sch.Id, &sch.Name, &sch.Mon, &sch.Tue, &sch.Wed, &sch.Thu, &sch.Fri, &sch.Sat, &sch.Sun, &sch.StartTime, &sch.EndTime)

	return sch
}

func (manager ScheduleManager) DoGet() []model.Schedule {

	db, err := sql.Open("sqlite3", "./goqueen.db?_busy_timeout=600")
	CheckErr(err)
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("select * from %s", manager.TableName))
	CheckErr(err)

	defer rows.Close()
	var results []model.Schedule

	for rows.Next() {
		sch := new(model.Schedule)
		rows.Scan(&sch.Id, &sch.Name, &sch.Mon, &sch.Tue, &sch.Wed, &sch.Thu, &sch.Fri, &sch.Sat, &sch.Sun, &sch.StartTime, &sch.EndTime)

		results = append(results, *sch)
	}

	return results
}
