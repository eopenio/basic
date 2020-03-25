package basic

import (
	"github.com/eopenio/basic/config"
	"github.com/eopenio/basic/db"
	"github.com/eopenio/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
