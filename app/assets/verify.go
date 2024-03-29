package assets

import (
	"net/http"

	"github.com/webability-go/xamboo/cms/context"
)

func Verify(ctx *context.Context, connection int) bool {

	// is appname is empty: search for "app" entry in ctx
	appname, _ := ctx.Sysparams.GetString("masterapp")
	if appname == "" {
		http.Error(ctx.Writer, "Master Library name is not available in config file (parameter masterapp missing)", http.StatusInternalServerError)
		return false
	}

	installed, _ := ctx.Sysparams.GetBool("installed")
	switch connection {
	case NOTINSTALLED:
		if installed {
			http.Error(ctx.Writer, "Error: the system is already installed.", http.StatusUnauthorized)
			return false
		}
	case ANY:
		if !installed {
			http.Error(ctx.Writer, "Error: the system is not installed.", http.StatusUnauthorized)
			return false
		}
	case USER:
		if !installed {
			http.Error(ctx.Writer, "Error: the system is not installed.", http.StatusUnauthorized)
			return false
		}
		sessionid, _ := ctx.Sessionparams.GetString("sessionid")
		if sessionid == "" {
			http.Error(ctx.Writer, "Error: the user is not connected.", http.StatusUnauthorized)
			return false
		}
	}
	return true
}

func IsInstalled(ctx *context.Context) bool {
	installed, _ := ctx.Sysparams.GetBool("installed")
	return installed
}
