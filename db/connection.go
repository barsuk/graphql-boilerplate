package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"graphql-boilerplate/configs"
	"os"
	"strconv"
	"time"
)

var Conn *sqlx.DB

func init() {
	var err error
	Conn, err = OpenSqlxViaPgxConnPool()
	if err != nil {
		// постгрес в докере ещё не запустился?
		pgWaitTime, _ := strconv.Atoi(os.Getenv("PG_WAIT_TIME"))
		fmt.Println("WAIT TIME: " + os.Getenv("PG_WAIT_TIME"))
		time.Sleep(time.Duration(pgWaitTime) * time.Second)
		Conn, err = OpenSqlxViaPgxConnPool()
		if err != nil {
			fmt.Printf(configs.ErrorColor, err)
			fmt.Println("Убиваюсь без базы-то")
			os.Exit(1)
		}
	}
	fmt.Printf(configs.InfoColor+"\n", "db started")
}

// OpenSqlxViaPgxConnPool открытие пула соединений
func OpenSqlxViaPgxConnPool() (*sqlx.DB, error) {
	port, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)
	connConfig := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"), // имя контейнера в докер-ком-пузе ???
		Port:     uint16(port),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 30,
		AfterConnect:   nil,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Call to pgx.NewConnPool failed")
	}

	nativeDB := stdlib.OpenDBFromPool(connPool)
	err = nativeDB.Ping()
	if err != nil {
		connPool.Close()
		return nil, errors.Wrap(err, "ping OpenDBFromPool failed")
	}

	return sqlx.NewDb(nativeDB, "pgx"), nil
}
