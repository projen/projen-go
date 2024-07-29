//go:build !no_runtime_type_checking

package java

import (
	"fmt"
)

func validateUpdatePolicy_IntervalParameters(minutes *float64) error {
	if minutes == nil {
		return fmt.Errorf("parameter minutes is required, but nil was provided")
	}

	return nil
}

