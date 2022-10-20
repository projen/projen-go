//go:build !no_runtime_type_checking

package python

import (
	"fmt"
)

func (i *jsiiProxy_IPythonDeps) validateAddDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IPythonDeps) validateAddDevDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

