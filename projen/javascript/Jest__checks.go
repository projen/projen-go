//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (j *jsiiProxy_Jest) validateAddIgnorePatternParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddReporterParameters(reporter JestReporter) error {
	if reporter == nil {
		return fmt.Errorf("parameter reporter is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddSetupFileParameters(file *string) error {
	if file == nil {
		return fmt.Errorf("parameter file is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddSetupFileAfterEnvParameters(file *string) error {
	if file == nil {
		return fmt.Errorf("parameter file is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddSnapshotResolverParameters(file *string) error {
	if file == nil {
		return fmt.Errorf("parameter file is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddTestMatchParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddWatchIgnorePatternParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateJest_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateJest_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateJest_OfParameters(project projen.Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func validateNewJestParameters(project NodeProject, options *JestOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

