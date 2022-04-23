package postgre

import (
	"fmt"
	"time"

	"github.com/cicingik/check-out/config"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jinzhu/gorm"
)

type (
	DbEngine struct {
		G   *gorm.DB
		cfg config.AppConfig
	}
)

func NewDbService(cfg config.AppConfig) *DbEngine {

	db, err := initDb(cfg)
	if err != nil {
		fmt.Printf("Cannot connect to %s database. Details %v", cfg.DBConfig.DbDriver, err)
	}

	return &DbEngine{
		G:   db,
		cfg: cfg,
	}
}

func initDb(cfg config.AppConfig) (connection *gorm.DB, err error) {

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.DBConfig.DbHost, cfg.DBConfig.DbPort, cfg.DBConfig.DbUser, cfg.DBConfig.DbName, cfg.DBConfig.DbPassword)

	// Register augments the provided driver with tracing, enabling it to be loaded by gormtrace.Open.
	sqltrace.Register(cfg.DBConfig.DbDriver, &pq.Driver{}, sqltrace.WithServiceName(config.AppName))

	// Open the registered driver, allowing all uses of the returned *gorm.DB to be traced.
	connection, err = gormtrace.Open(cfg.DBConfig.DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database. Details %v", cfg.DBConfig.DbDriver, err)
		return nil, err
	} else {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		connection.DB().SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		connection.DB().SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		connection.DB().SetConnMaxLifetime(time.Hour)

		fmt.Printf("Connected to the %s database", cfg.DBConfig.DbDriver)
		connection.LogMode(cfg.DBConfig.DbDebug == 1)

		connection.DB().SetMaxIdleConns(cfg.DBConfig.MaxIdleConns)
		connection.DB().SetMaxOpenConns(cfg.DBConfig.MaxOpenConns)
		connection.DB().SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxConnLifetimeSeconds) * time.Second)

		//defer connection.Close()
		return connection, err
	}
}
