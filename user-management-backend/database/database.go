package database

import (
	"database/sql"
	"encoding/json"
	"github.com/Masterminds/squirrel"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	PSQL = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

// Config struct to hold configuration parameters
type Config struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	DbName   string `json:"db_name"`
}

func ConnectDB() {
	// Load configuration during initialization
	config, err := loadConfig("config/db_config.json")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	username := config.UserName
	//	password := config.PassWord
	dbname := config.DbName

	connStr := "user=" + username + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("Successfully connected to the database")
}

// LoadConfig reads the configuration from a JSON file
func loadConfig(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
