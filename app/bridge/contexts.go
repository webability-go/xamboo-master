package bridge

import (
	"errors"
	"fmt"
	"plugin"

	"master/app/assets"
)

var Containers assets.ContainersList

func LinkContexts(lib *plugin.Plugin) error {

	obj, err := lib.Lookup("Containers")
	if err != nil {
		fmt.Println(err)
		return errors.New("ERROR: THE APPLICATION LIBRARY DOES NOT CONTAIN Containers OBJECT")
	}
	Containers = obj.(assets.ContainersList)

	return nil
}
