//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func (d *jsiiProxy_Dependencies) validateAddDependencyParameters(spec *string, type_ DependencyType) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_Dependencies) validateGetDependencyParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_Dependencies) validateRemoveDependencyParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_Dependencies) validateTryGetDependencyParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateDependencies_ParseDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func validateNewDependenciesParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

