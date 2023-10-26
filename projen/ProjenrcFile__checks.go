//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateProjenrcFile_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateProjenrcFile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateProjenrcFile_OfParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func validateNewProjenrcFileParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

