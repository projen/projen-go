//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package web

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen/javascript"
)

func validateNewTailwindConfigParameters(project javascript.NodeProject, options *TailwindConfigOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

