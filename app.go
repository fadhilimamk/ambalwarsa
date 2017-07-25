package main

import (
	"github.com/fadhilimamk/ambalwarsa/src/ambalwarsa"
	"github.com/fadhilimamk/ambalwarsa/src/conf"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	filename := "./config/config.development.ini"
	if err := conf.InitConfiguration(filename); err != nil {
		log.Fatal("Error initializing ambalwarsa!")
	}

	log.WithField("Configuration", conf.Configuration).Info("Config loaded")
}

func main() {
	router := gin.Default()
	gin.SetMode(conf.Configuration.Server.GINMODE)

	router.GET("/", ambalwarsa.DefaultHandler)

	router.Run(conf.Configuration.Server.PORT)
}
