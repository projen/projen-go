//go:build !no_runtime_type_checking

package java

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func validateNewProjenrcParameters(project projen.Project, pom Pom, options *ProjenrcOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if pom == nil {
		return fmt.Errorf("parameter pom is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

