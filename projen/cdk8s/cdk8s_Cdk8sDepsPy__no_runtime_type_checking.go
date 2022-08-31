//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateNewCdk8sDepsPyParameters(project projen.Project, options *Cdk8sDepsOptions) error {
	return nil
}

