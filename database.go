package paprika

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type DatabaseSetting struct {
	DbName         string
	User           string
	Password       string
	Host           string
	Port           string
	SslMode        string
	ConnectTimeout string
}

var databaseSetting DatabaseSetting

func initDb() *sql.DB {
	connectionString := formatConnectionString(databaseSetting)
	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Could not open connection to the database.")
		os.Exit(0)
	}
	return DB
}

func formatConnectionString(databaseSetting DatabaseSetting) string {
	// TODO: Detect what TOML settings were written in
	// and build appropriate connection string.

	// For now, let's go with the basics.
	connectionString := fmt.Sprintf("dbname=%s host=%s port=%s sslmode=disable", databaseSetting.DbName, databaseSetting.Host, databaseSetting.Port)
	return connectionString
}

func loadDatabaseConfiguration() {
	tomlData, err := ioutil.ReadFile("database.toml")
	if err != nil {
		fmt.Println("Could not find database.toml settings file.")
		os.Exit(0)
	}

	if _, err := toml.Decode(string(tomlData), &databaseSetting); err != nil {
		fmt.Println("The database.toml file is not formatted properly.")
		os.Exit(0)
	}

	initDb()
}
