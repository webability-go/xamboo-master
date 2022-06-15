package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms"
	"github.com/webability-go/xamboo/cms/context"
	xcore "github.com/webability-go/xcore/v2"

	"github.com/webability-go/xmodules/base"

	"masterapp/code"
	"masterapp/security"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

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

func Menu(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := security.Verify(ctx, security.USER)
	if !ok {
		return ""
	}

	Order := ctx.Request.Form.Get("Order")

	if Order == "get" {
		return getMenu(ctx, s.(*cms.CMS), language)
	}
	if Order == "getchildren" {
		Father := ctx.Request.Form.Get("father")
		switch Father {
		case "files":
			return getFiles(ctx, s.(*cms.CMS), language)
		case "services":
			return getServices(ctx, s.(*cms.CMS), language)
		case "listeners":
			return getListeners(ctx, s.(*cms.CMS), language)
		case "hosts":
			return getHosts(ctx, s.(*cms.CMS), language)
		case "engines":
			return getEngines(ctx, s.(*cms.CMS), language)
		case "datasources":
			return getDatasources(ctx, s.(*cms.CMS), language)
		case "datacourcecontainers":
			return getDatasourceContainers(ctx, s.(*cms.CMS), language)
		case "modules":
			return getModules(ctx, s.(*cms.CMS), language)
		default:
			if Father[0:4] == "ctx-" { // modules of context
				return getModulesOfContainer(ctx, s.(*cms.CMS), language, Father[4:])
			}
		}
		return getMenu(ctx, s.(*cms.CMS), language)
	}
	if Order == "openclose" {

		//    $id = $this->base->HTTPRequest->getParameter('id');
		//    $status = $this->base->HTTPRequest->getParameter('status');
		//    $this->base->security->setParameter('mastertree', $id, $status=='true'?1:0);
	}

	return map[string]string{
		"status": "OK",
	}
}

func getMenu(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	// Main Menu Title and reload button
	optr := map[string]interface{}{
		"id":        "maintitle",
		"template":  "maintitle",
		"loadable":  false,
		"closeable": false,
		"closed":    false,
	}
	rows = append(rows, optr)

	// stats Title
	optr = map[string]interface{}{
		"id":        "statstitle",
		"template":  "title",
		"loadable":  false,
		"loaded":    false,
		"closeable": false,
		"closed":    false,
		"title":     language.Get("STATS.TITLE"),
	}
	rows = append(rows, optr)

	// TODO(phil) Put here all the stats options
	// live stats
	optr = map[string]interface{}{
		"id":        "livestats",
		"template":  "livestats",
		"loadable":  false,
		"closeable": false,
		"closed":    false,
		"title":     language.Get("LIVESTATS.TITLE"),
	}
	rows = append(rows, optr)

	// System
	optr = map[string]interface{}{
		"id":        "systemtitle",
		"template":  "title",
		"loadable":  false,
		"closeable": false,
		"closed":    false,
		"title":     language.Get("SYSTEM.TITLE"),
	}
	rows = append(rows, optr)

	// 0. main config line
	optr = map[string]interface{}{
		"id":        "systemreload",
		"template":  "systemreload",
		"loadable":  false,
		"closeable": false,
		"closed":    false,
		"title":     language.Get("SYSTEMRELOAD.TITLE"),
	}
	rows = append(rows, optr)

	// config title
	optr = map[string]interface{}{
		"id":        "configtitle",
		"template":  "title",
		"loadable":  false,
		"closeable": false,
		"closed":    false,
		"title":     language.Get("CONFIG.TITLE"),
	}
	rows = append(rows, optr)

	// 0. main config line
	optr = map[string]interface{}{
		"id":        "mainconfig",
		"template":  "mainconfig",
		"loadable":  false,
		"closeable": false,
		"closed":    false,
		"title":     language.Get("MAINCONFIG.TITLE"),
	}
	rows = append(rows, optr)

	// 1. Files
	optr = map[string]interface{}{
		"id":        "files",
		"template":  "files",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	// 2. Services
	optr = map[string]interface{}{
		"id":        "services",
		"template":  "services",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	// 3. Contexts
	optr = map[string]interface{}{
		"id":        "datasources",
		"template":  "datasources",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	// 4. Contexts containers
	optr = map[string]interface{}{
		"id":        "datasourcecontainers",
		"template":  "datasourcecontainers",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	// 5. Modules
	optr = map[string]interface{}{
		"id":        "modules",
		"template":  "modules",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	/*

		config := s.GetFullConfig()


		// Containers de los hosts x APPs
		optr = map[string]interface{}{
			"id":        "containers",
			"template":  "containers",
			"loadable":  false,
			"closeable": true,
			"closed":    true,
		}
		rows = append(rows, optr)

		apps := map[string]int{}
		cnt := 1

		// Carga los APPs Libraries de cada Host config
		for _, h := range config.Hosts {
			for id, lib := range h.Applications {

				ptr := fmt.Sprintf("%p", lib)
				num, ok := apps[ptr]
				if !ok {
					num = cnt
					apps[ptr] = cnt
					cnt++
				}

				opt := map[string]interface{}{
					"id":        "cnt-" + h.Name + "-" + id,
					"hostid":    h.Name,
					"appid":     id,
					"template":  "container",
					"name":      h.Name + "::" + id,
					"status":    "[#" + fmt.Sprint(num) + "]",
					"father":    "containers",
					"loadable":  false,
					"closeable": true,
					"closed":    true,
				}
				rows = append(rows, opt)

				configfile := lib.GetDatasourcesConfigFile()
				contextcontainer := lib.GetDatasourceSet()
				icompiledmodules := lib.GetCompiledModules()
				compiledmodules := icompiledmodules.(*base.Modules)

				//			fmt.Println(configfile)
				//			fmt.Println(contextcontainer)
				//			fmt.Println(compiledmodules)

				// Lets load the configuration structure of the context
				bridge.Containers.Load(ctx, h.Name+"_"+id, configfile)
				container := bridge.Containers.GetContainer(h.Name + "_" + id)
				contexts := container.Contexts

				//			fmt.Println(contexts)

				for _, context := range contexts {
					thiscontext := contextcontainer.GetDatasource(context.ID)
					icon := "context.png"
					if thiscontext == nil {
						icon = "context-notavailable.png"
					}

					opt := map[string]interface{}{
						"id":        "ctx-" + h.Name + "-" + id + "-" + context.ID,
						"hostid":    h.Name,
						"appid":     id,
						"conid":     context.ID,
						"template":  "context",
						"icon":      icon,
						"name":      context.ID,
						"color":     "black",
						"father":    "cnt-" + h.Name + "-" + id,
						"loadable":  false,
						"closeable": false,
					}
					rows = append(rows, opt)

					if context.Config == nil {
						continue // nothing to scan: only config link exists
					}
					authorizedmodules, _ := context.Config.GetStringCollection("module")
					if len(authorizedmodules) == 0 && len(*compiledmodules) == 0 {
						continue // nothing to show
					}

					for _, authmod := range authorizedmodules {
						xauthmod := strings.Split(authmod, "/")
						modid := xauthmod[0]
						modprefix := ""
						if len(xauthmod) > 1 {
							modprefix = xauthmod[1]
						}

						// Verify if the module is compiled/installed for this DB
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

						opt := map[string]interface{}{
							"id":        "mod-" + h.Name + "-" + id + "-" + context.ID + "-" + modprefix + modid,
							"icon":      icon,
							"hostid":    h.Name,
							"appid":     id,
							"conid":     context.ID,
							"modid":     modid,
							"modprefix": modprefix,
							"template":  "module",
							"name":      prefix + modid + " " + version,
							"color":     "black",
							"status":    status,
							"father":    "ctx-" + h.Name + "-" + id + "-" + context.ID,
							"loadable":  false,
							"closeable": false,
						}
						rows = append(rows, opt)
					}

					// Now we add compiled modules not authorized
					for _, compmod := range *compiledmodules {
						found := false
						for _, authmod := range authorizedmodules {
							xauthmod := strings.Split(authmod, ""/"")
							if compmod.GetID() == xauthmod[0] {
								found = true
								break
							}
						}
						if found {
							continue
						}

						opt := map[string]interface{}{
							"id":        "mod-" + h.Name + "-" + id + "-" + context.ID + "-" + compmod.GetID(),
							"hostid":    h.Name,
							"appid":     id,
							"modid":     compmod.GetID(),
							"modprefix": "",
							"template":  "module",
							"icon":      "module-noaction.png",
							"name":      compmod.GetID() + " " + compmod.GetVersion(),
							"color":     "black",
							"status":    "(NOT AUTHORIZED)",
							"father":    "ctx-" + h.Name + "-" + id + "-" + context.ID,
							"loadable":  false,
							"closeable": false,
						}
						rows = append(rows, opt)
					}
				}
			}
		}
	*/
	return map[string]interface{}{
		"row": rows,
	}

}

func getFiles(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	cfg := code.GetMainConfig()
	ifiles := cfg["include"]

	for _, file := range ifiles.([]interface{}) {
		optr := map[string]interface{}{
			"id":        file.(string),
			"file":      file.(string),
			"template":  "file",
			"father":    "files",
			"loadable":  false,
			"closeable": false,
			"name":      file.(string),
		}
		rows = append(rows, optr)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

func getServices(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	//   listeners
	optr := map[string]interface{}{
		"id":        "listeners",
		"template":  "listeners",
		"father":    "services",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	//   Hosts
	optr = map[string]interface{}{
		"id":        "hosts",
		"template":  "hosts",
		"father":    "services",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	//   Engines
	optr = map[string]interface{}{
		"id":        "engines",
		"template":  "engines",
		"father":    "services",
		"loadable":  true,
		"loaded":    false,
		"closeable": true,
		"closed":    true,
	}
	rows = append(rows, optr)

	return map[string]interface{}{
		"row": rows,
	}

}

func getListeners(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	config := s.GetFullConfig()

	for _, l := range config.Listeners {
		ip := "*"
		if l.IP != "" {
			ip = l.IP
		}
		opt := map[string]interface{}{
			"id":        "lis-" + l.Name,
			"lisid":     l.Name,
			"template":  "listener",
			"name":      l.Name + " [" + l.Protocol + "://" + ip + ":" + l.Port + "]",
			"father":    "listeners",
			"loadable":  false,
			"closeable": false,
		}
		rows = append(rows, opt)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

func getHosts(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	config := s.GetFullConfig()

	for _, h := range config.Hosts {
		hn := ""
		if len(h.HostNames) > 0 {
			hn = h.HostNames[0]
		}
		opt := map[string]interface{}{
			"id":        "hos-" + h.Name,
			"hosid":     h.Name,
			"template":  "host",
			"name":      h.Name + " [" + hn + "]",
			"father":    "hosts",
			"loadable":  false,
			"closeable": false,
		}
		rows = append(rows, opt)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

func getEngines(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	config := s.GetFullConfig()

	for _, e := range config.Engines {
		opt := map[string]interface{}{
			"id":        "eng-" + e.Name,
			"engid":     e.Name,
			"template":  "engine",
			"name":      e.Name,
			"father":    "engines",
			"loadable":  false,
			"closeable": false,
		}
		rows = append(rows, opt)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

type ct struct {
	Ptr string
	ID  string
}

func getDatasources(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	for cntid, ct := range *base.Containers {
		dss := ct.GetDatasources()
		for dtsid := range dss {
			opt := map[string]interface{}{
				"id":        "ds-" + cntid + "-" + dtsid,
				"cntid":     cntid,
				"dtsid":     dtsid,
				"template":  "datasource",
				"name":      cntid + "::" + dtsid,
				"father":    "datasources",
				"loadable":  true,
				"loaded":    false,
				"closeable": true,
				"closed":    true,
			}
			rows = append(rows, opt)
		}
	}

	return map[string]interface{}{
		"row": rows,
	}

}

type cc struct {
	Ptr        string
	ID         string
	ConfigFile string
}

func getDatasourceContainers(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	config := s.GetFullConfig()

	contextcontainerslist := map[string]cc{}
	// TODO(phil) Add stats fo # used in applications

	for _, h := range config.Hosts {
		for _, plg := range h.Plugins {

			lib := applications.GetApplication(plg.Id)
			ptr := fmt.Sprintf("%p", lib)
			configfile := lib.GetDatasourcesConfigFile()

			contextcontainerslist[ptr] = cc{
				Ptr:        ptr,
				ID:         plg.Name,
				ConfigFile: configfile,
			}
		}
	}

	for _, cc := range contextcontainerslist {
		opt := map[string]interface{}{
			"id":        "ctxcnt-" + cc.Ptr,
			"template":  "datasourcecontainer",
			"name":      cc.ID + " [" + cc.Ptr + "] " + cc.ConfigFile,
			"father":    "datasourcecontainers",
			"loadable":  false,
			"closeable": false,
		}
		rows = append(rows, opt)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

func getModules(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage) map[string]interface{} {

	rows := []interface{}{}

	config := s.GetFullConfig()

	modules := map[string]applications.Module{}
	keys := []string{}

	// Carga los APPs Libraries de cada Host config
	for _, h := range config.Hosts {
		for _, plg := range h.Plugins {

			// TODO(phil) add and calculate how many modules are authorized and installed
			application := applications.GetApplication(plg.Id)
			icompiledmodules := application.GetCompiledModules()
			compiledmodules := icompiledmodules.(*base.Modules)

			for _, mod := range *compiledmodules {
				modid := mod.GetID()
				modversion := mod.GetVersion()
				modules[modid+"/"+modversion] = mod
			}
		}
	}

	for k := range modules {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, mod := range keys {

		opt := map[string]interface{}{
			"id":        "mod-" + modules[mod].GetID(),
			"modid":     modules[mod].GetID(),
			"template":  "module",
			"icon":      "module.png",
			"name":      modules[mod].GetID() + " " + modules[mod].GetVersion(),
			"father":    "modules",
			"loadable":  false,
			"closeable": false,
		}
		rows = append(rows, opt)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

type md struct {
	Ptr              string
	ID               string
	Ctxid            string
	Modversion       string
	Installedversion string
	Prefix           string
}

func getModulesOfContainer(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage, container string) map[string]interface{} {

	rows := []interface{}{}

	config := s.GetFullConfig()
	cnt := strings.Split(container, "/")

	moduleslist := map[string]md{}

	// Carga los APPs Libraries de cada Host config
	for _, h := range config.Hosts {

		for _, plg := range h.Plugins {

			lib := applications.GetApplication(plg.Id)
			ptr := fmt.Sprintf("%p", lib)
			if ptr != cnt[0] {
				continue
			}

			configfile := lib.GetDatasourcesConfigFile()
			//			contextcontainer := lib.GetDatasourceSet()
			icompiledmodules := lib.GetCompiledModules()
			compiledmodules := icompiledmodules.(*base.Modules)

			code.Containers.Load(ctx, h.Name+"_"+plg.Name, configfile)
			container := code.Containers.GetContainer(h.Name + "_" + plg.Name)
			contexts := container.Datasources

			for _, context := range contexts {

				if context.Config == nil {
					continue // nothing to scan: only config link exists
				}

				if context.ID != cnt[1] {
					continue
				}

				authorizedmodules, _ := context.Config.GetStringCollection("module")
				if len(authorizedmodules) == 0 && len(*compiledmodules) == 0 {
					continue // nothing to show
				}

				for _, authmod := range authorizedmodules {
					xauthmod := strings.Split(authmod, "|")
					modid := xauthmod[0]
					modprefix := ""
					if len(xauthmod) > 1 {
						modprefix = xauthmod[1]
					}

					moduleslist[modid+"/"+modprefix] = md{
						Ptr:              ptr,
						ID:               modid,
						Ctxid:            context.ID,
						Modversion:       "",
						Installedversion: "",
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

				}
				/*
					// Now we add compiled modules not authorized
					for _, compmod := range *compiledmodules {
						found := false
						for _, authmod := range authorizedmodules {
							xauthmod := strings.Split(authmod, "/")
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

	for _, cc := range moduleslist {

		opt := map[string]interface{}{
			"id":        "ctxmod-" + cc.Ptr + "/" + cc.Ctxid + "/" + cc.ID + "/" + cc.Prefix,
			"modid":     cc.ID,
			"template":  "module",
			"icon":      "module.png",
			"name":      cc.ID + " " + cc.Modversion,
			"father":    "ctx-" + container,
			"loadable":  false,
			"closeable": false,
		}
		rows = append(rows, opt)
	}

	return map[string]interface{}{
		"row": rows,
	}

}

func Reloadconfig(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := security.Verify(ctx, security.USER)
	if !ok {
		return ""
	}

	code.ReloadConfig()

	return "OK"
}
