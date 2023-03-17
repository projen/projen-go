//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Jest) validateAddIgnorePatternParameters(pattern *string) error {
	return nil
}

func (j *jsiiProxy_Jest) validateAddReporterParameters(reporter interface{}) error {
	return nil
}

func (j *jsiiProxy_Jest) validateAddSetupFileParameters(file *string) error {
	return nil
}

func (j *jsiiProxy_Jest) validateAddSetupFileAfterEnvParameters(file *string) error {
	return nil
}

func (j *jsiiProxy_Jest) validateAddSnapshotResolverParameters(file *string) error {
	return nil
}

func (j *jsiiProxy_Jest) validateAddTestMatchParameters(pattern *string) error {
	return nil
}

func (j *jsiiProxy_Jest) validateAddWatchIgnorePatternParameters(pattern *string) error {
	return nil
}

func validateJest_OfParameters(project projen.Project) error {
	return nil
}

func validateNewJestParameters(project NodeProject, options *JestOptions) error {
	return nil
}

