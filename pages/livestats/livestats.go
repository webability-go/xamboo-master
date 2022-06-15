package main

import (
	"github.com/webability-go/xamboo/cms/context"
	xcore "github.com/webability-go/xcore/v2"

	"masterapp/security"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := security.Verify(ctx, security.USER)
	if !ok {
		return ""
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"#": language,
	}

	return template.Execute(params)
}
