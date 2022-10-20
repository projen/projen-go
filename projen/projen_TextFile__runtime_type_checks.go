//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_TextFile) validateAddLineParameters(line *string) error {
	if line == nil {
		return fmt.Errorf("parameter line is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TextFile) validateSynthesizeContentParameters(_arg IResolver) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TextFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TextFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTextFileParameters(project Project, filePath *string, options *TextFileOptions) error {
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

