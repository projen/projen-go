//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package circleci

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (c *jsiiProxy_Circleci) validateAddOrbParameters(name *string, orb *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if orb == nil {
		return fmt.Errorf("parameter orb is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Circleci) validateAddWorkflowParameters(workflow *Workflow) error {
	if workflow == nil {
		return fmt.Errorf("parameter workflow is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(workflow, func() string { return "parameter workflow" }); err != nil {
		return err
	}

	return nil
}

func validateNewCircleciParameters(project projen.Project, options *CircleCiProps) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

