package app

import (
	"database/sql"
	_ "fmt"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/orloc/goqueen/model"
	"log"
)

type CardManager struct {
	DbName    string
	TableName string
	Options   []string
}

func (manager CardManager) SetupDB(truncate bool) {

	db, err := sql.Open("sqlite3", "./goqueen.db?_busy_tmeout=600")
	CheckErr(err)
	defer db.Close()

	sqlStmt := `
		create table if not exists ` + manager.TableName + ` (
			id integer not null primaray key, 
			name text not null
			code string not null
			pin string not null
			isActive numeric not null
			createdAt integer not null
			updatedAt integer not null
		);
	`

	if truncate {
		sqlStmt = sqlStmt + `delete from ` + manager.TableName + `;`
	}

	if _, stmtErr := db.Exec(sqlStmt); stmtErr != nil {
		log.Printf("%q: %s\n", stmtErr, sqlStmt)
		panic(stmtErr)
	}
}

func (manager CardManager) Save(card *ModelMap) int64 {

	return 3849
}

func (manager CardManager) Update(card *ModelMap, id int64) {

}

func (manager CardManager) GetAll() (results []*ModelMap) {
	m := new(ModelMap)
	results = append(results, m)
	return
}

func (manager CardManager) GetById(id string) (model *ModelMap) {
	return
}

func (manager CardManager) getResponse(card model.Card) *ModelMap {
	response := new(ModelMap)
	response.Card = &card
	return response
}
