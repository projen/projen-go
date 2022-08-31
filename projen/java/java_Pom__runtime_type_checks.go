//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package java

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (p *jsiiProxy_Pom) validateAddDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Pom) validateAddPluginParameters(spec *string, options *PluginOptions) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Pom) validateAddPropertyParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Pom) validateAddRepositoryParameters(repository *MavenRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(repository, func() string { return "parameter repository" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Pom) validateAddTestDependencyParameters(spec *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	return nil
}

func validateNewPomParameters(project projen.Project, options *PomOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

