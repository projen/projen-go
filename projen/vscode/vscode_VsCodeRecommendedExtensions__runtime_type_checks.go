//go:build !no_runtime_type_checking

package vscode

import (
	"fmt"
)

func validateNewVsCodeRecommendedExtensionsParameters(vscode VsCode) error {
	if vscode == nil {
		return fmt.Errorf("parameter vscode is required, but nil was provided")
	}

	return nil
}

