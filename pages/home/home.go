package main

import (
	"github.com/webability-go/xamboo/cms"
	"github.com/webability-go/xamboo/cms/context"
	xcore "github.com/webability-go/xcore/v2"

	"masterapp/assets"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	// If config is not already done, call install
	installed := assets.IsInstalled(ctx)
	if !installed {
		return s.(*cms.CMS).Run("home/install", true, nil, "", "", "")
	}

	return s.(*cms.CMS).Run("home/home", true, nil, "", "", "")
}
