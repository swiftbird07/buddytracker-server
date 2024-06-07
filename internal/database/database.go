package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error

	Db, err = openConnection("database", 5432, "postgres", os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	if err != nil {
		log.Fatalln(err)
	}
}

func openConnection(host string, port uint16, user string, password string, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	setupDatabase(db)

	log.Println("database connection established")

	return db, nil
}

//go:embed migrations/*
var files embed.FS

func setupDatabase(db *sql.DB) {
	sqlFiles, err := files.ReadDir("migrations")
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(sqlFiles, func(i, j int) bool {
		return sqlFiles[i].Name() < sqlFiles[j].Name()
	})

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS migrations(filename VARCHAR(20) PRIMARY KEY, timestamp TIMESTAMPTZ DEFAULT NOW());")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("created migration table")

	for _, file := range sqlFiles {
		if !file.Type().IsDir() {
			row := db.QueryRow("SELECT filename, timestamp FROM migrations WHERE filename=?;", file.Name())
			var filename string
			var timestamp time.Time
			row.Scan(&filename, &timestamp)
			if filename != file.Name() {
				log.Println("migrating database to " + file.Name())
				sqlFile, err := files.ReadFile("migrations/" + file.Name())
				if err != nil {
					log.Fatal(err)
				}
				sqlStr := string(sqlFile)

				_, err = db.Exec(sqlStr)
				if err != nil {
					log.Fatal(err)
				} else {
					_, err = db.Exec("INSERT INTO migrations(filename) VALUES(?);", file.Name())
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	log.Println("database is ready")
}
