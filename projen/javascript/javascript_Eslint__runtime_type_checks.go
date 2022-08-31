//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package javascript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (e *jsiiProxy_Eslint) validateAddIgnorePatternParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Eslint) validateAddOverrideParameters(override *EslintOverride) error {
	if override == nil {
		return fmt.Errorf("parameter override is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(override, func() string { return "parameter override" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_Eslint) validateAddRulesParameters(rules *map[string]interface{}) error {
	if rules == nil {
		return fmt.Errorf("parameter rules is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Eslint) validateAllowDevDepsParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateEslint_OfParameters(project projen.Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func validateNewEslintParameters(project NodeProject, options *EslintOptions) error {
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

