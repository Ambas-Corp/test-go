package main

import (
	"os"
	"os/signal"
	"syscall"

	cf "github.com/bns-engineering/platformbanking-card/common/config"
	"github.com/bns-engineering/platformbanking-card/common/redis"

	"github.com/bns-engineering/platformbanking-card/common/database"
	"github.com/bns-engineering/platformbanking-card/common/logging"
	"github.com/bns-engineering/platformbanking-card/handler/grpc"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logging.OpenFileLog()
	logging.InfoLn("Starting...")

	//load config
	Config := cf.LoadConfig()
	cf.SetConfig(&Config)

	//Set debug mode
	debugmode := Config.Debug
	logging.Infof("Debug Mode: %v", debugmode)
	if debugmode == true {
		logging.EnableDebugMode()
	}

	//get db connection
	getdb := database.GetdbConn()
	database.SetEngine(getdb)

	//connect redis
	redis.ConnectRedis(Config.Redis.Host, Config.Redis.Port)

	//set and start GRPC server
	GrpcPort := Config.Server.Port
	grpcServer := grpc.New()
	grpcServer.Start(GrpcPort)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	logging.InfoLn("All server stopped!")
}
