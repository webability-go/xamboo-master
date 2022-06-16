package main

import (
	"golang.org/x/text/language"

	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"

	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/tools"

	"masterapp/assets"
	"masterapp/code"
)

const (
	VERSION = "0.0.2"
)

var (
	Application = LocalApp{}
	Containers  = code.Containers
)

func init() {
	// Have to use the system language or english by default
	// You may change default language of the system (known languages: English, Spanish, French)
	tools.Language = language.Spanish

	assets.MasterAppEntries = &assets.AppEntries{
		VerifySession:  code.VerifySession,
		ReloadConfig:   code.ReloadConfig,
		GetMainConfig:  code.GetMainConfig,
		GenerateConfig: code.GenerateConfig,
		Containers:     &Containers,
	}
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
