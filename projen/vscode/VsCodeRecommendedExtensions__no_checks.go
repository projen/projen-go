//go:build no_runtime_type_checking

package vscode

// Building without runtime type checking enabled, so all the below just return nil

func validateVsCodeRecommendedExtensions_IsComponentParameters(x interface{}) error {
	return nil
}

func validateVsCodeRecommendedExtensions_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewVsCodeRecommendedExtensionsParameters(vscode VsCode) error {
	return nil
}

