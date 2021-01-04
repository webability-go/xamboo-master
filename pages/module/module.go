package main

import (
	"fmt"
	"strings"

	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo"
	"github.com/webability-go/xamboo/assets"
	"github.com/webability-go/xmodules/base"

	"master/app/bridge"
)

func Run(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}

	module := ctx.Request.Form.Get("module")

	data := generateData(ctx, s.(*xamboo.Server), module)

	params := &xcore.XDataset{
		"module": module,
		"DATA":   data,
		"#":      language,
	}
	return template.Execute(params)
}

func generateData(ctx *assets.Context, s *xamboo.Server, module string) string {

	config := s.GetFullConfig()
	data := "<table class=\"module-table\">"
	data += "<tr><td>Host/Applicación</td><td>Contexto</td><td>Versión compilada</td><td>Versión instalada</td><td>Prefijo</td><td>Comandos</td>"

	// 1. Get all the contexts with and without the module authorized

	// Carga los APPs Libraries de cada Host config
	for _, h := range config.Hosts {

		for id, lib := range h.Applications {

			configfile := lib.GetDatasourcesConfigFile()
			contextcontainer := lib.GetDatasourceSet()
			icompiledmodules := lib.GetCompiledModules()
			compiledmodules := icompiledmodules.(*base.Modules)

			bridge.Containers.Load(ctx, h.Name+"_"+id, configfile)
			container := bridge.Containers.GetContainer(h.Name + "_" + id)
			contexts := container.Contexts

			for _, context := range contexts {

				if context.Config == nil {
					continue // nothing to scan: only config link exists
				}
				authorizedmodules, _ := context.Config.GetStringCollection("module")
				if len(authorizedmodules) == 0 && len(*compiledmodules) == 0 {
					continue // nothing to show
				}

				var moduledata assets.Module = nil
				moduledata = compiledmodules.Get(module)
				if moduledata == nil {
					continue
				}
				thiscontext := contextcontainer.GetDatasource(context.ID)
				modversion := moduledata.GetVersion()
				installedversion := moduledata.GetInstalledVersion(thiscontext)

				for _, authmod := range authorizedmodules {
					xauthmod := strings.Split(authmod, "|")
					modid := xauthmod[0]
					modprefix := ""
					if len(xauthmod) > 1 {
						modprefix = xauthmod[1]
					}
					if modid != module {
						continue
					}
					ptr := fmt.Sprintf("%p", lib)

					data += "<tr>"
					data += "<td>" + h.Name + " / "
					data += id + " [" + ptr + "]</td>"
					data += "<td>" + context.ID + "</td>"
					data += "<td>" + modversion + "</td>"
					data += "<td>" + installedversion + "</td>"
					data += "<td>" + modprefix + "</td>"
					data += "<td>[INSTALAR] [DESINSTALAR]</td>"

					// Verify if the module is compiled/installed for this DB
					/*
						modversion := ""
						installedversion := ""
						for _, mod := range *compiledmodules {
							if modid == mod.GetID() {
								modversion = mod.GetVersion()
								if thiscontext != nil {
									installedversion = mod.GetInstalledVersion(thiscontext)
								}
								break
							}
						}
					*/

					/*
						// Get version from module table to know installed version etc
						prefix := ""
						if modprefix != "" {
							prefix = "[" + modprefix + "]"
						}

						icon := "module.png"
						status := language.Get("OK")
						version := ""
						if modversion != "" {
							version = "v" + modversion
							if installedversion == "" {
								status = language.Get("NOTINSTALLED")
								icon = "module-installable.png" // not installed
							} else if modversion != installedversion {
								status = language.Get("UPGRADE")
								icon = "module-updatable.png" // have to update
							}
						} else {
							status = language.Get("NOTCOMPILE")
							icon = "module-notcompiled.png" // not compiled
						}
					*/

					data += "</tr>"

				}
				/*
					// Now we add compiled modules not authorized
					for _, compmod := range *compiledmodules {
						found := false
						for _, authmod := range authorizedmodules {
							xauthmod := strings.Split(authmod, "|")
							if compmod.GetID() == xauthmod[0] {
								found = true
								break
							}
						}
						if found {
							continue
						}

					}
				*/
			}
		}
	}
	data += "</table>"

	return data
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
