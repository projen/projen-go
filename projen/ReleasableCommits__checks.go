//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func validateReleasableCommits_ExecParameters(cmd *string) error {
	if cmd == nil {
		return fmt.Errorf("parameter cmd is required, but nil was provided")
	}

	return nil
}

func validateReleasableCommits_OfTypeParameters(types *[]*string) error {
	if types == nil {
		return fmt.Errorf("parameter types is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ReleasableCommits) validateSetCmdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

