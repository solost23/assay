package configs

type ServerConfig struct {
	Name         string      `mapstructure:"name"`
	Version      string      `mapstructure:"version"`
	Mode         string      `mapstructure:"mode"`
	Port         int         `mapstructure:"port"`
	TimeLocation string      `mapstructure:"time_location"`
	MySQL        MySQLConfig `mapstructure:"mysql"`
	Redis        RedisConfig `mapstructure:"redis"`
	JWT          JWTConfig   `mapstructure:"jwt"`
}

type MySQLConfig struct {
	DataSourceName  string `mapstructure:"dsn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Key      string `mapstructure:"key"`
	Salt     string `mapstructure:"salt"`
	Duration int64  `mapstructure:"duration"`
}
