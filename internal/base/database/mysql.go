package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST" envDefault:"localhost:6379"`
	Port     string `env:"MYSQL_PORT" envDefault:"3306"`
	DBName   string `env:"MYSQL_DBNAME" envDefault:""`
	Username string `env:"MYSQL_USERNAME" envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD" envDefault:"root"`
}

func NewMySQlConnection(cfg *MySQLConfig) (*sql.DB, error) {
	lcfg := toLibConfig(cfg)
	conn, err := mysql.NewConnector(lcfg)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(conn)
	return db, nil
}

func toLibConfig(cfg *MySQLConfig) *mysql.Config {
	return &mysql.Config{
		Addr:      fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Net:       "tcp",
		DBName:    cfg.DBName,
		User:      cfg.Username,
		Passwd:    cfg.Password,
		ParseTime: true,
		Loc:       time.Local,
	}
}
