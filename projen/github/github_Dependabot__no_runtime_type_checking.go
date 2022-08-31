//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Dependabot) validateAddIgnoreParameters(dependencyName *string) error {
	return nil
}

func validateNewDependabotParameters(github GitHub, options *DependabotOptions) error {
	return nil
}

