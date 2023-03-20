//go:build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrc_OfParameters(project projen.Project) error {
	return nil
}

func validateNewProjenrcParameters(project projen.Project, pom Pom, options *ProjenrcOptions) error {
	return nil
}

