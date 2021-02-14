package main

import (
	"fmt"
	"strings"

	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xmodules/base"

	"master/app/bridge"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}

	module := ctx.Request.Form.Get("module")

	data := generateData(ctx, s.(*cms.CMS), module)

	params := &xcore.XDataset{
		"module": module,
		"DATA":   data,
		"#":      language,
	}
	return template.Execute(params)
}

type cc struct {
	Ptr              string
	ID               string
	Ctxid            string
	Modversion       string
	Installedversion string
	Prefix           string
}

func generateData(ctx *context.Context, s *cms.CMS, module string) string {

	config := s.GetFullConfig()
	data := "<table class=\"module-table\">"
	data += "<tr><td>Applicación</td><td>Contexto</td><td>Versión compilada</td><td>Versión instalada</td><td>Prefijo</td><td>Comandos</td>"

	// 1. Get all the contexts with and without the module authorized
	contextcontainerslist := map[string]cc{}

	// Carga los APPs Libraries de cada Host config
	for _, h := range config.Hosts {

		for _, plg := range h.Plugins {
			lib := applications.GetApplication(plg.Id)

			configfile := lib.GetDatasourcesConfigFile()

			contextcontainer := lib.GetDatasourceSet()
			icompiledmodules := lib.GetCompiledModules()
			compiledmodules := icompiledmodules.(*base.Modules)

			bridge.Containers.Load(ctx, h.Name+"_"+plg.Name, configfile)
			container := bridge.Containers.GetContainer(h.Name + "_" + plg.Name)
			contexts := container.Contexts

			for _, context := range contexts {
				if context.Config == nil {
					continue // nothing to scan: only config link exists
				}
				authorizedmodules, _ := context.Config.GetStringCollection("module")
				if len(authorizedmodules) == 0 && len(*compiledmodules) == 0 {
					continue // nothing to show
				}

				var moduledata applications.Module = nil
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

					contextcontainerslist[ptr+context.ID] = cc{
						Ptr:              ptr,
						ID:               plg.Name,
						Ctxid:            context.ID,
						Modversion:       modversion,
						Installedversion: installedversion,
						Prefix:           modprefix,
					}

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

	for _, cc := range contextcontainerslist {
		data += "<tr>"
		data += "<td>" + cc.ID + " [" + cc.Ptr + "]</td>"
		data += "<td>" + cc.Ctxid + "</td>"
		data += "<td>" + cc.Modversion + "</td>"
		data += "<td>" + cc.Installedversion + "</td>"
		data += "<td>" + cc.Prefix + "</td>"

		data += "<td>"
		if cc.Installedversion == "" {
			data += "[<span style=\"background-color: #dfd; cursor: pointer;\" onclick=\"install('" + cc.Ptr + "', '" + cc.Ctxid + "', '" + module + "', '" + cc.Prefix + "');\">INSTALAR</span>]"
		} else if cc.Installedversion != cc.Modversion {
			data += "[<span style=\"background-color: #ddf; cursor: pointer;\" onclick=\"upgrade('" + cc.Ptr + "', '" + cc.Ctxid + "', '" + module + "', '" + cc.Prefix + "');\">ACTUALIZAR</span>]"
		}
		if cc.Installedversion != "" {
			data += "[<span style=\"background-color: #ddf; cursor: pointer;\" onclick=\"reinstall('" + cc.Ptr + "', '" + cc.Ctxid + "', '" + module + "', '" + cc.Prefix + "');\">REINSTALAR</span>]"
			data += "[<span style=\"background-color: #fdd; cursor: pointer;\" onclick=\"uninstall('" + cc.Ptr + "', '" + cc.Ctxid + "', '" + module + "', '" + cc.Prefix + "');\">DESINSTALAR</span>]"
		}
		data += "</td>"

	}

	data += "</table>"

	return data
}

func Install(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}
	srv := s.(*cms.CMS)

	app := ctx.Request.Form.Get("app")
	scontext := ctx.Request.Form.Get("context")
	module := ctx.Request.Form.Get("module")
	prefix := ctx.Request.Form.Get("prefix")

	// Get config to access things
	config := srv.GetFullConfig()
	result := []string{}
	//	var err error

	// Extract the module interface from the APP Plugin
	done := false
	var err error
	for _, h := range config.Hosts {
		for _, plg := range h.Plugins {
			lib := applications.GetApplication(plg.Id)
			ptr := fmt.Sprintf("%p", lib)
			if ptr != app {
				continue
			}

			compiledmodules := lib.GetCompiledModules()

			var moduledata applications.Module = nil
			moduledata = compiledmodules.Get(module)
			if moduledata == nil {
				continue
			}

			datasourceset := lib.GetDatasourceSet()

			datasource := datasourceset.GetDatasource(scontext)

			if datasource == nil {
				continue
			}

			//			fmt.Println("MODULE AND CONTEXT FOUND:", moduledata, contextdata, prefix)
			// do install/update
			result, err = moduledata.Synchronize(datasource, prefix)
			done = true

			//			fmt.Println(err, result)
			break
		}
		if done {
			break
		}
	}

	success := true
	message := "Installed"
	if err != nil {
		success = false
		message = "Error"
	}

	return map[string]interface{}{
		"success": success, "messages": message, "popup": false, "result": result, "error": err,
	}
}
