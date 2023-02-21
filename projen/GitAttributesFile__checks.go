//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func (g *jsiiProxy_GitAttributesFile) validateAddAttributesParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateAddLfsPatternParameters(glob *string) error {
	if glob == nil {
		return fmt.Errorf("parameter glob is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateSynthesizeContentParameters(_arg IResolver) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GitAttributesFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GitAttributesFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGitAttributesFileParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

