//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_Version) validateEnvForBranchParameters(branchOptions *VersionBranchOptions) error {
	if branchOptions == nil {
		return fmt.Errorf("parameter branchOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(branchOptions, func() string { return "parameter branchOptions" }); err != nil {
		return err
	}

	return nil
}

func validateVersion_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateVersion_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewVersionParameters(scope constructs.IConstruct, options *VersionOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

