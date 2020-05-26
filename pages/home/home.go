package main

import (
	"github.com/webability-go/xamboo"
	"github.com/webability-go/xamboo/assets"
	"github.com/webability-go/xcore/v2"
)

func Run(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	// If config is not already done, call install
	installed, _ := ctx.Sysparams.GetBool("installed")
	if !installed {
		return e.(*xamboo.Server).Run("home/install", true, nil, "", "", "")
	}

	return e.(*xamboo.Server).Run("home/home", true, nil, "", "", "")
}
