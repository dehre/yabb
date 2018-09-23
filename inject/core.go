package inject

import (
	"net/http"

	"github.com/L-oris/yabb/inject/types"
	"github.com/L-oris/yabb/models/tpl"
	"github.com/sarulabs/di"
)

func core() []di.Def {
	templates := di.Def{
		Name: types.Templates.String(),
		Build: func(ctn di.Container) (interface{}, error) {
			return &tpl.TPL{}, nil
		},
	}

	fileserver := di.Def{
		Name: types.FileServer.String(),
		Build: func(ctn di.Container) (interface{}, error) {
			return http.ServeFile, nil
		},
	}

	return []di.Def{
		templates, fileserver,
	}
}