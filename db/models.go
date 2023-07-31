package db

import (
	"database/sql"
	"fmt"
	"github.com/swlee90/batch-go/configuration"
	"github.com/swlee90/batch-go/logger"

	_ "github.com/lib/pq"
)

// logger

var log = logger.NewLogger()
var configs = configuration.NewDBConfig()

type Pdb struct {
	DbObj *sql.DB
	Table string
}

func DbConn() *sql.DB {
	dbinfo := fmt.Sprintf("host =%s user=%s port=%d password=%s dbname=%s sslmode=disable",
		configs.DB_URL, configs.DB_USER, configs.DB_PORT, configs.DB_PASSWORD, configs.DB_NAME)

	conn, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Error(err)
	}

	return conn
}

func (pg *Pdb) PgPing() {
	err := pg.DbObj.Ping()
	if err != nil {
		panic(err)
	}
	log.Info("DB Ping Success")
}

func (pg *Pdb) CreateTable() {
	creStr := fmt.Sprintf("CREATE TABLE %s (id serial PRIMARY KEY, name VARCHAR(20), quantity INTEGER);", pg.Table)
	_, err := pg.DbObj.Exec(creStr)
	if err != nil {
		log.Error(err)
	}
	log.Info("Finished creating table")
}

func (pg *Pdb) SelectTbl() *sql.Rows {
	stmt1 := fmt.Sprintf("SELECT * from %s;", pg.Table)
	rows, err := pg.DbObj.Query(stmt1)
	if err != nil {
		log.Error(err)
	}
	return rows
}

func (pg *Pdb) InsertTbl() {
	stmtIns := fmt.Sprintf("INSERT INTO %s (name, quantity) VALUES ($1, $2);", pg.Table)
	_, err := pg.DbObj.Exec(stmtIns, "test0", 100)
	if err != nil {
		log.Error(err)
	}
	_, err = pg.DbObj.Exec(stmtIns, "test1", 101)
	if err != nil {
		log.Error(err)
	}
	log.Info("Inserted 2 records")
}

func PrintRows(rows *sql.Rows) {
	var id int
	var name string
	var quantity int

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			log.Error("No rows were returned")
		case nil:
			log.Infof("%d, %s, %d\n", id, name, quantity)
		default:
			if err != nil {
				panic(err)
			}
		}
	}

}
