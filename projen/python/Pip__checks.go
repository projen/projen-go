//go:build !no_runtime_type_checking

package python

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (p *jsiiProxy_Pip) validateAddDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Pip) validateAddDevDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func validatePip_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePip_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewPipParameters(project projen.Project, _options *PipOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

