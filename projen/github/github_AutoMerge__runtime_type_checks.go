//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_AutoMerge) validateAddConditionsLaterParameters(later IAddConditionsLater) error {
	if later == nil {
		return fmt.Errorf("parameter later is required, but nil was provided")
	}

	return nil
}

func validateNewAutoMergeParameters(github GitHub, options *AutoMergeOptions) error {
	if github == nil {
		return fmt.Errorf("parameter github is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

