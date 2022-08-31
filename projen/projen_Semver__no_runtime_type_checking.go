//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateSemver_CaretParameters(version *string) error {
	return nil
}

func validateSemver_OfParameters(spec *string) error {
	return nil
}

func validateSemver_PinnedParameters(version *string) error {
	return nil
}

func validateSemver_TildeParameters(version *string) error {
	return nil
}

