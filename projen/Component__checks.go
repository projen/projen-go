//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateComponent_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateComponent_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewComponentParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

