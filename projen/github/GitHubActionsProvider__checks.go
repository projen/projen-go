//go:build !no_runtime_type_checking

package github

import (
	"fmt"
)

func (g *jsiiProxy_GitHubActionsProvider) validateGetParameters(action *string) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitHubActionsProvider) validateSetParameters(action *string, override *string) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	if override == nil {
		return fmt.Errorf("parameter override is required, but nil was provided")
	}

	return nil
}

