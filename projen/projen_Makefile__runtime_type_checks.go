//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_Makefile) validateAddAllParameters(target *string) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_Makefile) validateAddRuleParameters(rule *Rule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_Makefile) validateAddRulesParameters(rules *[]*Rule) error {
	for idx_6c621d, v := range *rules {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter rules[%#v]", idx_6c621d) }); err != nil {
			return err
		}
	}

	return nil
}

func (m *jsiiProxy_Makefile) validateSynthesizeContentParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Makefile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Makefile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewMakefileParameters(project Project, filePath *string, options *MakefileOptions) error {
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

