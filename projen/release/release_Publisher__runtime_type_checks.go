//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package release

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

func (p *jsiiProxy_Publisher) validateAddGitHubPrePublishingStepsParameters(steps *[]*workflows.JobStep) error {
	for idx_b7595e, v := range *steps {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter steps[%#v]", idx_b7595e) }); err != nil {
			return err
		}
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToGitParameters(options *GitPublishOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToGitHubReleasesParameters(options *GitHubReleasesPublishOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToGoParameters(options *GoPublishOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToMavenParameters(options *MavenPublishOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToNpmParameters(options *NpmPublishOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToNugetParameters(options *NugetPublishOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Publisher) validatePublishToPyPiParameters(options *PyPiPublishOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewPublisherParameters(project projen.Project, options *PublisherOptions) error {
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

