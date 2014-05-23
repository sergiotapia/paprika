package paprika

import (
	"fmt"
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

var database_setting DatabaseSetting

func loadDatabaseConfiguration() {
	tomlData, err := ioutil.ReadFile("database.toml")
	if err != nil {
		fmt.Println("Could not find database.toml settings file.")
		os.Exit(0)
	}

	if _, err := toml.Decode(string(tomlData), &database_setting); err != nil {
		fmt.Println("The database.toml file is not formatted properly.")
		os.Exit(0)
	}
}
