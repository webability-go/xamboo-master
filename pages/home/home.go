package main

import (
	"github.com/webability-go/xamboo/cms"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xcore/v2"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	// If config is not already done, call install
	installed, _ := ctx.Sysparams.GetBool("installed")
	if !installed {
		return s.(*cms.CMS).Run("home/install", true, nil, "", "", "")
	}

	return s.(*cms.CMS).Run("home/home", true, nil, "", "", "")
}
