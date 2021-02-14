package bridge

import (
	"errors"
	"fmt"
	"io/ioutil"
	"plugin"

	"github.com/webability-go/xamboo/cms/context"
)

var GetMainConfig func() map[string]interface{}

func LinkConfig(lib *plugin.Plugin) error {

	fct, err := lib.Lookup("GetMainConfig")
	if err != nil {
		fmt.Println(err)
		return errors.New("ERROR: THE APPLICATION LIBRARY DOES NOT CONTAIN GetMainConfig FUNCTION")
	}
	GetMainConfig = fct.(func() map[string]interface{})

	return nil
}

func GenerateConfig(ctx *context.Context, L string, C string, serial string, username string, password string, email string) {

	md5password := ""
	if password != "" {
		md5password = GetMD5Hash(password)
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
