//go:build no_runtime_type_checking

package release

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Publisher) validateAddGitHubPrePublishingStepsParameters(steps *[]*workflows.JobStep) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToGitParameters(options *GitPublishOptions) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToGitHubReleasesParameters(options *GitHubReleasesPublishOptions) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToGoParameters(options *GoPublishOptions) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToMavenParameters(options *MavenPublishOptions) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToNpmParameters(options *NpmPublishOptions) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToNugetParameters(options *NugetPublishOptions) error {
	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToPyPiParameters(options *PyPiPublishOptions) error {
	return nil
}

func validateNewPublisherParameters(project projen.Project, options *PublisherOptions) error {
	return nil
}

