//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (x *jsiiProxy_XmlFile) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (x *jsiiProxy_XmlFile) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (x *jsiiProxy_XmlFile) validateAddToArrayParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (x *jsiiProxy_XmlFile) validateSynthesizeContentParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_XmlFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_XmlFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewXmlFileParameters(project Project, filePath *string, options *XmlFileOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

