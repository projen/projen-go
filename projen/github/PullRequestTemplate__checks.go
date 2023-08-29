//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (p *jsiiProxy_PullRequestTemplate) validateAddLineParameters(line *string) error {
	if line == nil {
		return fmt.Errorf("parameter line is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PullRequestTemplate) validateSynthesizeContentParameters(_arg projen.IResolver) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func validatePullRequestTemplate_OfParameters(project projen.Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PullRequestTemplate) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PullRequestTemplate) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPullRequestTemplateParameters(github GitHub, options *PullRequestTemplateOptions) error {
	if github == nil {
		return fmt.Errorf("parameter github is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

