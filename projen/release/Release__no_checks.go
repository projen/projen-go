//go:build no_runtime_type_checking

package release

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Release) validateAddBranchParameters(branch *string, options *BranchOptions) error {
	return nil
}

func (r *jsiiProxy_Release) validateAddJobsParameters(jobs *map[string]*workflows.Job) error {
	return nil
}

func validateRelease_IsComponentParameters(x interface{}) error {
	return nil
}

func validateRelease_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRelease_OfParameters(project github.GitHubProject) error {
	return nil
}

func validateNewReleaseParameters(project github.GitHubProject, options *ReleaseOptions) error {
	return nil
}

