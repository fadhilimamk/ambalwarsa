package main

import (
	"encoding/json"
	"net/http"

	"github.com/fadhilimamk/ambalwarsa/src/conf"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	filename := "./config/config.development.ini"
	err := conf.InitConfiguration(filename)
	if err != nil {
		log.Fatal("Error initializing ambalwarsa!")
	}

	log.WithField("Configuration", conf.Configuration).Info("Config loaded")

}

func main() {
	router := gin.Default()
	gin.SetMode("test")
	router.GET("/", mainHandler)

	router.Run(conf.Configuration.Server.PORT)
}

// Handler Section --------------------------------------
func mainHandler(c *gin.Context) {
	writer := c.Writer

	response := Response{
		Data: "mantap",
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
