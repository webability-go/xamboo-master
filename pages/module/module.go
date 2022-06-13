package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/webability-go/xcore/v2"

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

	moduledata := generateModuleData(ctx, s.(*cms.CMS), module)
	data := generateData(ctx, s.(*cms.CMS), module)

	params := &xcore.XDataset{
		"module":     module,
		"moduledata": moduledata,
		"DATA":       data,
		"#":          language,
	}
	return template.Execute(params)
}

type cc struct {
	MODID            string
	CNTID            string
	DSID             string
	Modversion       string
	Needs            string
	Installedversion string
	Prefix           string
	Path             map[string]string
}

func generateModuleData(ctx *context.Context, s *cms.CMS, module string) string {

	// config := s.GetFullConfig()
	md := base.ModulesList.Get(module)

	lang, _ := ctx.Sysparams.GetString("language")
	ilang := language.English
	sl := ""

	l := md.GetLanguages()
	for le := range l {
		if fmt.Sprint(le) == lang {
			ilang = le
			sl += "[<b style=\"color: #00cc00;\">" + fmt.Sprint(le) + "</b>]"
		} else {
			sl += "[<b>" + fmt.Sprint(le) + "</b>]"
		}
	}

	data := "[" + md.GetID() + "@v" + md.GetVersion() + "] "
	if l[ilang] != "" {
		data += l[ilang]
	} else {
		data += l[language.English]
	}
	data += "<br />Supported languages: " + sl
	data += "<br />Modules needed: " + fmt.Sprint(md.GetNeeds())

	return data
}

func generateData(ctx *context.Context, s *cms.CMS, module string) string {

	md := base.ModulesList.Get(module)
	if md == nil {
		return "-- Module IS NOT COMPILED into xamboo !!"
	}

	// config := s.GetFullConfig()
	data := "<table class=\"module-table\">"
	data += "<tr><td>Fuente de datos</td><td>Versión compilada</td><td>Necesidades</td><td>Versión instalada</td><td>Prefijo</td><td>Data</td><td>Paths</td>"

	// 1. Get all the contexts with and without the module authorized
	contextcontainerslist := map[string]cc{}

	// Module installed x datasource
	for cntid, ct := range *base.Containers {
		dss := ct.GetDatasources()
		for dsid, ds := range dss {

			// filter only dss that use this module
			if !ds.IsModuleAuthorized(module) {
				//				fmt.Println("Module IS NOT authorized on ", cntid, dsid)
				continue
			}
			//			fmt.Println("Module authorized on ", module, cntid, dsid, mv)

			// yes:
			//			fmt.Println("++ Module compiled and authorized on xamboo")
			modprefix := ""
			// search for the paths
			cds := ds.(*base.Datasource)
			pathadmin, _ := cds.Config.GetString("pathinstalladmin")
			pathadminstatic, _ := cds.Config.GetString("pathinstalladminstatic")
			pathsite, _ := cds.Config.GetString("pathinstallsite")
			pathsitestatic, _ := cds.Config.GetString("pathinstallsitestatic")

			// needed Modules
			needs := md.GetNeeds()
			nmodules := ""
			for _, idmod := range needs {
				nmd := base.ModulesList.Get(idmod)
				if nmd == nil {
					nmodules += "<span style=\"background-color: red;\">[" + idmod + "]</span>"
					continue
				}
				ninst := nmd.GetInstalledVersion(ds)
				if ninst == "" {
					nmodules += "<span style=\"background-color: red;\">[" + idmod + "]</span>"
					continue
				}
				nmodules += "<span style=\"background-color: #afa;\">[" + idmod + "]</span>"
			}

			contextcontainerslist[cntid+"::"+dsid] = cc{
				MODID:            module,
				CNTID:            cntid,
				DSID:             dsid,
				Modversion:       md.GetVersion(),
				Needs:            nmodules,
				Installedversion: md.GetInstalledVersion(ds),
				Prefix:           modprefix,
				Path: map[string]string{
					"pathadmin":       pathadmin,
					"pathadminstatic": pathadminstatic,
					"pathsite":        pathsite,
					"pathsitestatic":  pathsitestatic,
				},
			}
		}
	}

	for _, cc := range contextcontainerslist {
		data += "<tr>"
		data += "<td>" + cc.CNTID + "::" + cc.DSID + "</td>"
		data += "<td>" + cc.Modversion + "</td>"
		data += "<td>" + cc.Needs
		data += "<td>" + cc.Installedversion + "</td>"
		data += "<td>" + cc.Prefix + "</td>"

		data += "<td>"
		if cc.Installedversion == "" {
			data += "[<span style=\"background-color: #dfd; cursor: pointer;\" onclick=\"module_install('" + cc.MODID + "', '" + cc.CNTID + "', '" + cc.DSID + "', '" + cc.Prefix + "', null);\">INSTALAR</span>]"
		} else if cc.Installedversion != cc.Modversion {
			data += "[<span style=\"background-color: #ddf; cursor: pointer;\" onclick=\"module_upgrade('" + cc.MODID + "', '" + cc.CNTID + "', '" + cc.DSID + "', '" + cc.Prefix + "', null);\">ACTUALIZAR</span>]"
		}
		if cc.Installedversion != "" {
			//		data += "[<span style=\"background-color: #ddf; cursor: pointer;\" onclick=\"reinstall('" + cc.MODID + "', '" + cc.DSID + "', '" + module + "', '" + cc.Prefix + "');\">REINSTALAR</span>]"
			//		data += "[<span style=\"background-color: #fdd; cursor: pointer;\" onclick=\"uninstall('" + cc.MODID + "', '" + cc.DSID + "', '" + module + "', '" + cc.Prefix + "');\">DESINSTALAR</span>]"
		}
		data += "</td>"

		data += "<td>"
		for ip, p := range cc.Path {
			// Verify installation of files into directory to know if we put "ok" or "missing" flag
			if p == "" {
				continue
			}
			data += ip + ": " + p + " [<span style=\"background-color: #dfd; cursor: pointer;\" onclick=\"module_install('" + cc.MODID + "', '" + cc.CNTID + "', '" + cc.DSID + "', '" + cc.Prefix + "', '" + ip + "');\">COPY FILES</span>]<br />"
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
	//	srv := s.(*cms.CMS)

	module := ctx.Request.Form.Get("module")
	container := ctx.Request.Form.Get("container")
	datasource := ctx.Request.Form.Get("datasource")
	prefix := ctx.Request.Form.Get("prefix")
	//	idpath := ctx.Request.Form.Get("idpath")

	// find module
	md := base.ModulesList.Get(module)
	if md == nil {
		return map[string]interface{}{
			"success": false, "messages": "Module " + module + " does not exists in the system", "popup": false, "result": "", "error": "",
		}
	}

	// find datasource to use
	dsc := base.Containers.GetContainer(container)
	if dsc == nil {
		return map[string]interface{}{
			"success": false, "messages": "Container " + container + " does not exists in the system", "popup": false, "result": "", "error": "",
		}
	}

	ds := dsc.GetDatasource(datasource)
	if ds == nil {
		return map[string]interface{}{
			"success": false, "messages": "Datasource " + datasource + " does not exists in the container " + container, "popup": false, "result": "", "error": "",
		}
	}

	result, err := md.Synchronize(ds, prefix)
	success := true
	message := ""
	return map[string]interface{}{
		"success": success, "messages": message, "popup": false, "result": result, "error": err,
	}
}

func Uninstall(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {
	return map[string]interface{}{
		"success": true, "messages": "Not implemented yet", "popup": false, "result": "", "error": "",
	}
}
