package config

import (
	"fmt"
	"os"

	"github.com/bns-engineering/platformbanking-card/common/logging"
	"github.com/tkanos/gonfig"
)

//Config - global config
var Config *Configuration

//LoadConfig Function
func LoadConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		logging.Error(err, "fail read config")
		os.Exit(1)
	}
	return configuration
}

func getFileName() string {
	env := os.Getenv("ENV")
	if env == "DEV" || len(env) == 0 {
		env = "development"
	} else if env == "STAGING" {
		env = "staging"
	} else if env == "PRODUCTION" {
		env = "production"
	} else {
		err := fmt.Errorf("error get env")
		logging.Error(err, "please define env")
	}
	strfilename := fmt.Sprintf("config.%s.json", env)
	return strfilename
}

//SetConfig Function
func SetConfig(cfg *Configuration) {
	Config = cfg
}
