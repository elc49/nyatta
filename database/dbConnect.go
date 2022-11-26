package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	log "github.com/sirupsen/logrus"
)

var (
	DatabaseError = errors.New("DatabaseError")
	dbTables      = []string{
		`
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email text NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);`,
		`
CREATE TABLE IF NOT EXISTS properties (
  id BIGSERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  town varchar(50) NOT NULL,
  postal_code varchar(20) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  created_by bigint NOT NULL REFERENCES users ON DELETE CASCADE
);		`,
		`
CREATE TABLE IF NOT EXISTS amenities (
  id BIGSERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  provider varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  property_id bigint NOT NULL REFERENCES properties ON DELETE CASCADE
);
`,
	}
)

var dbClient *sql.DB

// InitDB - setup db and return connection instance/error
func InitDB() (*sql.DB, error) {
	configureDB := config.GetConfig().Database.RDBMS

	host := configureDB.Env.Host
	port := configureDB.Env.Port
	driver := configureDB.Env.Driver
	user := configureDB.Access.User
	pass := configureDB.Access.Pass
	name := configureDB.Access.DbName
	ssl_mode := configureDB.Ssl.SslMode

	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl_mode)

	db, err := sql.Open(driver, dbUri)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err == nil {
		dbClient = db
		log.Info("Database is connected")
	}

	if config.IsPrototypeEnv() {
		if err := dropAllTables(dbClient); err != nil {
			return nil, err
		}

		if err := startMigration(dbClient); err != nil {
			return nil, err
		}
	}

	return dbClient, nil
}

// GetDB - get database client
func GetDB() *sql.DB {
	return dbClient
}

// dropAllTables - cleanup database tables
func dropAllTables(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS users, properties, amenities CASCADE")
	if err == nil {
		log.Infoln("Database tables deleted")
	}
	return err
}

// startMigration - setup database tables/columns and any missing indexes
func startMigration(db *sql.DB) error {
	var err error
	for _, table := range dbTables {
		_, err = db.Exec(table)
	}
	if err == nil {
		log.Infoln("Tables migrated successfully")
	}
	return err
}
