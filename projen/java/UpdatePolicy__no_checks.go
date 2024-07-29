//go:build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func validateUpdatePolicy_IntervalParameters(minutes *float64) error {
	return nil
}

