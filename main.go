package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/swlee90/batch-go/configuration"
	"github.com/swlee90/batch-go/db"
	"github.com/swlee90/batch-go/logger"
)

func main() {
	log := logger.NewLogger()
	log.Info("Start Batch Program")
	gocron.Every(5).Seconds().Do(minutesTask)
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
