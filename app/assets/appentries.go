package assets

import (
	"github.com/webability-go/xamboo/cms/context"
)

type AppEntries struct {
	VerifySession  func(ctx *context.Context)
	ReloadConfig   func() error
	GetMainConfig  func() map[string]interface{}
	GenerateConfig func(ctx *context.Context, L string, C string, serial string, username string, password string, email string)

	Containers ContainersList
}

var MasterAppEntries *AppEntries

func GetEntries() *AppEntries {
	return MasterAppEntries
}
