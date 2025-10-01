package database

import (
	"api-e-ticketing/src/models"
	"api-e-ticketing/src/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable to hold the database connection

func DatabaseInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't read environment")
	}

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign the connection to the global variable
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err)
	}

	fmt.Println("Successfully connected to the database.")
}

func CreateDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't read environment")
	}

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsnRoot := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s",
		dbHost, dbPort, dbUser, dbPassword,
	)

	DB, err = gorm.Open(postgres.Open(dsnRoot), &gorm.Config{}) // Assign the connection to the global variable
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	createCmd := fmt.Sprintf("CREATE DATABASE %s", dbName)
	if err := DB.Exec(createCmd).Error; err != nil {
		// Jika error karena sudah ada, bisa diabaikan, tapi handle kemungkinan error lain
		log.Println("Warning during create database:", err)
	} else {
		fmt.Println("Successfully create database.")
	}


	dsnDB := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	DB, err = gorm.Open(postgres.Open(dsnDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the target database:", err)
	}

	// Enable citext extension
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS citext;").Error; err != nil {
		log.Fatal("Failed to enable citext extension:", err)
	} else {
		fmt.Println("citext extension enabled ✅")
	}
}

func DropTables() {
	err := DB.Migrator().DropTable(
		&models.User{},
		&models.Trip{},
		&models.Transaction{},
		&models.Terminal{},
		&models.Node{},
		&models.Distance{},
		&models.Balance{},
	)
	if err != nil {
		log.Fatal("Failed to drop tables...")
	}

	fmt.Println("Dropped tables successfully")
}

func Migration() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Trip{},
		&models.Transaction{},
		&models.Terminal{},
		&models.Node{},
		&models.Distance{},
		&models.Balance{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}

func UserSeeder() {
	// Create a new user
	hashPassword, err := utils.HashPassword("password")
	if err != nil {
		log.Fatal("Failed to hash password")
	}

	//Admin
	atmin := models.User{
		Username: "admin",
		Email:    "admintest@mail.com",
		Password: hashPassword,
		Role: models.ADMIN,
	}

	err = DB.Create(&atmin).Error
	if err != nil {
		log.Fatal("Failed to seed user...")
	}

	user := models.User{
		Username: "userbiasa",
		Email:    "usertest@mail.com",
		Password: hashPassword,
		Role: models.USER,
	}

	err = DB.Create(&user).Error
	if err != nil {
		log.Fatal("Failed to seed user...")
	}
}