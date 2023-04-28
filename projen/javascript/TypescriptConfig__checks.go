//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (t *jsiiProxy_TypescriptConfig) validateAddExcludeParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TypescriptConfig) validateAddExtendsParameters(value TypescriptConfig) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TypescriptConfig) validateAddIncludeParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TypescriptConfig) validateResolveExtendsPathParameters(configPath *string) error {
	if configPath == nil {
		return fmt.Errorf("parameter configPath is required, but nil was provided")
	}

	return nil
}

func validateNewTypescriptConfigParameters(project projen.Project, options *TypescriptConfigOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

