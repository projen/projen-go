//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen/github/workflows"
)

func (t *jsiiProxy_TaskWorkflow) validateAddJobParameters(id *string, job interface{}) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if job == nil {
		return fmt.Errorf("parameter job is required, but nil was provided")
	}
	switch job.(type) {
	case *workflows.JobCallingReusableWorkflow:
		job := job.(*workflows.JobCallingReusableWorkflow)
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	case workflows.JobCallingReusableWorkflow:
		job_ := job.(workflows.JobCallingReusableWorkflow)
		job := &job_
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	case *workflows.Job:
		job := job.(*workflows.Job)
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	case workflows.Job:
		job_ := job.(workflows.Job)
		job := &job_
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(job) {
			return fmt.Errorf("parameter job must be one of the allowed types: *workflows.JobCallingReusableWorkflow, *workflows.Job; received %#v (a %T)", job, job)
		}
	}

	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateAddJobsParameters(jobs *map[string]interface{}) error {
	if jobs == nil {
		return fmt.Errorf("parameter jobs is required, but nil was provided")
	}
	for idx_5d9a17, v := range *jobs {
		switch v.(type) {
		case *workflows.JobCallingReusableWorkflow:
			v := v.(*workflows.JobCallingReusableWorkflow)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		case workflows.JobCallingReusableWorkflow:
			v_ := v.(workflows.JobCallingReusableWorkflow)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		case *workflows.Job:
			v := v.(*workflows.Job)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		case workflows.Job:
			v_ := v.(workflows.Job)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter jobs[%#v] must be one of the allowed types: *workflows.JobCallingReusableWorkflow, *workflows.Job; received %#v (a %T)", idx_5d9a17, v, v)
			}
		}
	}

	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateGetJobParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateOnParameters(events *workflows.Triggers) error {
	if events == nil {
		return fmt.Errorf("parameter events is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(events, func() string { return "parameter events" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateRemoveJobParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateUpdateJobParameters(id *string, job interface{}) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if job == nil {
		return fmt.Errorf("parameter job is required, but nil was provided")
	}
	switch job.(type) {
	case *workflows.JobCallingReusableWorkflow:
		job := job.(*workflows.JobCallingReusableWorkflow)
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	case workflows.JobCallingReusableWorkflow:
		job_ := job.(workflows.JobCallingReusableWorkflow)
		job := &job_
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	case *workflows.Job:
		job := job.(*workflows.Job)
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	case workflows.Job:
		job_ := job.(workflows.Job)
		job := &job_
		if err := _jsii_.ValidateStruct(job, func() string { return "parameter job" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(job) {
			return fmt.Errorf("parameter job must be one of the allowed types: *workflows.JobCallingReusableWorkflow, *workflows.Job; received %#v (a %T)", job, job)
		}
	}

	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateUpdateJobsParameters(jobs *map[string]interface{}) error {
	if jobs == nil {
		return fmt.Errorf("parameter jobs is required, but nil was provided")
	}
	for idx_5d9a17, v := range *jobs {
		switch v.(type) {
		case *workflows.JobCallingReusableWorkflow:
			v := v.(*workflows.JobCallingReusableWorkflow)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		case workflows.JobCallingReusableWorkflow:
			v_ := v.(workflows.JobCallingReusableWorkflow)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		case *workflows.Job:
			v := v.(*workflows.Job)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		case workflows.Job:
			v_ := v.(workflows.Job)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
				return err
			}
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter jobs[%#v] must be one of the allowed types: *workflows.JobCallingReusableWorkflow, *workflows.Job; received %#v (a %T)", idx_5d9a17, v, v)
			}
		}
	}

	return nil
}

func validateTaskWorkflow_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTaskWorkflow_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewTaskWorkflowParameters(github GitHub, options *TaskWorkflowOptions) error {
	if github == nil {
		return fmt.Errorf("parameter github is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

