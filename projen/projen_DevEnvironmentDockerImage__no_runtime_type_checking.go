//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateDevEnvironmentDockerImage_FromFileParameters(dockerFile *string) error {
	return nil
}

func validateDevEnvironmentDockerImage_FromImageParameters(image *string) error {
	return nil
}

