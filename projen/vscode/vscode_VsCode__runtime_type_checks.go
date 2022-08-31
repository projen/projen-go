//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package vscode

import (
	"fmt"

	"github.com/projen/projen-go/projen"
)

func validateNewVsCodeParameters(project projen.Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

