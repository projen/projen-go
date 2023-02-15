//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHubActionsProvider) validateGetParameters(action *string) error {
	return nil
}

func (g *jsiiProxy_GitHubActionsProvider) validateSetParameters(action *string, override *string) error {
	return nil
}

