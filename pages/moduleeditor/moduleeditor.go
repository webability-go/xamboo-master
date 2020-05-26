package main

import (
	"fmt"

	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo"
	"github.com/webability-go/xamboo/assets"

	"master/app/bridge"
)

func Run(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}

	data := generateData(ctx, s.(*xamboo.Server))

	host := ctx.Request.Form.Get("host")
	app := ctx.Request.Form.Get("app")
	scontext := ctx.Request.Form.Get("context")
	module := ctx.Request.Form.Get("module")
	prefix := ctx.Request.Form.Get("prefix")

	params := &xcore.XDataset{
		"host":    host,
		"app":     app,
		"context": scontext,
		"module":  module,
		"prefix":  prefix,
		"DATA":    data,
		"#":       language,
	}
	return template.Execute(params)
}

func generateData(ctx *assets.Context, s *xamboo.Server) string {

	host := ctx.Request.Form.Get("host")
	app := ctx.Request.Form.Get("app")
	scontext := ctx.Request.Form.Get("context")
	module := ctx.Request.Form.Get("module")
	prefix := ctx.Request.Form.Get("prefix")

	return "Host:<br/>" + host + "<br/>" +
		"APP:<br/>" + app + "<br/>" +
		"Context:<br/>" + scontext + "<br/>" +
		"Module:<br/>" + module + "<br/>" +
		"Prefix:<br/>" + prefix + "<br/>"
}

func Install(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}
	srv := s.(*xamboo.Server)

	host := ctx.Request.Form.Get("host")
	app := ctx.Request.Form.Get("app")
	scontext := ctx.Request.Form.Get("context")
	module := ctx.Request.Form.Get("module")
	prefix := ctx.Request.Form.Get("prefix")

	// Get config to access things
	config := srv.GetFullConfig()

	// Extract the module interface from the APP Plugin
	for _, h := range config.Hosts {
		if h.Name != host {
			continue
		}
		for id, lib := range h.Applications {
			if id != app {
				continue
			}

			compiledmodules := lib.GetCompiledModules()

			var moduledata assets.Module = nil
			moduledata = compiledmodules.Get(module)
			if moduledata == nil {
				continue
			}

			contextcontainer := lib.GetDatasourceSet()

			contextdata := contextcontainer.GetDatasource(scontext)

			if contextdata == nil {
				continue
			}

			fmt.Println("MODULE AND CONTEXT FOUND:", moduledata, contextdata, prefix)
			// do install/update

			result, err := moduledata.Synchronize(contextdata, prefix)

			fmt.Println(err, result)
		}
	}

	return map[string]interface{}{
		"success": true, "messages": "Installed", "popup": false,
	}
}
