package conf

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// Configuration : configuration object, represented INI config file.
var Configuration configuration

// Configuration : connection object to access data.
var Connection connection

type configuration struct {
	Server   server
	Database database
	CDN      cdn
}

type server struct {
	PORT        string
	ENVIRONMENT string
}

type database struct {
	MASTER_DB string
	SLAVE_DB  string
	REDIS     string
}

type cdn struct {
	IMAGE string
	VIDEO string
}

type connection struct {
	AWMasterDB *sqlx.DB
	AWSlaveDB  *sqlx.DB
	AWRedis    *redis.Pool
}
