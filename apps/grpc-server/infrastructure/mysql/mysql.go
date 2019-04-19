package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysqldriver
	"github.com/jmoiron/sqlx"
)

const (
	driver    = "mysql"
	charset   = "charset=utf8mb4"
	parseTime = "parseTime=true"
	loc       = "loc=Local"
)

// Open Open
func Open(conf *Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@%s/?%s&%s&%s", conf.User, conf.Password, conf.ConnectionName, charset, parseTime, loc)
	return sqlx.Open(driver, dataSourceName)
}
