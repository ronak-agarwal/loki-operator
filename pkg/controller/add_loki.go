package controller

import (
	"github.com/ronak-agarwal/loki-operator/pkg/controller/loki"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, loki.Add)
}
