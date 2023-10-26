//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen/github/workflows"
)

func (u *jsiiProxy_UpgradeDependencies) validateAddPostBuildStepsParameters(steps *[]*workflows.JobStep) error {
	for idx_b7595e, v := range *steps {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter steps[%#v]", idx_b7595e) }); err != nil {
			return err
		}
	}

	return nil
}

func validateUpgradeDependencies_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateUpgradeDependencies_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_UpgradeDependencies) validateSetContainerOptionsParameters(val *workflows.ContainerOptions) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewUpgradeDependenciesParameters(project NodeProject, options *UpgradeDependenciesOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

