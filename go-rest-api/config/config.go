package config

import (
	"encoding/json"
	"io"
	"os"
)

type config struct {
	InitialAmount int `json:"initialBalanceAmount"`
	MinimumAmount int `json:"minimumBalanceAmount"`
}

var cfg = &config{}

func init() {
	os.Chdir("..")
	file, err := os.Open(".config/" + Env() + ".json")
	checkError(err)
	defer file.Close()

	read, err := io.ReadAll(file)
	checkError(err)

	err = json.Unmarshal(read, cfg)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetConfig() *config {
	return cfg
}
