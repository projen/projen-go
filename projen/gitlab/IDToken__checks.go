//go:build !no_runtime_type_checking

package gitlab

import (
	"fmt"
)

func (j *jsiiProxy_IDToken) validateSetAudParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *[]*string; received %#v (a %T)", val, val)
	}

	return nil
}

