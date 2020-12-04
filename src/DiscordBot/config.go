package discordbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadConfig(filename string, bot *BotConfig) (err error) {
	fmt.Printf("Reading file %s...\n", filename)

	file, er := ioutil.ReadFile(filename)
	if er != nil {
		err = fmt.Errorf("ReadConfig ReadFile: %s => %s", filename, er)
		return
	}

	// fmt.Println(string(file))

	er = json.Unmarshal(file, bot)
	if er != nil {
		err = fmt.Errorf("ReadConfig Unmarshal: %s", er)
		return
	}

	return
}
