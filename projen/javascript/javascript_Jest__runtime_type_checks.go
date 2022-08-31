//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package javascript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_Jest) validateAddIgnorePatternParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Jest) validateAddReporterParameters(reporter interface{}) error {
	if reporter == nil {
		return fmt.Errorf("parameter reporter is required, but nil was provided")
	}
	switch reporter.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *map[string]interface{}:
		// ok
	case map[string]interface{}:
		// ok
	default:
		return fmt.Errorf("parameter reporter must be one of the allowed types: *string, *map[string]interface{}; received %#v (a %T)", reporter, reporter)
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

func validateNewJestParameters(project NodeProject, options *JestOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

