package config

type ServerConfig struct {
	Port int
	Name string
}

type MysqlConfig struct {
	Url      string // 数据库连接地址
	Username string // 数据库用户名
	Password string
	Database string
}
