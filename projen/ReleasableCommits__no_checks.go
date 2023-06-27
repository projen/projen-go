//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateReleasableCommits_ExecParameters(cmd *string) error {
	return nil
}

func validateReleasableCommits_OfTypeParameters(types *[]*string) error {
	return nil
}

func (j *jsiiProxy_ReleasableCommits) validateSetCmdParameters(val *string) error {
	return nil
}

