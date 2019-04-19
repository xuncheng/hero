package model

type Config struct {
	MysqlHost string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	MysqlPort string `env:"MYSQL_PORT" envDefault:"3306"`
	MysqlUser string `env:"MYSQL_USER" envDefault:"root"`
	MysqlPass string `env:"MYSQL_PASS" envDefault:"admin"`
	MysqlDB   string `env:"MYSQL_DB" envDefault:"hero_dev"`
}
