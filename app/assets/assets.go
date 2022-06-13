package assets

import (
	"fmt"

	"github.com/webability-go/xconfig"

	"github.com/webability-go/xamboo/cms/context"
)

// The datasource containers file is {resourcesdir}/containers.conf
type Datasource struct {
	ID     string
	Path   string
	Config *xconfig.XConfig
}

type Container struct {
	Name        string
	Path        string
	LogFile     string
	Config      *xconfig.XConfig
	Datasources []*Datasource
}

type ContainersList interface {
	fmt.Stringer   // please implement String()
	fmt.GoStringer // Please implement GoString()

	Load(ctx *context.Context, id string, datasourcefile string)
	UpsertContainer(containerid string, newid string, path string) *Container
	UpsertDatasource(containerid string, datasourceid string, newid string, path string) *Datasource

	GetContainersList() []*Container
	GetContainer(containerid string) *Container
	GetDatasource(containerid string, datasourceid string) *Datasource

	// save only the list of containers ~/resources/datasourcescontainers.conf
	SaveContainers(ctx *context.Context)
	// save only the container contexts config ~/config/contexts.conf
	SaveContainer(ctx *context.Context, containerid string)
	// save only the context config ~/config/ID.conf
	SaveDatasource(ctx *context.Context, containerid string, datasourceid string)
}

func (c *Container) String() string {
	s := "CONTAINER: " + fmt.Sprintf("%v\n", *c)
	return s
}

func (c *Container) GoString() string {
	s := "#CONTAINER: " + fmt.Sprintf("%#v\n", *c)
	return s
}

func (c *Datasource) String() string {
	s := "CONTEXT: " + fmt.Sprintf("%v\n", *c)
	return s
}

func (c *Datasource) GoString() string {
	s := "#CONTEXT: " + fmt.Sprintf("%#v\n", *c)
	return s
}
