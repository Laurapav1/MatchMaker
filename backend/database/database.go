package database

import (
	"MatchMaker/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB sets up the connection to the SQL Server database and initializes the tables.
func InitDB() {
	var err error
	connString := "sqlserver://sa:MatchMaker!@matchmaker_database:1433?database=master"

	// Open the database connection using GORM with SQL Server driver
	DB, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err.Error())
	}

	// Auto migrate the models to create tables if they do not exist
	err = DB.AutoMigrate(&models.User{}, &models.GameRequest{})
	if err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

	fmt.Println("Connected to the database and ensured tables exist.")

	// Insert dummy data if necessary
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
	var existingGameRequest models.GameRequest
	err := DB.Where("niveau = ? AND location = ? AND time = ? AND gender = ?", niveau, location, time, gender).First(&existingGameRequest).Error

	if err == gorm.ErrRecordNotFound {
		// Record does not exist, insert the dummy data
		dummyGameRequest := models.GameRequest{
			UserEmail: userEmail,
			Niveau:    niveau,
			Location:  location,
			Time:      time,
			Gender:    gender,
			Amount:    10,
			Price:     99.99,
		}

		if err := DB.Create(&dummyGameRequest).Error; err != nil {
			log.Fatal("Error inserting dummy data: ", err.Error())
		}
		fmt.Println("Inserted dummy data into GameRequest table successfully!")
	} else if err != nil {
		log.Fatal("Error checking for existing record: ", err.Error())
	} else {
		fmt.Println("Record already exists, skipping insert.")
	}
}
