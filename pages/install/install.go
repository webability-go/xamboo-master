package main

import (
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"
	xcore "github.com/webability-go/xcore/v2"

	"masterapp/assets"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := assets.Verify(ctx, assets.NOTINSTALLED)
	if !ok {
		return ""
	}

	// PAGE depends on COUNTRY variable (if already selected) or not
	L := ctx.Request.Form.Get("LANGUAGE")
	C := ctx.Request.Form.Get("COUNTRY")
	// verify validity of L,C

	page := "install/language"
	// If not selected, call language, or call account
	if L != "" && C != "" {
		page = "install/account"
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"VERSION":  config.Config.Version,
		"LANGUAGE": L,
		"COUNTRY":  C,
		"PAGE":     page,
		"#":        language,
	}
	return template.Execute(params)
}
