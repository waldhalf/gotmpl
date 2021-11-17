package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute


// Creates database pool for pgsql
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxIdleTime(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	// Ping the DB
	err = testDb(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// NewDatabase create a new database for the application
func NewDatabase(dsn string)(*sql.DB, error){
	// Open the DB
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// testDb tries to ping the db
func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}