package service

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// MySQLSeeder is seeder for MySQL DB engine
type MySQLSeeder struct {
	db *sql.DB
}

// Seed is method for seeding the database
func (s MySQLSeeder) Seed(data SeedData) error {
	queryTemplate := "INSERT IGNORE INTO %s (%s) VALUES (%s);"

	for _, dt := range data.Data {
		if kv, ok := dt.(map[interface{}]interface{}); !ok {
			continue
		} else {
			keys, values := make([]string, 0), make([]string, 0)
			for k, v := range kv {
				keys = append(keys, k.(string))
				// check if it is a function
				value := v.(string)
				if strings.HasSuffix(value, "()") {
					values = append(values, fmt.Sprintf("%s", value))
				} else {
					values = append(values, fmt.Sprintf("'%s'", value))
				}
			}
			_, err := s.db.Exec(fmt.Sprintf(queryTemplate, data.Table, strings.Join(keys, ","), strings.Join(values, ",")))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
