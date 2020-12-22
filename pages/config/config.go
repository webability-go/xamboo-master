package main

import (
	// "fmt"
	//	"strings"
	"encoding/xml"

	"github.com/webability-go/xcore/v2"
	"github.com/webability-go/xdommask"

	//	"github.com/webability-go/xamboo"
	"github.com/webability-go/xamboo/assets"

	//	"github.com/webability-go/xmodules/base"

	"master/app/bridge"
)

func Run(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}

	L := ctx.Request.Form.Get("LANGUAGE")
	C := ctx.Request.Form.Get("COUNTRY")

	mask := getMask(L, C).Compile()
	xmlmask, _ := xml.Marshal(mask)

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"FORM": string(xmlmask),
		"#":    language,
	}

	return template.Execute(params)
}

func getMask(lang string, country string) *xdommask.Mask {

	mask := xdommask.NewMask("formconfig")
	mask.Mode = xdommask.VIEW
	mask.AuthModes = xdommask.UPDATE | xdommask.VIEW
	mask.Key = "main"

	mask.AlertMessage = "##mask.errormessage##"
	mask.ServerMessage = "##mask.servermessage##"
	mask.UpdateTitle = "##mask.updatetitle##"
	mask.ViewTitle = "##mask.viewtitle##"

	// bloc 1
	fb1 := xdommask.NewInfoField("title1", "##block1.title##")
	fb1.AuthModes = xdommask.UPDATE | xdommask.VIEW
	mask.AddField(fb1)

	// username
	f11 := xdommask.NewTextField("username")
	f11.Title = "##username.title##"
	f11.HelpDescription = "##username.help.description##"
	f11.NotNullModes = xdommask.UPDATE
	f11.AuthModes = xdommask.UPDATE | xdommask.VIEW
	f11.HelpModes = xdommask.UPDATE
	f11.ViewModes = xdommask.VIEW
	f11.StatusNotNull = "##username.status.notnull##"
	f11.MinLength = 4
	f11.MaxLength = 80
	f11.Size = "400"
	f11.StatusTooShort = "##username.status.tooshort##"
	f11.URLVariable = "username"
	f11.DefaultValue = ""
	mask.AddField(f11)

	// email
	f13 := xdommask.NewMailField("email")
	f13.Title = "##email.title##"
	f13.HelpDescription = "##email.help.description##"
	f13.NotNullModes = xdommask.UPDATE
	f13.AuthModes = xdommask.UPDATE | xdommask.VIEW
	f13.HelpModes = xdommask.UPDATE
	f13.ViewModes = xdommask.VIEW
	f13.StatusNotNull = "##email.status.notnull##"
	f13.MaxLength = 200 // chars
	f13.Size = "400"
	f13.URLVariable = "email"
	f13.DefaultValue = ""
	mask.AddField(f13)

	// password
	f12 := xdommask.NewPWField("password")
	f12.Title = "##password.title##"
	f12.HelpDescription = "##password.help.description##"
	f12.NotNullModes = xdommask.UPDATE
	f12.AuthModes = xdommask.UPDATE
	f12.HelpModes = xdommask.UPDATE
	f12.ViewModes = xdommask.VIEW
	f12.StatusNotNull = "##password.status.notnull##"
	f12.MaxLength = 80
	f12.MinLength = 4
	f12.Size = "400"
	f12.StatusTooShort = "##password.status.tooshort##"
	f12.URLVariable = "password"
	f12.DefaultValue = ""
	mask.AddField(f12)

	fb2 := xdommask.NewInfoField("title2", "##block2.title##")
	fb2.AuthModes = xdommask.UPDATE | xdommask.VIEW
	mask.AddField(fb2)

	// language
	f21 := xdommask.NewTextField("language")
	f21.Title = "##language.title##"
	f21.HelpDescription = "##language.help.description##"
	f21.NotNullModes = xdommask.UPDATE
	f21.AuthModes = xdommask.UPDATE | xdommask.VIEW
	f21.HelpModes = xdommask.UPDATE
	f21.ViewModes = xdommask.VIEW
	f21.StatusNotNull = "##language.status.notnull##"
	f21.MinLength = 2
	f21.MaxLength = 2
	f21.Size = "400"
	f21.StatusTooShort = "##language.status.tooshort##"
	f21.URLVariable = "language"
	f21.DefaultValue = ""
	mask.AddField(f21)

	// country
	f22 := xdommask.NewTextField("country")
	f22.Title = "##country.title##"
	f22.HelpDescription = "##country.help.description##"
	f22.NotNullModes = xdommask.UPDATE
	f22.AuthModes = xdommask.UPDATE | xdommask.VIEW
	f22.HelpModes = xdommask.UPDATE
	f22.ViewModes = xdommask.VIEW
	f22.StatusNotNull = "##country.status.notnull##"
	f22.MaxLength = 2
	f22.MinLength = 2
	f22.Size = "400"
	f22.StatusTooShort = "##country.status.tooshort##"
	f22.URLVariable = "country"
	f22.DefaultValue = ""
	mask.AddField(f22)

	// serial
	f23 := xdommask.NewTextField("serial")
	f23.Title = "##serial.title##"
	f23.HelpDescription = "##serial.help.description##"
	f23.NotNullModes = xdommask.UPDATE
	f23.AuthModes = xdommask.UPDATE | xdommask.VIEW
	f23.HelpModes = xdommask.UPDATE
	f23.ViewModes = xdommask.VIEW
	f23.StatusNotNull = "##serial.status.notnull##"
	f23.MaxLength = 16 // chars
	f23.MinLength = 16
	f23.Size = "400"
	f23.StatusTooShort = "##serial.status.tooshort##"
	f23.URLVariable = "serial"
	f23.DefaultValue = ""
	mask.AddField(f23)

	f7 := xdommask.NewButtonField("", "update")
	f7.AuthModes = xdommask.VIEW
	f7.TitleView = "##form.update##"
	mask.AddField(f7)

	f8 := xdommask.NewButtonField("", "submit")
	f8.AuthModes = xdommask.UPDATE
	f8.TitleUpdate = "##form.continue##"
	mask.AddField(f8)

	f9 := xdommask.NewButtonField("", "reset")
	f9.AuthModes = xdommask.UPDATE
	f9.TitleUpdate = "##form.reset##"
	mask.AddField(f9)

	return mask
}

func Formconfig(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}

	order := ctx.Request.Form.Get("Order")
	if order == "getrecord" {

		data := map[string]string{}
		data["serial"], _ = ctx.Sysparams.GetString("serial")
		data["username"], _ = ctx.Sysparams.GetString("username")
		data["email"], _ = ctx.Sysparams.GetString("email")
		data["language"], _ = ctx.Sysparams.GetString("language")
		data["country"], _ = ctx.Sysparams.GetString("country")

		edata := map[string]interface{}{}
		edata["main"] = data

		return edata
	}

	serial := ctx.Request.Form.Get("serial")
	username := ctx.Request.Form.Get("username")
	password := ctx.Request.Form.Get("password")
	email := ctx.Request.Form.Get("email")
	L := ctx.Request.Form.Get("language")
	C := ctx.Request.Form.Get("country")

	// check params ok
	success := true
	messages := map[string]string{}
	messagetext := ""
	if username == "" {
		success = false
		messages["username"] = language.Get("error.username.missing")
		messagetext += language.Get("error.username.missing")
	}
	if username != "" && username == password {
		success = false
		messages["password"] = language.Get("error.password.same")
		messagetext += language.Get("error.password.same")
	}

	// TODO(phil) verify language and country are authorized, password is strong enough

	if success {
		bridge.GenerateConfig(ctx, L, C, serial, username, password, email)
		messages["text"] = language.Get("success")
	} else {
		messages["text"] = messagetext
	}

	return map[string]interface{}{
		"success": success, "messages": messages, "popup": false,
	}
}
