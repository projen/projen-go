//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_GitAttributesFile) validateAddAttributesParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateAddLfsPatternParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateRemoveAttributesParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateSynthesizeContentParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func validateGitAttributesFile_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGitAttributesFile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GitAttributesFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GitAttributesFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGitAttributesFileParameters(scope constructs.IConstruct, options *GitAttributesFileOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

