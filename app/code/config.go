package code

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/webability-go/xamboo"
	"github.com/webability-go/xamboo/config"
)

func ReloadConfig() error {
	return xamboo.OverLoad()
}

func GetMainConfig() map[string]interface{} {

	// read main config file. The main config file is into the sysconfig "config" parameter
	filename := config.Config.File

	// load the file into dataset
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	data := map[string]interface{}{}

	json.Unmarshal(byteValue, &data)

	//	fmt.Println(data)
	return data
}
