//go:build !no_runtime_type_checking

package biomeconfig

import (
	"fmt"
)

func (j *jsiiProxy_IHook) validateSetStableResultParameters(val interface{}) error {
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case *[]*float64:
		// ok
	case []*float64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *bool, *[]*float64; received %#v (a %T)", val, val)
	}

	return nil
}

