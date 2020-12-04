package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BotConfig struct {
	Token		string `json:"Token"`
	BotPrefix	string `json:"BotPrefix"`
}

func ReadConfig(filename string) (botConfig *BotConfig, err error) {
	fmt.Printf("Reading file %s...\n", filename)

	file, er := ioutil.ReadFile(filename)
	if er != nil {
		err = fmt.Errorf("ReadConfig ReadFile: %s => %s", filename, er)
		return
	}

	fmt.Println(string(file))

	er = json.Unmarshal(file, botConfig)
	if er != nil {
		err = fmt.Errorf("ReadConfig Unmarshal: %s", er)
		return
	}

	return
}
