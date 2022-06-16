package main

import (
	"fmt"

	"github.com/webability-go/xamboo/cms/context"
	xcore "github.com/webability-go/xcore/v2"

	"masterapp/assets"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := assets.Verify(ctx, assets.USER)
	if !ok {
		return ""
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"#":    language,
		"data": getData(ctx, language),
	}

	return template.Execute(params)
}

func getData(ctx *context.Context, language *xcore.XLanguage) string {

	resourcesdir, _ := ctx.Sysparams.GetString("resourcesdir")

	code := assets.GetEntries()
	cfg := code.GetMainConfig()
	/*
		fmt.Println("DATA:", data)

		log := cfg["log"]
		if log != nil {
			ilog := log.(map[string]interface{})
			data["errors"] = ilog["errors"].(string)
			data["sys"] = ilog["sys"].(string)
		}
		includes := cfg["include"]
		files := ""
		if includes != nil {
			iinclude := includes.([]interface{})
			for _, include := range iinclude {
				files += include.(string) + "\n"
			}
		}
		data["files"] = files
	*/
	//		fmt.Printf("%#v \n", cfg)

	/*

		fb3 := xdommask.NewInfoField("title3", "##block3.title##")
		fb3.AuthModes = xdommask.UPDATE | xdommask.VIEW
		mask.AddField(fb3)

		// sys
		f31 := xdommask.NewTextField("sys")
		f31.Title = "##sys.title##"
		f31.HelpDescription = "##sys.help.description##"
		f31.NotNullModes = xdommask.UPDATE
		f31.AuthModes = xdommask.UPDATE | xdommask.VIEW
		f31.HelpModes = xdommask.UPDATE
		f31.ViewModes = xdommask.VIEW
		f31.StatusNotNull = "##sys.status.notnull##"
		f31.MaxLength = 400
		f31.Size = "400"
		f31.URLVariable = "sys"
		f31.DefaultValue = ""
		mask.AddField(f31)

		// errors
		f32 := xdommask.NewTextField("errors")
		f32.Title = "##errors.title##"
		f32.HelpDescription = "##errors.help.description##"
		f32.NotNullModes = xdommask.UPDATE
		f32.AuthModes = xdommask.UPDATE | xdommask.VIEW
		f32.HelpModes = xdommask.UPDATE
		f32.ViewModes = xdommask.VIEW
		f32.StatusNotNull = "##errors.status.notnull##"
		f32.MaxLength = 80
		f32.MinLength = 4
		f32.Size = "400"
		f32.StatusTooShort = "##errors.status.tooshort##"
		f32.URLVariable = "errors"
		f32.DefaultValue = ""
		mask.AddField(f32)

		fb4 := xdommask.NewInfoField("title4", "##block4.title##")
		fb4.AuthModes = xdommask.UPDATE | xdommask.VIEW
		mask.AddField(fb4)

		// files
		f41 := xdommask.NewTextAreaField("files")
		f41.Title = "##file.title##"
		f41.HelpDescription = "##file.help.description##"
		f41.NotNullModes = xdommask.UPDATE
		f41.AuthModes = xdommask.UPDATE | xdommask.VIEW
		f41.HelpModes = xdommask.UPDATE
		f41.ViewModes = xdommask.VIEW
		f41.StatusNotNull = "##file.status.notnull##"
		f41.MinLength = 4
		f41.MaxLength = 80
		f41.Width = 500
		f41.Height = 300
		f41.StatusTooShort = "##file.status.tooshort##"
		f41.URLVariable = "file"
		f41.DefaultValue = ""
		mask.AddField(f41)

		smask := xdommask.NewMask("formfilesconfig")
		smask.Mode = xdommask.VIEW
		smask.AuthModes = xdommask.UPDATE | xdommask.VIEW
		smask.Key = "main"

		smask.AlertMessage = "##mask.errormessage##"
		smask.ServerMessage = "##mask.servermessage##"
		smask.UpdateTitle = "##mask.updatetitle##"
		smask.ViewTitle = "##mask.viewtitle##"

		// file
		sf11 := xdommask.NewTextField("file")
		sf11.Title = "##file.title##"
		sf11.HelpDescription = "##file.help.description##"
		sf11.NotNullModes = xdommask.UPDATE
		sf11.AuthModes = xdommask.UPDATE | xdommask.VIEW
		sf11.HelpModes = xdommask.UPDATE
		sf11.ViewModes = xdommask.VIEW
		sf11.StatusNotNull = "##file.status.notnull##"
		sf11.MinLength = 4
		sf11.MaxLength = 80
		sf11.Size = "400"
		sf11.StatusTooShort = "##file.status.tooshort##"
		sf11.URLVariable = "file"
		sf11.DefaultValue = ""
		mask.AddField(sf11)
	*/
	/*
		<entry id="block3.title"><![CDATA[Logs:]]></entry>

		<entry id="sys.title"><![CDATA[Log del sistema:]]></entry>
		<entry id="sys.help.description"><![CDATA[El log del sistema puede ser "discard", "stdout:", "stderr:", o un archivo en disco duro. Asegurese de tener los accesos de escrutura sobre el archivo y el directorio.]]></entry>
		<entry id="sys.status.notnull"><![CDATA[No puede ser vacio.]]></entry>

		<entry id="errors.title"><![CDATA[Log de errores del sistema:]]></entry>
		<entry id="errors.help.description"><![CDATA[El log de errores del sistema puede ser "discard", "stdout:", "stderr:", o un archivo en disco duro. Asegurese de tener los accesos de escrutura sobre el archivo y el directorio.]]></entry>
		<entry id="errors.status.notnull"><![CDATA[No puede ser vacio.]]></entry>


		<entry id="block4.title"><![CDATA[Archivos contenedores:]]></entry>
	*/

	return fmt.Sprintf("%s <br /> %#v", resourcesdir, cfg)
}
