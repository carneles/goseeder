package service

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DBEngine is the type of DB engine
type DBEngine string

const (
	// DBEnginePostgres is PostgreSQL DB engine
	DBEnginePostgres DBEngine = "postgres"
	// DBEngineMySQL is MySQL DB engine
	DBEngineMySQL DBEngine = "mysql"
)

// Seeder is the interface for Database seeder
type Seeder interface {
	Seed(SeedData) error
}

// NewSeeder will construct Database seeder service using selected DB engine
func NewSeeder(connectionString string) Seeder {
	engine := getDBEngineFromConnectionString(connectionString)
	switch engine {
	case DBEnginePostgres:
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			return nil
		}
		if err = db.Ping(); err != nil {
			return nil
		}
		return PostgresSeeder{db}
	case DBEngineMySQL:
		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			return nil
		}
		if err = db.Ping(); err != nil {
			return nil
		}
		return MySQLSeeder{db}
	default:
		return nil
	}
}

func getDBEngineFromConnectionString(connectionString string) DBEngine {
	if strings.Contains(connectionString, "postgres") {
		return DBEnginePostgres
	}
	return DBEngineMySQL
}

// SeedData represent data to be seed in YAML files
type SeedData struct {
	Schema *string       `yaml:"schema"`
	Table  string        `yaml:"table"`
	Data   []interface{} `yaml:"data"`
}
