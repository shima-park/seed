package plugin

import (
	"errors"
	"fmt"
	goplugin "plugin"
	"time"
)

type Plugin struct {
	Path     string
	Module   string
	Name     string
	OpenTime time.Time
	Error    error
}

func LoadPlugins(path string) ([]Plugin, error) {
	p, err := goplugin.Open(path)
	if err != nil {
		return nil, err
	}

	sym, err := p.Lookup("Bundle")
	if err != nil {
		return nil, err
	}

	ptr, ok := sym.(*map[string][]interface{})
	if !ok {
		return nil, errors.New("invalid bundle type")
	}

	bundle := *ptr
	var loadedPlugins []Plugin
	for name, plugins := range bundle {
		loader := registry[name]
		if loader == nil {
			continue
		}

		for _, plugin := range plugins {
			pluginName, err := loader(plugin)
			loadedPlugins = append(loadedPlugins, Plugin{
				Path:     path,
				Module:   name,
				Name:     pluginName,
				OpenTime: time.Now(),
				Error:    err,
			})
		}

	}

	return loadedPlugins, nil
}
