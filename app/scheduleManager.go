package app

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/orloc/goqueen/model"
	"log"
	_ "os"
)

type ScheduleManager struct {
	DbName    string
	TableName string
}

func (manager ScheduleManager) SetupDB(trucate bool) {

	db, err := sql.Open("sqlite3", "./goqueen.db")
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
	db, err := sql.Open("sqlite3", "./goqueen.db")
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

func (manager ScheduleManager) DoGet() map[string]model.Schedule {
	var results map[string]model.Schedule

	db, err := sql.Open("sqlite3", "./goqueen.db")
	CheckErr(err)
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("select * from %s", manager.TableName))
	CheckErr(err)

	defer rows.Close()
	// NOTE not implemented yet
	/*
		for rows.Next() {

		}
	*/

	return results
}
