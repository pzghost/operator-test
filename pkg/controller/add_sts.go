package controller

import (
	"github.com/pzghost/operator-sdk-statefulset/pkg/controller/sts"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sts.Add)
}
