package db

import (
	"fmt"

	"os"
	// "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func InitpgDb() (*DbInstance,error) {

	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, fmt.Errorf("error loading .env file %v",err)
	// }

	// Get the database connection parameters from the environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")

	DbInstance,error1:=ConnectTODb(dbUser,dbPassword,dbName,dbHost,dbPort,dbSslMode)
	if error1 != nil {
		// If there was an error, log it and exit
		return nil, fmt.Errorf("error connecting nto the database: %v", error1)
		
	}

	return DbInstance, nil

	// fmt.Println("This is the PG assessing", PG.Db)

	// If no error, use the db connection (example)

	// fmt.Println("Connected to the database successfully!",value.Db)

	// Don't forget to close the database when don

}
