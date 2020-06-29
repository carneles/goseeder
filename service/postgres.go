package service

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// PostgresSeeder is seeder for PostgreSQL DB engine
type PostgresSeeder struct {
	db *sql.DB
}

// Seed is method for seeding the database
func (s PostgresSeeder) Seed(data SeedData) error {
	queryTemplate := "INSERT INTO %s.%s (%s) VALUES (%s) ON CONFLICT DO NOTHING;"

	schema := "public"
	if data.Schema != nil {
		schema = *data.Schema
	}
	for _, dt := range data.Data {
		if kv, ok := dt.(map[interface{}]interface{}); !ok {
			continue
		} else {
			keys, values := make([]string, 0), make([]string, 0)
			for k, v := range kv {
				keys = append(keys, k.(string))
				value := v.(string)
				// check if it is a function
				if strings.HasSuffix(value, "()") {
					values = append(values, fmt.Sprintf("%s", value))
				} else {
					values = append(values, fmt.Sprintf("'%s'", value))
				}
			}
			_, err := s.db.Exec(fmt.Sprintf(queryTemplate, schema, data.Table, strings.Join(keys, ","), strings.Join(values, ",")))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
