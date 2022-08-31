//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func validateNewMavenCompileParameters(project projen.Project, pom Pom, options *MavenCompileOptions) error {
	return nil
}

