package global

import (
	"assay/configs"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/minio/minio-go"
	"github.com/tarm/serial"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ServerConfig = &configs.ServerConfig{}
	DB           *gorm.DB
	RDB          *redis.Client
	Mqtt         mqtt.Client
	Cat          *serial.Port
	Minio        *minio.Client
)
