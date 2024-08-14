package database

import (
	"MatchMaker/models"
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

	insertDummyData()
}

func insertDummyData() {
	// Example unique constraint could be based on a combination of Niveau, Location, Time, and Gender
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
		_, err = DB.Exec(`INSERT INTO GameRequest (Niveau, Location, Time, Gender, Amount, Price)
		                  VALUES (@p1, @p2, @p3, @p4, @p5, @p6)`,
			niveau, location, time, gender, 10, 99.99)
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

func GetAllGameRequests() ([]models.GameRequest, error) {
	rows, err := DB.Query("SELECT ID, Niveau, Location, Time, Gender, Amount, Price FROM GameRequest")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GameRequest
	for rows.Next() {
		var cs models.GameRequest
		if err := rows.Scan(&cs.ID, &cs.Niveau, &cs.Location, &cs.Time, &cs.Gender, &cs.Amount, &cs.Price); err != nil {
			return nil, err
		}
		results = append(results, cs)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}
