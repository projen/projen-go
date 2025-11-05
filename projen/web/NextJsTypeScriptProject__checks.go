//go:build !no_runtime_type_checking

package web

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
)

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddBinsParameters(bins *map[string]*string) error {
	if bins == nil {
		return fmt.Errorf("parameter bins is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddFieldsParameters(fields *map[string]interface{}) error {
	if fields == nil {
		return fmt.Errorf("parameter fields is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddGitIgnoreParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddPackageIgnoreParameters(_pattern *string) error {
	if _pattern == nil {
		return fmt.Errorf("parameter _pattern is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddScriptsParameters(scripts *map[string]*string) error {
	if scripts == nil {
		return fmt.Errorf("parameter scripts is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAddTipParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateAnnotateGeneratedParameters(_glob *string) error {
	if _glob == nil {
		return fmt.Errorf("parameter _glob is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateHasScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateRemoveScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateRemoveTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateRenderWorkflowSetupParameters(options *javascript.RenderWorkflowSetupOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateRunTaskCommandParameters(task projen.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateSetScriptParameters(name *string, command *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateTryFindFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateTryFindJsonFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateTryFindObjectFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NextJsTypeScriptProject) validateTryRemoveFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateNextJsTypeScriptProject_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNextJsTypeScriptProject_IsProjectParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNextJsTypeScriptProject_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewNextJsTypeScriptProjectParameters(options *NextJsTypeScriptProjectOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

