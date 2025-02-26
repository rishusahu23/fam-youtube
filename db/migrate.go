//go:build migrate
// +build migrate

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rishusahu23/fam-youtube/config"
)

func main() {
	// Load configuration
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Command-line argument (up or down)
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run migrate.go <up|down>")
	}

	// creating database if not exists
	if err = createDatabaseIfNotExists(conf, os.Args[2]); err != nil {
		panic(err)
	}

	// Construct PostgreSQL connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conf.PostgresConfig.User,
		conf.PostgresConfig.Password,
		conf.PostgresConfig.Host,
		conf.PostgresConfig.Port,
		os.Args[2],
		"disable",
	)

	action := os.Args[1]

	fmt.Println(os.Args)

	// Initialize migration
	migrationsDir := fmt.Sprintf("file://./db/%v/migrations", os.Args[2])
	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	switch action {
	// make upgradeDB DB_NAME={db_name} will run migrations file present in db/{db_name} directory
	case "up":
		err = m.Up()
	// make downgradeDB DB_NAME={db_name} will downgrade migrations by 1 step from files present in db/{db_name} directory
	// as of now it won't work as down files are empty
	case "down":
		err = m.Steps(-1)
	// useful when you need to unit test for dao layer
	case "snapshot":
		err = snapshotDB(dbURL, os.Args[2])
	default:
		log.Fatalf("Invalid action: %s. Use 'up' or 'down'.", action)
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration successful!")
}

// snapshotDB takes a snapshot of the current database schema and writes it to a file.
func snapshotDB(dbURL, dbName string) error {
	// Open connection to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Query to get the list of tables in the database (considering public schema and others)
	rows, err := db.Query(`
		SELECT table_schema || '.' || table_name
		FROM information_schema.tables
		WHERE table_schema NOT IN ('information_schema', 'pg_catalog')
		ORDER BY table_schema, table_name
	`)
	if err != nil {
		return fmt.Errorf("failed to fetch table list: %v", err)
	}
	defer rows.Close()

	// Start building the schema snapshot
	var schemaBuilder strings.Builder
	schemaBuilder.WriteString(fmt.Sprintf("-- Snapshot for database: %s\n", dbName))

	// Loop through the tables and fetch their DDL (CREATE TABLE statements)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return fmt.Errorf("failed to scan table name: %v", err)
		}

		// Get the DDL for each table
		var ddl string
		err := db.QueryRow(fmt.Sprintf(`
			SELECT 'CREATE TABLE ' || table_name || E'\n(\n' ||
			string_agg(column_name || ' ' || data_type, E',\n') || E'\n);\n'
			FROM information_schema.columns
			WHERE table_schema = '%s' AND table_name = '%s'
			GROUP BY table_name`,
			strings.Split(table, ".")[0], strings.Split(table, ".")[1])).Scan(&ddl)
		if err != nil {
			// If there's an error (table doesn't exist or some other issue), log it and continue
			log.Printf("Skipping table %s due to error: %v\n", table, err)
			continue
		}

		// Append the DDL to the schema builder
		schemaBuilder.WriteString(ddl + "\n\n")
	}

	// Write the schema snapshot to a file
	snapshotFile := fmt.Sprintf("db/%s/latest.sql", dbName)
	err = os.WriteFile(snapshotFile, []byte(schemaBuilder.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write schema to file: %v", err)
	}

	log.Printf("Snapshot saved to %s\n", snapshotFile)
	return nil
}

// createDatabaseIfNotExists checks if the database exists and creates it if it doesn't.
func createDatabaseIfNotExists(conf *config.Config, dbName string) error {
	// Construct PostgreSQL connection string
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conf.PostgresConfig.User,
		conf.PostgresConfig.Password,
		conf.PostgresConfig.Host,
		conf.PostgresConfig.Port,
		"postgres",
		"disable",
	)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to default database: %v", err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbName)
	var exists int
	err = db.QueryRow(query).Scan(&exists)
	if err == sql.ErrNoRows {
		// Database does not exist, create it
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
		log.Printf("Database '%s' created successfully.", dbName)
	} else if err != nil {
		return fmt.Errorf("error checking database existence: %v", err)
	} else {
		log.Printf("Database '%s' already exists.", dbName)
	}
	return nil
}
