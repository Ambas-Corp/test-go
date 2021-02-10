package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bns-engineering/platformbanking-card/common/config"
	"github.com/bns-engineering/platformbanking-card/common/logging"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

//Engine - database engine
var Engine *xorm.Engine

//GetdbConn Function
func GetdbConn() (engine *xorm.Engine) {
	var err error
	dbUser := config.Config.Database.User
	dbPass := config.Config.Database.Pass
	dbName := config.Config.Database.Name
	dbIPAddress := config.Config.Database.Host
	dbPort, _ := strconv.Atoi(config.Config.Database.Port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbIPAddress, dbPort, dbName)
	engine, err = xorm.NewEngine("mysql", dsn)
	// engine.ShowSQL(true) // Displays the execution of SQL for easy debugging and analysis
	if err != nil {
		strlog := fmt.Sprintf("Fail to sync database: %v\n", err)
		logging.InfoLn(strlog)
		os.Exit(1)
	}
	if err := engine.Ping(); err != nil {
		logging.InfoLn(err.Error())
		os.Exit(1)
	}
	engine.SetTableMapper(core.SnakeMapper{})
	// if err = engine.Sync2(new(modelinternal.Logjnl)); err != nil {
	// 	strlog := fmt.Sprintf("Fail to sync database: %v\n", err)
	// 	logging.InfoLn(strlog)
	// }
	return engine
}

//SetEngine function
func SetEngine(engine *xorm.Engine) {
	Engine = engine
}
