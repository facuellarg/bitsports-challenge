package dataprovider

import (
	"bitsports/ent"
	"bitsports/utils"
	"errors"
	"fmt"
	"strconv"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type DataSource struct {
	Host     string
	User     string
	Password string
	DataBase string
	Port     int
}

//errors
var (
	ErrInvalidPort = errors.New("invalid port")
)

var (
	envDatasource DataSource
)

func init() {
	if err := LoadDataSource(); err != nil {
		log.Warn(err.Error())
	}
}

func LoadDataSource() error {
	envDatasource.DataBase = utils.GetEnvOrDefault("DB_NAME", "bitsport")
	envDatasource.Host = utils.GetEnvOrDefault("DB_HOST", "127.0.0.1")
	envDatasource.Password = utils.GetEnvOrDefault("DB_PASS", "1234")
	envDatasource.User = utils.GetEnvOrDefault("DB_USER", "root")
	portInt, err := strconv.Atoi(utils.GetEnvOrDefault("DB_PORT", "5432"))
	if err != nil {
		return ErrInvalidPort
	}
	envDatasource.Port = portInt
	return nil
}

func SetDataSource(newDataSource DataSource) {
	envDatasource = newDataSource
}

func GetPostgressUri() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		envDatasource.Host,
		envDatasource.Port,
		envDatasource.User,
		envDatasource.Password,
		envDatasource.DataBase,
	)
}

func GetClient(driverName, uri string) (*ent.Client, error) {
	return ent.Open(driverName, uri)
}

func GetPostgresClient() (*ent.Client, error) {
	return GetClient(dialect.Postgres, GetPostgressUri())

}

func GetSqliteClient() (*ent.Client, error) {
	return GetClient(dialect.SQLite, "file:./ent.sql?mode=memory&cache=shared&_fk=1")
}
