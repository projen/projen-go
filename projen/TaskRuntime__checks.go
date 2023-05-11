//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func (t *jsiiProxy_TaskRuntime) validateRunTaskParameters(name *string, args *[]interface{}) error {
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

func (t *jsiiProxy_TaskRuntime) validateTryFindTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewTaskRuntimeParameters(workdir *string) error {
	if workdir == nil {
		return fmt.Errorf("parameter workdir is required, but nil was provided")
	}

	return nil
}

