//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/projen/projen-go/projen"
)

func validateCdkTasks_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCdkTasks_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCdkTasksParameters(project projen.Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

