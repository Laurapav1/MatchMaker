package main

import (
	"MatchMaker/routes"
	"fmt"

	"github.com/gin-gonic/gin"

	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Gender string `json:"gender"`
}

var db *sql.DB

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()
	addRoutes(r)
	r.Run()
}

func addRoutes(r *gin.Engine) {
	routes.AuthRoutes(r)
}

// initDB opretter forbindelse til SQL Server databasen
func initDB() {
	var err error
	// Opdater med dine forbindelsesdetaljer
	connString := "server=matchmaker_database;user id=sa;Password=MatchMaker!;database=master"
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}
	fmt.Println("Connected to the database successfully!")

	// Tjek om DummyTable eksisterer, og opret den, hvis ikke
	_, err = db.Exec(`IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'DummyTable')
		BEGIN
			CREATE TABLE DummyTable (
				ID INT PRIMARY KEY,
				Name NVARCHAR(50)
			);
		END`)
	if err != nil {
		log.Fatal("Error creating DummyTable: ", err.Error())
	}
	fmt.Println("Ensured DummyTable exists.")

	// Insert a dummy entry
	_, err = db.Exec("INSERT INTO DummyTable (ID, Name) VALUES (@p1, @p2)", 1, "DummyName")
	if err != nil {
		log.Fatal("Error inserting dummy entry: ", err.Error())
	}
	fmt.Println("Dummy entry inserted successfully!")
}
