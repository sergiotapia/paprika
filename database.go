package paprika

import (
	"encoding/json"
	"fmt"
	"os"

	// Import pg to support postgres database.
	_ "github.com/lib/pq"
)

// DatabaseConfiguration is a struct that represents the database
// configuration set in conf.json.
type DatabaseConfiguration struct {
	DbName         string
	User           string
	Password       string
	Host           string
	Port           string
	SslMode        string
	ConnectTimeout string
	SslCert        string
	SslKey         string
	SslRootCert    string
}

var databaseConfiguration DatabaseConfiguration

func init() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&databaseConfiguration)
	if err != nil {
		fmt.Println("Error decoding conf.json:", err)
		os.Exit(0)
	}
}

// func initDb() *sql.DB {
//   connectionString := formatConnectionString(databaseSetting)
//   DB, err := sql.Open("postgres", connectionString)
//   if err != nil {
//     fmt.Println("Could not open connection to the database.")
//     os.Exit(0)
//   }
//   return DB
// }
//
// func formatConnectionString(databaseSetting DatabaseSetting) string {
//   // TODO: Detect what TOML settings were written in
//   // and build appropriate connection string.
//
//   // For now, let's go with the basics.
//   connectionString := fmt.Sprintf("dbname=%s host=%s port=%s sslmode=disable", databaseSetting.DbName, databaseSetting.Host, databaseSetting.Port)
//   return connectionString
// }
