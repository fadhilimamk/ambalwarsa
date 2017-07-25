package main

import (
	"fmt"
	"os"

	"github.com/fadhilimamk/ambalwarsa/src/ambalwarsa"
	"github.com/fadhilimamk/ambalwarsa/src/conf"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func init() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	filename := fmt.Sprintf("./config/config.%s.ini", env)
	if err := conf.InitConfiguration(filename); err != nil {
		log.Fatal("Error initializing ambalwarsa!")
	}

	log.WithField("Configuration", conf.Configuration).Info("Config loaded")

}

func main() {

	gin.SetMode(conf.Configuration.Server.GINMODE)
	router = gin.Default()
	log.Info("Ambalwarsa is listening you ...")

	router.GET("/", ambalwarsa.DefaultHandler)

	router.Run(conf.Configuration.Server.PORT)

}
