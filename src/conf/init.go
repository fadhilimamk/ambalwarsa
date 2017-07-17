package conf

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	log "github.com/prometheus/common/log"
	gcfg "gopkg.in/gcfg.v1"
)

// InitConfiguration : Read configuration file for Ambalwarsa.
func InitConfiguration(filename string) error {
	log.Info("Init Ambalwarsa with config file from :", filename)

	err := gcfg.ReadFileInto(&Configuration, filename)
	if err != nil {
		log.Fatal("Init config fail :", err)
		return err
	}

	return nil
}

// InitConnection : open connection based on configuration
func InitConnection() error {
	var err error

	// Prepare connection for master database
	masterDB, err := sqlx.Open("postgres", Configuration.Database.MASTER_DB)
	if err != nil {
		log.Fatal("Error opening master database :", err)
		return err
	}
	Connection.AWMasterDB = masterDB
	_, err = Connection.AWMasterDB.Query("SELECT 1")
	if err != nil {
		log.Fatal("Error accessing master database :", err)
		return err
	}

	// Prepare connection for slave database
	slaveDB, err := sqlx.Open("postgres", Configuration.Database.SLAVE_DB)
	if err != nil {
		log.Fatal("Error opening slave database :", err)
		return err
	}
	Connection.AWSlaveDB = slaveDB
	_, err = Connection.AWSlaveDB.Query("SELECT 1")
	if err != nil {
		log.Fatal("Error accessing slave database :", err)
		return err
	}

	// Prepare connection for redis pool
	redisPool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", Configuration.Database.REDIS) },
	}
	Connection.AWRedis = redisPool

	return nil
}
