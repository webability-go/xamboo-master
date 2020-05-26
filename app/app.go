package main

import (
	"github.com/webability-go/xamboo/assets"

	"github.com/webability-go/xmodules/base"

	_ "github.com/webability-go/xmodules/client"
	_ "github.com/webability-go/xmodules/clientlink"
	_ "github.com/webability-go/xmodules/clientp18n"
	_ "github.com/webability-go/xmodules/clientsecurity"
	_ "github.com/webability-go/xmodules/country"
	_ "github.com/webability-go/xmodules/ingredient"
	_ "github.com/webability-go/xmodules/material"
	_ "github.com/webability-go/xmodules/metric"
	_ "github.com/webability-go/xmodules/stat"
	_ "github.com/webability-go/xmodules/suggestions"
	_ "github.com/webability-go/xmodules/translation"
	_ "github.com/webability-go/xmodules/usda"
	_ "github.com/webability-go/xmodules/user"
	_ "github.com/webability-go/xmodules/userlink"
	//	_ "github.com/webability-go/xmodules/usermenu"
	_ "github.com/webability-go/xmodules/wiki"
)

const (
	VERSION              = "1.0.0"
	DATASOURCECONFIGFILE = "./master/config/datasources.conf"
)

var (
	Container   *base.Container
	Application = LocalApp{}
)

func init() {
	Container = base.Create(DATASOURCECONFIGFILE)
}

type LocalApp struct{}

func (la *LocalApp) StartHost(h assets.Host) {
}

func (la *LocalApp) StartContext(ctx *assets.Context) {
	VerifyLogin(ctx)
}

func (la *LocalApp) GetDatasourcesConfigFile() string {
	return DATASOURCECONFIGFILE
}

func (la *LocalApp) GetDatasourceSet() assets.DatasourceSet {
	return Container
}

func (la *LocalApp) GetCompiledModules() assets.ModuleSet {
	return base.ModulesList
}
