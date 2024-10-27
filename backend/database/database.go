package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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

	// Tjek om GameRequest-tabellen eksisterer, og opret den, hvis ikke
	_, err = DB.Exec(`
		IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'GameRequest')
		BEGIN
			CREATE TABLE GameRequest (
				ID INT PRIMARY KEY IDENTITY,
				UserEmail NVARCHAR(100),
				Niveau INT,
				Location NVARCHAR(100),
				Time DATETIME,
				Gender NVARCHAR(50),
				Amount INT,
				Price DECIMAL(10, 2)
			)
		END`)
	if err != nil {
		log.Fatal("Error creating GameRequest table: ", err.Error())
	}
	fmt.Println("Ensured GameRequest table exists.")

	// Tjek om User-tabellen eksisterer, og opret den, hvis ikk
	_, err = DB.Exec(`
		IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'User')
		BEGIN
			CREATE TABLE [User] (
				ID INT PRIMARY KEY IDENTITY,
            	FirstName NVARCHAR(100),
            	LastName NVARCHAR(100),
            	Email NVARCHAR(100) UNIQUE,
            	Password NVARCHAR(255)
			)
		END`)
	if err != nil {
		log.Fatal("Error creating User table: ", err.Error())
	}
	fmt.Println("Ensured User table exists.")

	insertDummyData()
}

func insertDummyData() {
	// Example unique constraint could be based on a combination of Niveau, Location, Time, and Gender
	userEmail := "SampleUser@email.com"
	location := "Sample Location"
	gender := "Any"
	niveau := 1
	time := time.Now()

	// Check if the record already exists
	var id int
	err := DB.QueryRow(`SELECT ID FROM GameRequest WHERE Niveau = @p1 AND Location = @p2 AND Gender = @p4`,
		niveau, location, time, gender).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		// Record does not exist, insert the dummy data
		_, err = DB.Exec(`INSERT INTO GameRequest (UserEmail, Niveau, Location, Time, Gender, Amount, Price)
		                  VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)`,
			userEmail, niveau, location, time, gender, 10, 99.99)
		if err != nil {
			log.Fatal("Error inserting dummy data: ", err.Error())
		}
		fmt.Println("Inserted dummy data into GameRequest table successfully!")
	case err != nil:
		log.Fatal("Error checking for existing record: ", err.Error())
	default:
		fmt.Println("Record already exists, skipping insert.")
	}
}
