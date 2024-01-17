//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (n *jsiiProxy_NodePackage) validateAddBinParameters(bins *map[string]*string) error {
	if bins == nil {
		return fmt.Errorf("parameter bins is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateAddEngineParameters(engine *string, version *string) error {
	if engine == nil {
		return fmt.Errorf("parameter engine is required, but nil was provided")
	}

	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateAddFieldParameters(name *string, value interface{}) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateAddVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateHasScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateRemoveScriptParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateSetScriptParameters(name *string, command *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NodePackage) validateTryResolveDependencyVersionParameters(dependencyName *string) error {
	if dependencyName == nil {
		return fmt.Errorf("parameter dependencyName is required, but nil was provided")
	}

	return nil
}

func validateNodePackage_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNodePackage_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNodePackage_OfParameters(project projen.Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func validateNewNodePackageParameters(project projen.Project, options *NodePackageOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

