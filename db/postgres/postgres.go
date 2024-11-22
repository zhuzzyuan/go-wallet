package postgres

import (
	"database/sql"
	"strings"
	"time"

	"go-wallet/config"
	"go-wallet/db/reconnector"
	"go-wallet/util/log"

	// Import postgres driver
	_ "github.com/lib/pq"
)

// DB encapsulates PostgreSQL DB variables.
type DB struct {
	db *sql.DB
}

var (
	dbClient *DB

	autoReconnectDisabled = false
)

// Init opens db connection and checks if the connection is valid.
func Init() {
	connStr := config.GetPostgresConnectionURL()
	// log.Infof("Connect to db: %s", config.GetDBInfo())

	db, err := sql.Open("postgres", connStr)
	if err == nil {
		err = db.Ping()
		if err == nil {
			log.Infof("Connected to db: %s", connStr)
		}
	}

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	dbClient = &DB{db}
}

// DisableAutoReconnect disables auto-reconnect feature.
func DisableAutoReconnect() {
	autoReconnectDisabled = true
}

func (db *DB) reconnect() {
	reconnector.Reconnect("Postgres", func() bool {
		return db.db.Ping() == nil && serverAlive()
	})
}

func (db *DB) lostConnection(err error) bool {
	errMsg := err.Error()
	if strings.HasSuffix(errMsg, "Server shutdown in progress") ||
		strings.HasSuffix(errMsg, "invalid connection") ||
		strings.HasSuffix(errMsg, "Too many connections") ||
		strings.HasPrefix(errMsg, "database is not available") ||
		strings.HasSuffix(errMsg, "no such host") {
		log.Debug("server shutdown or connection invalid")
		return true
	}

	return db.db.Ping() != nil || !serverAlive()
}

func serverAlive() bool {
	var t time.Time
	query := "SELECT CURRENT_TIMESTAMP"
	err := QueryRow(query, nil, &t)
	if err != nil {
		log.Error(err)
	}

	return t.Unix() > 0 && err == nil
}

func dbReady() {
	if dbClient == nil {
		log.Fatal("Cannot execute sql: please init db client first")
	}
}
