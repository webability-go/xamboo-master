package code

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/webability-go/xamboo"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"

	"github.com/webability-go/xmodules/tools"
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

func GenerateConfig(ctx *context.Context, L string, C string, serial string, username string, password string, email string) {

	md5password := ""
	if password != "" {
		md5password = tools.GetMD5Hash(password)
	} else {
		md5password, _ = ctx.Sysparams.GetString("password")
	}

	local := "username=" + username + "\n"
	local += "password=" + md5password + "\n"
	local += "email=" + email + "\n"
	local += "language=" + L + "\n"
	local += "country=" + C + "\n"
	local += "serial=\"" + serial + "\n"
	local += "installed=yes\n"

	// write local
	resourcesdir, _ := ctx.Sysparams.GetString("resourcesdir")
	path := resourcesdir + "/local.conf"
	ioutil.WriteFile(path, []byte(local), 0644)

	// inject into Host.config
	ctx.Sysparams.LoadString(local)
}
