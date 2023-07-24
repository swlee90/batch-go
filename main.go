package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/sw90lee/batch-sample/configuration"
	"github.com/sw90lee/batch-sample/db"
	"github.com/sw90lee/batch-sample/logger"
)

func main() {
	log := logger.NewLogger()
	log.Info("Start Batch Program")
	gocron.Every(1).Minutes().Do(minutesTask)
	<-gocron.Start()

	s := gocron.NewScheduler()
	<-s.Start()

}

func minutesTask() {
	cfg := configuration.NewDBConfig()
	conn := db.Pdb{Table: cfg.DB_TABLE, DbObj: db.DbConn()}
	defer conn.DbObj.Close()
	rows := conn.SelectTbl()
	db.PrintRows(rows)
}

func test() {
	fmt.Println("123")
}
