//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_Tasks) validateAddEnvironmentParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateAddTaskParameters(name *string, options *TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Tasks) validatePostProjectCreationParameters(initProject *InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateProjectCreationParameters(initProject *InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateRemoveTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateResolveTasksManifestParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateRunTaskParameters(name *string, args *[]interface{}) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	for idx_090772, v := range *args {
		switch v.(type) {
		case *string:
			// ok
		case string:
			// ok
		case *float64:
			// ok
		case float64:
			// ok
		case *int:
			// ok
		case int:
			// ok
		case *uint:
			// ok
		case uint:
			// ok
		case *int8:
			// ok
		case int8:
			// ok
		case *int16:
			// ok
		case int16:
			// ok
		case *int32:
			// ok
		case int32:
			// ok
		case *int64:
			// ok
		case int64:
			// ok
		case *uint8:
			// ok
		case uint8:
			// ok
		case *uint16:
			// ok
		case uint16:
			// ok
		case *uint32:
			// ok
		case uint32:
			// ok
		case *uint64:
			// ok
		case uint64:
			// ok
		default:
			return fmt.Errorf("parameter args[%#v] must be one of the allowed types: *string, *float64; received %#v (a %T)", idx_090772, v, v)
		}
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateTryFindParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateTasks_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTasks_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewTasksParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

