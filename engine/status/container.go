package status

import (
	"gitlab.ricebook.net/platform/agent/types"
	"gitlab.ricebook.net/platform/agent/utils"
)

func GenerateContainerMeta(ID string, attrs map[string]string) (*types.Container, error) {
	//干掉 image 的信息
	delete(attrs, "image")

	name, entrypoint, ident, err := utils.GetAppInfo(attrs["name"])
	if err != nil {
		return nil, err
	}
	delete(attrs, "name")

	container := &types.Container{}
	container.ID = ID
	container.Name = name
	container.EntryPoint = entrypoint
	container.Ident = ident
	if v, ok := attrs["version"]; ok {
		container.Version = v
		delete(attrs, "version")
	} else {
		container.Version = "UNKNOWN"
	}

	// stat: vendor  CentOS
	// stat: build-date  20160906
	// stat: license  GPLv2
	// 都不知道为什么会出现这三项

	delete(attrs, "vendor")
	delete(attrs, "build-date")
	delete(attrs, "license")
	container.Extend = attrs

	return container, nil
}
