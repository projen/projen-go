//go:build no_runtime_type_checking

package release

// Building without runtime type checking enabled, so all the below just return nil

func validateReleaseTrigger_ManualParameters(options *ManualReleaseOptions) error {
	return nil
}

func validateReleaseTrigger_ScheduledParameters(options *ScheduledReleaseOptions) error {
	return nil
}

