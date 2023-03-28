//go:build !no_runtime_type_checking

package vscode

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IDevContainerEnvironment) validateAddFeaturesParameters(features *[]*DevContainerFeature) error {
	for idx_5b8a8b, v := range *features {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter features[%#v]", idx_5b8a8b) }); err != nil {
			return err
		}
	}

	return nil
}

