//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package build

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildJobParameters(id *string, job *workflows.Job) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if job == nil {
		return fmt.Errorf("parameter job is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildJobCommandsParameters(id *string, commands *[]*string, options *AddPostBuildJobCommandsOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if commands == nil {
		return fmt.Errorf("parameter commands is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildJobTaskParameters(task projen.Task, options *AddPostBuildJobTaskOptions) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildStepsParameters(steps *[]*workflows.JobStep) error {
	for idx_b7595e, v := range *steps {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter steps[%#v]", idx_b7595e) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNewBuildWorkflowParameters(project projen.Project, options *BuildWorkflowOptions) error {
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

