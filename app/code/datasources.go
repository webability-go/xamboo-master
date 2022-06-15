package code

import (
	"fmt"

	"github.com/webability-go/xconfig"

	"masterapp/assets"

	"github.com/webability-go/xamboo/cms/context"
)

type TContainers []*assets.Container

var Containers = TContainers{}

func (c *TContainers) Load(ctx *context.Context, id string, configfile string) {

	datacontainer := c.UpsertContainer(id, id, configfile)
	// load container
	datacontainer.Config = xconfig.New()
	err := datacontainer.Config.LoadFile(configfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	datasources, _ := datacontainer.Config.GetStringCollection("datasource")
	for _, datasource := range datasources {
		cfgpath, _ := datacontainer.Config.GetString(datasource + "+config")
		datadatasource := c.UpsertDatasource(id, datasource, datasource, cfgpath)
		datadatasource.Config = xconfig.New()
		err := datadatasource.Config.LoadFile(cfgpath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func (c *TContainers) String() string {
	s := "CONTAINERS: "
	for _, ct := range *c {
		s += fmt.Sprintf("%#v\n", *ct)
	}
	return s
}

func (c *TContainers) GoString() string {
	s := "#CONTAINERS: "
	for _, ct := range *c {
		s += fmt.Sprintf("%#v\n", *ct)
	}
	return s
}

func (c *TContainers) UpsertContainer(containerid string, newid string, path string) *assets.Container {

	for _, cont := range *c {
		if cont.Name == newid {
			cont.Name = newid
			cont.Path = path
			return cont
		}
	}

	container := &assets.Container{
		Name: newid,
		Path: path,
	}
	*c = append(*c, container)
	return container
}

func (c *TContainers) UpsertDatasource(containerid string, datasourceid string, newid string, path string) *assets.Datasource {

	ct := c.GetContainer(containerid)
	if ct == nil {
		return nil
	}

	ctx := c.GetDatasource(containerid, datasourceid)
	if ctx != nil {
		ctx.ID = newid
		ctx.Path = path
		return ctx
	}

	datasource := &assets.Datasource{
		ID:   newid,
		Path: path,
	}
	ct.Datasources = append(ct.Datasources, datasource)
	return datasource
}

func (c *TContainers) GetContainersList() []*assets.Container {

	data := []*assets.Container{}
	for _, x := range *c {
		data = append(data, x)
	}
	return data
}

func (c *TContainers) GetContainer(containerid string) *assets.Container {

	for _, x := range *c {
		if x.Name == containerid {
			return x
		}
	}
	return nil
}

func (c *TContainers) GetDatasource(containerid string, datasourceid string) *assets.Datasource {

	ct := c.GetContainer(containerid)
	if ct == nil {
		return nil
	}

	for _, x := range ct.Datasources {
		if x.ID == datasourceid {
			return x
		}
	}
	return nil
}

func (c *TContainers) SaveContainers(ctx *context.Context) {
}

func (c *TContainers) SaveContainer(ctx *context.Context, containerid string) {

	container := c.GetContainer(containerid)

	if container != nil {
		path := container.Path
		// We rebuild the config: the contexts are as objects
		if container.Config == nil {
			container.Config = xconfig.New()
		}

		container.Config.Set("logcore", container.LogFile)
		container.Config.Del("datasource")
		for _, ct := range container.Datasources {
			container.Config.Add("datasource", ct.ID)
			container.Config.Set(ct.ID+"+config", ct.Path)
		}
		container.Config.SaveFile(path)
	}
}

func (c *TContainers) SaveDatasource(ctx *context.Context, containerid string, datasourceid string) {

	datasource := c.GetDatasource(containerid, datasourceid)
	if datasource != nil {
		path := datasource.Path
		if datasource.Config != nil {
			datasource.Config.SaveFile(path)
		}
	}
}
