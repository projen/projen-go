//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateRenovatebot_IsComponentParameters(x interface{}) error {
	return nil
}

func validateRenovatebot_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRenovatebotParameters(project Project, options *RenovatebotOptions) error {
	return nil
}

