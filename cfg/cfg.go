package cfg

import (
	"encoding/json"
	"os"

	"trinity/includes/Log"
)

type Config struct {
	User     string `json:"user"`
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Sslmode  string `json:"sslmode"`
}

func RowAccessString() string {
	var c Config
	fileConfig, err := os.Open("/config.json")
	if err != nil {
		Log.WriteLog(err)
	}

	defer fileConfig.Close()

	decoder := json.NewDecoder(fileConfig)
	if err := decoder.Decode(&c); err != nil {
		Log.WriteLog(err)
	}

	return "user=" + c.User + " password=" + c.Password + " dbname=" + c.Dbname + " sslmode=" + c.Sslmode
}
