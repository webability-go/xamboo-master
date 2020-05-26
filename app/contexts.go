package main

import (
	"fmt"

	"github.com/webability-go/xconfig"

	serverassets "github.com/webability-go/xamboo/assets"
	"master/app/assets"
)

type TContainers []*assets.Container

var Containers = TContainers{}

func (c *TContainers) Load(ctx *serverassets.Context, id string, configfile string) {

	datacontainer := c.UpsertContainer(id, id, configfile)
	// load container
	datacontainer.Config = xconfig.New()
	err := datacontainer.Config.LoadFile(configfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	contexts, _ := datacontainer.Config.GetStringCollection("datasource")
	for _, context := range contexts {
		cfgpath, _ := datacontainer.Config.GetString(context + "+config")
		datacontext := c.UpsertContext(id, context, context, cfgpath)
		datacontext.Config = xconfig.New()
		err := datacontext.Config.LoadFile(cfgpath)
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

func (c *TContainers) UpsertContext(containerid string, contextid string, newid string, path string) *assets.Context {

	ct := c.GetContainer(containerid)
	if ct == nil {
		return nil
	}

	ctx := c.GetContext(containerid, contextid)
	if ctx != nil {
		ctx.ID = newid
		ctx.Path = path
		return ctx
	}

	context := &assets.Context{
		ID:   newid,
		Path: path,
	}
	ct.Contexts = append(ct.Contexts, context)
	return context
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

func (c *TContainers) GetContext(containerid string, contextid string) *assets.Context {

	ct := c.GetContainer(containerid)
	if ct == nil {
		return nil
	}

	for _, x := range ct.Contexts {
		if x.ID == contextid {
			return x
		}
	}
	return nil
}

func (c *TContainers) SaveContainers(ctx *serverassets.Context) {
}

func (c *TContainers) SaveContainer(ctx *serverassets.Context, containerid string) {

	container := c.GetContainer(containerid)

	if container != nil {
		path := container.Path
		// We rebuild the config: the contexts are as objects
		if container.Config == nil {
			container.Config = xconfig.New()
		}

		container.Config.Set("logcore", container.LogFile)
		container.Config.Del("context")
		for _, ct := range container.Contexts {
			container.Config.Add("context", ct.ID)
			container.Config.Set(ct.ID+"+config", ct.Path)
		}
		container.Config.SaveFile(path)
	}
}

func (c *TContainers) SaveContext(ctx *serverassets.Context, containerid string, contextid string) {

	context := c.GetContext(containerid, contextid)
	if context != nil {
		path := context.Path
		if context.Config != nil {
			context.Config.SaveFile(path)
		}
	}
}
