package initialize

func init() {
	initLogger()
	initConfig()
	initMysql()
	initRedis()
}
