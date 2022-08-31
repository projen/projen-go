//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package cdktf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
)

func (c *jsiiProxy_ConstructLibraryCdktf) validateAddBinsParameters(bins *map[string]*string) error {
	if bins == nil {
		return fmt.Errorf("parameter bins is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateAddFieldsParameters(fields *map[string]interface{}) error {
	if fields == nil {
		return fmt.Errorf("parameter fields is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateAddGitIgnoreParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateAddPackageIgnoreParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateAddTipParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateAnnotateGeneratedParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateHasScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateRemoveScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateRemoveTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateRenderWorkflowSetupParameters(options *javascript.RenderWorkflowSetupOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateRunTaskCommandParameters(task projen.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateSetScriptParameters(name *string, command *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateTryFindFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateTryFindJsonFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateTryFindObjectFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructLibraryCdktf) validateTryRemoveFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateNewConstructLibraryCdktfParameters(options *ConstructLibraryCdktfOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

