package db

import (
	"database/sql"
	"fmt"

	// "reflect"

	_ "github.com/lib/pq"
)

type DbInstance struct {
	Db *sql.DB
}

var PG *DbInstance

// func ConnectTODb() error {
// 	connStr := "user=postgres password=0549martin dbname=ayigyadb sslmode=disable"

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		fmt.Println("Something happpened", err)
// 		return err
// 	}

// 	PG = &DbInstance{Db: db}

// 	// fmt.Println("This is  the db instance%", PG.Db)

// 	return nil
// }

func ConnectTODb(dbUser, dbPassword, dbName, dbHost, dbPort, dbSslMode string) (*DbInstance,error) {
	// Construct the connection string using the provided parameters
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSslMode)

	// Open the connection to the database
	fmt.Println("connection string", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error gconnecting to the database: %v", err)
	}

	// Check if the connection is valid by pinging the database
	if err := db.Ping(); err != nil {
		 return nil,fmt.Errorf("unable to reach the database: %v", err)
	}

	// Assign the database instance to the global PG variable
	PG := &DbInstance{Db: db}
	fmt.Println("Successfully connected to the database!",dbHost,db.Ping())

	// Return nil if everything went well
	return PG, nil
}
