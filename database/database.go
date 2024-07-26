package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

// initDB opretter forbindelse til SQL Server databasen
func InitDB() {
	var err error
	// Opdater med dine forbindelsesdetaljer
	connString := "server=matchmaker_database;user id=sa;Password=MatchMaker!;database=master"
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening database: ", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}
	fmt.Println("Connected to the database successfully!")

		// Tjek om CreateSearch-tabellen eksisterer, og opret den, hvis ikke
		_, err = DB.Exec(`
		IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'Search')
		BEGIN
			CREATE TABLE Search (
				ID INT PRIMARY KEY IDENTITY,
				Niveau INT,
				Location NVARCHAR(100),
				Time DATETIME,
				Gender NVARCHAR(50),
				Amount INT,
				Price DECIMAL(10, 2)
			)
		END`)
		if err != nil {
			log.Fatal("Error creating Search table: ", err.Error())
		}
		fmt.Println("Ensured Search table exists.")
}
