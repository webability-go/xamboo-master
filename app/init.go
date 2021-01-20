package main

import (
	"golang.org/x/text/language"

	"github.com/webability-go/xamboo/assets"

	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/tools"
)

const (
	VERSION = "0.0.1"
)

var (
	Application = LocalApp{}
)

func init() {
	tools.Language = language.Spanish
}

type LocalApp struct{}

func (la *LocalApp) StartHost(h assets.Host) {
}

func (la *LocalApp) StartContext(ctx *assets.Context) {
	VerifyLogin(ctx)
}

func (la *LocalApp) GetDatasourcesConfigFile() string {
	return ""
}

func (la *LocalApp) GetDatasourceSet() assets.DatasourceSet {
	return nil
}

func (la *LocalApp) GetCompiledModules() assets.ModuleSet {
	return base.ModulesList
}
