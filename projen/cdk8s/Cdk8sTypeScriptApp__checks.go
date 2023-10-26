//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
)

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddBinsParameters(bins *map[string]*string) error {
	if bins == nil {
		return fmt.Errorf("parameter bins is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddFieldsParameters(fields *map[string]interface{}) error {
	if fields == nil {
		return fmt.Errorf("parameter fields is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddGitIgnoreParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddPackageIgnoreParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddScriptsParameters(scripts *map[string]*string) error {
	if scripts == nil {
		return fmt.Errorf("parameter scripts is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAddTipParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateAnnotateGeneratedParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateHasScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateRemoveScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateRemoveTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateRenderWorkflowSetupParameters(options *javascript.RenderWorkflowSetupOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateRunTaskCommandParameters(task projen.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateSetScriptParameters(name *string, command *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateTryFindFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateTryFindJsonFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateTryFindObjectFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) validateTryRemoveFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateCdk8sTypeScriptApp_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCdk8sTypeScriptApp_IsProjectParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCdk8sTypeScriptApp_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewCdk8sTypeScriptAppParameters(options *Cdk8sTypeScriptAppOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

