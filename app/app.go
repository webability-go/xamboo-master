package main

import (
	"golang.org/x/text/language"

	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"

	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/tools"

	"masterapp/code"
)

const (
	VERSION = "0.0.1"
)

var (
	Application = LocalApp{}
	Containers  = code.Containers
)

func init() {
	// Have to use the system language or english by default
	// YOu may change default language of the system (know languages: English, Spanish, French)
	tools.Language = language.Spanish
}

type LocalApp struct{}

func (la *LocalApp) StartHost(h config.Host) {
}

func (la *LocalApp) StartContext(ctx *context.Context) {
	code.VerifySession(ctx)
}

func (la *LocalApp) GetDatasourcesConfigFile() string {
	return ""
}

func (la *LocalApp) GetDatasourceSet() applications.DatasourceSet {
	return nil
}

func (la *LocalApp) GetCompiledModules() applications.ModuleSet {
	return base.ModulesList
}

func VerifySession(ctx *context.Context) {
	code.VerifySession(ctx)
}

func ReloadConfig() error {
	return code.ReloadConfig()
}

func GetMainConfig() map[string]interface{} {
	return code.GetMainConfig()
}

func GenerateConfig(ctx *context.Context, L string, C string, serial string, username string, password string, email string) {
	code.GenerateConfig(ctx, L, C, serial, username, password, email)
}
