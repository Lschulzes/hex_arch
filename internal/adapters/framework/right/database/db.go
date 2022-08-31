package database

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) *Adapter {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("DB connection failure: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("DB connection failure: %v", err)
	}
	return &Adapter{db: db}
}

func (a Adapter) CloseDBConnection() {
	err := a.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (a Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}
	_, err = a.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
