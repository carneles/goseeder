[![Go Report Card](https://goreportcard.com/badge/github.com/carneles/goseeder)](https://goreportcard.com/report/github.com/carneles/goseeder)


# goseeder - Golang Database Seeder
This is a tool to seed data into database tables (like Laravel's `php artisan db:seed` and Node's `sequelize-cli seed`).

## Data
The data that will be inserted into tables, should be placed in YAML formatted files. One YAML file represent one table and the order of files in the folder is respected.
Goseeder uses `INSERT IGNORE` (on MySQL) and `INSERT ... ON CONFLICT DO NOTHING` (on Postgres) commands, so it is pretty safe to re-run the application again and again on tables having primary key or unique indexes.

This is the example of the YAML files:
```
schema: public
table: test
data:
  - id: 5b5d2fec-fbdb-493a-aa39-85878da3e08e
    name: one
    created_at: now()
  - id: c445876c-05bc-4479-b581-bcfd4356a9eb
    name: two
    created_at: now()
```

Store this file inside a folder, for example in `/folder/to/data`:
```
/folder
    /to
        /data
            01_first_data.yaml
            02_second_data.yaml
```

To execute:
```
goseeder seed -s /folder/to/data -d postgres://user:password@server:port/dbname?sslmode=disable
```

