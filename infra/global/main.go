package global

import (
	"assay/configs"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ServerConfig = &configs.ServerConfig{}
	DB           *gorm.DB
	RDB          *redis.Client
)
