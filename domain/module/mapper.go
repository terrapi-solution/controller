package module

import "github.com/terrapi-solution/controller/data/module"

func (r ModuleRequest) toDBModel() module.Module {
	return module.Module{
		Name: r.Name,
	}
}
