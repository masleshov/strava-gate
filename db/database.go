package db

import (
	"database/sql"
	"log"

	"bitbucket.org/virtualtrainer/strava-gate/config"
	_ "github.com/lib/pq"
)

type statementType int

const (
	oneRow statementType = iota
	manyRow
)

// Database is representation of configuration of database
type Database struct {
	ConnectionString string
}

// QueryArgs represents array of QueryArg
type QueryArgs []interface{}

// IsNilOrEmpty checks that QueryArgs is not nil and contains at least one argument of query which is not nil
func (args QueryArgs) IsNilOrEmpty() bool {
	length := len(args)
	return length == 0 || (length == 1 && args[0] == nil)
}

// ExecSelect executes a select query with some arguments which returns many rows
func (dbase *Database) ExecSelect(query string, args QueryArgs) (*sql.Rows, error) {
	rows, err := dbase.execQuery(manyRow, query, args)
	return rows.(*sql.Rows), err
}

// ExecCRUD executes some DML request to database
func (dbase *Database) ExecCRUD(query string, args QueryArgs) (*sql.Row, error) {
	row, err := dbase.execQuery(oneRow, query, args)
	return row.(*sql.Row), err
}

// NewDatabase creates new instance of database type with fileds from config
func NewDatabase() Database {
	db := &Database{
		ConnectionString: config.Vars.DatabaseURL,
	}
	return *db
}

func (dbase *Database) establishConn() *sql.DB {
	database, err := sql.Open("postgres", dbase.ConnectionString)
	if err != nil {
		log.Fatalf("Can't create database connection: " + err.Error())
	}

	return database
}

func (dbase *Database) close(db *sql.DB) {
	db.Close()
}

func (dbase *Database) execQuery(stmnt statementType, query string, args QueryArgs) (interface{}, error) {
	db := dbase.establishConn()

	var (
		rows interface{}
		err  error
	)

	if args.IsNilOrEmpty() && stmnt == manyRow {
		rows, err = db.Query(query)
	} else if !args.IsNilOrEmpty() && stmnt == manyRow {
		rows, err = db.Query(query, args...)
	} else if args.IsNilOrEmpty() && stmnt == oneRow {
		rows = db.QueryRow(query)
	} else if !args.IsNilOrEmpty() && stmnt == oneRow {
		rows = db.QueryRow(query, args...)
	}

	dbase.close(db)

	return rows, err
}
