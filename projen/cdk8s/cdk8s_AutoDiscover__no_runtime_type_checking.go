//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateNewAutoDiscoverParameters(project projen.Project, options *AutoDiscoverOptions) error {
	return nil
}

