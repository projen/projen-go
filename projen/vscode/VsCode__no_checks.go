//go:build no_runtime_type_checking

package vscode

// Building without runtime type checking enabled, so all the below just return nil

func validateVsCode_IsComponentParameters(x interface{}) error {
	return nil
}

func validateVsCode_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewVsCodeParameters(project projen.Project) error {
	return nil
}

