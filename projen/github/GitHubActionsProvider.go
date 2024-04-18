package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Manage the versions used for GitHub Actions used in steps.
// Experimental.
type GitHubActionsProvider interface {
	// Resolve an action name to the version that should be used, taking into account any overrides.
	// Experimental.
	Get(action *string) *string
	// Define an override for a given action.
	//
	// Specify the action name without a version to override all usages of the action.
	// You can also override a specific action version, by providing the version string.
	// Specific overrides take precedence over overrides without a version.
	//
	// If an override for the same action name is set multiple times, the last override is used.
	//
	// Example:
	//   // Force any use of `actions/checkout` to use a pin a specific commit
	//   project.github.actions.set("actions/checkout", "actions/checkout@aaaaaa");
	//
	//   // But pin usage of `v4` to a different commit
	//   project.github.actions.set("actions/checkout@v4", "actions/checkout@ffffff");
	//
	// Experimental.
	Set(action *string, override *string)
}

// The jsii proxy struct for GitHubActionsProvider
type jsiiProxy_GitHubActionsProvider struct {
	_ byte // padding
}

// Experimental.
func NewGitHubActionsProvider() GitHubActionsProvider {
	_init_.Initialize()

	j := jsiiProxy_GitHubActionsProvider{}

	_jsii_.Create(
		"projen.github.GitHubActionsProvider",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubActionsProvider_Override(g GitHubActionsProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.GitHubActionsProvider",
		nil, // no parameters
		g,
	)
}

func (g *jsiiProxy_GitHubActionsProvider) Get(action *string) *string {
	if err := g.validateGetParameters(action); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubActionsProvider) Set(action *string, override *string) {
	if err := g.validateSetParameters(action, override); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"set",
		[]interface{}{action, override},
	)
}

