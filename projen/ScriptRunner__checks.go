//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_ScriptRunner) validateAttachParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScriptRunner) validateConfigForParameters(entrypoint *string) error {
	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScriptRunner) validateTryAttachParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateScriptRunner_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateScriptRunner_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

