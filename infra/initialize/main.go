package initialize

func init() {
	initLogger()
	initConfig()
	initMysql()
	initRedis()
	initMqtt()
	// initCat()
	// initMinio()
}
