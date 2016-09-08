package local

/**
 * Local handler provides operations based entirely on the
 * Local environment, primarily based on config files in
 * a project, based on the current path'
 */

import (
	"github.com/james-nesbitt/wundertools-go/api"
	"github.com/james-nesbitt/wundertools-go/api/handler"
)

func NewLocalHandler() api.API {
	localHandler := LocalHandler{}
	return &nullHandler
}

type LocalHandler struct {
}