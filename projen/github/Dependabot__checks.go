//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_Dependabot) validateAddAllowParameters(dependencyName *string) error {
	if dependencyName == nil {
		return fmt.Errorf("parameter dependencyName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_Dependabot) validateAddIgnoreParameters(dependencyName *string) error {
	if dependencyName == nil {
		return fmt.Errorf("parameter dependencyName is required, but nil was provided")
	}

	return nil
}

func validateDependabot_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateDependabot_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDependabotParameters(github GitHub, options *DependabotOptions) error {
	if github == nil {
		return fmt.Errorf("parameter github is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

