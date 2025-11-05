//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

func (c *jsiiProxy_Cdk8sPythonApp) validateAddDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateAddDevDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateAddGitIgnoreParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateAddPackageIgnoreParameters(_pattern *string) error {
	if _pattern == nil {
		return fmt.Errorf("parameter _pattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateAddTipParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateAnnotateGeneratedParameters(_glob *string) error {
	if _glob == nil {
		return fmt.Errorf("parameter _glob is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateRemoveTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateRunTaskCommandParameters(task projen.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateTryFindFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateTryFindJsonFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateTryFindObjectFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sPythonApp) validateTryRemoveFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateCdk8sPythonApp_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCdk8sPythonApp_IsProjectParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCdk8sPythonApp_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewCdk8sPythonAppParameters(options *Cdk8sPythonOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

