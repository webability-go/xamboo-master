package main

import (
	"golang.org/x/text/language"

	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"

	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/tools"

	"master/app/code"
)

const (
	VERSION = "0.0.1"
)

var (
	Application = LocalApp{}
	Containers  = code.Containers
)

func init() {
	tools.Language = language.Spanish
}

type LocalApp struct{}

func (la *LocalApp) StartHost(h config.Host) {
}

func (la *LocalApp) StartContext(ctx *context.Context) {
	code.VerifyLogin(ctx)
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

func VerifyLogin(ctx *context.Context) {
	code.VerifyLogin(ctx)
}

func GetMD5Hash(text string) string {
	return code.GetMD5Hash(text)
}

func CreateKey(length int, chartype int) string {
	return code.CreateKey(length, chartype)
}

func GetMainConfig() map[string]interface{} {
	return code.GetMainConfig()
}
