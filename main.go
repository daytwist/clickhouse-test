package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var ConnectionCH *sql.DB

func setDriver() error {
	var err error
	databaseURL := fmt.Sprintf("tcp://%s:%s", "clickhouse", "9000")
	ConnectionCH, err = sql.Open("clickhouse", databaseURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	if err := ConnectionCH.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			return fmt.Errorf("failed to ping database: %w", err)
		}
		return err
	}
	return nil
}

func repositoryInsertUser() error {
	var (
		tx, err = ConnectionCH.Begin()
	)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	stmt, err := tx.Prepare("INSERT INTO user (id, name ,date) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < 100; i++ {
		if _, err := stmt.Exec(
			i,
			"name",
			time.Now(),
		); err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func main() {
	if err := setDriver(); err != nil {
		log.Fatal(err)
	}
	if err := repositoryInsertUser(); err != nil {
		log.Fatal(err)
	}
}
