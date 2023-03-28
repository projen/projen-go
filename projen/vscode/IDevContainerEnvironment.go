package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// Experimental.
type IDevContainerEnvironment interface {
	projen.IDevEnvironment
	// Adds a list of VSCode features that should be automatically installed in the container.
	// Experimental.
	AddFeatures(features ...*DevContainerFeature)
}

// The jsii proxy for IDevContainerEnvironment
type jsiiProxy_IDevContainerEnvironment struct {
	internal.Type__projenIDevEnvironment
}

func (i *jsiiProxy_IDevContainerEnvironment) AddFeatures(features ...*DevContainerFeature) {
	if err := i.validateAddFeaturesParameters(&features); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range features {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addFeatures",
		args,
	)
}

