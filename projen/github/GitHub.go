package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Experimental.
type GitHub interface {
	projen.Component
	// Experimental.
	Actions() GitHubActionsProvider
	// Whether downloading from LFS is enabled for this GitHub project.
	// Experimental.
	DownloadLfs() *bool
	// The `Mergify` configured on this repository.
	//
	// This is `undefined` if Mergify
	// was not enabled when creating the repository.
	// Experimental.
	Mergify() Mergify
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// GitHub API authentication method used by projen workflows.
	// Experimental.
	ProjenCredentials() GithubCredentials
	// All workflows.
	// Experimental.
	Workflows() *[]GithubWorkflow
	// Are workflows enabled?
	// Experimental.
	WorkflowsEnabled() *bool
	// Experimental.
	AddDependabot(options *DependabotOptions) Dependabot
	// Experimental.
	AddPullRequestTemplate(content ...*string) PullRequestTemplate
	// Adds a workflow to the project.
	//
	// Returns: a GithubWorkflow instance.
	// Experimental.
	AddWorkflow(name *string) GithubWorkflow
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Finds a GitHub workflow by name.
	//
	// Returns `undefined` if the workflow cannot be found.
	// Experimental.
	TryFindWorkflow(name *string) GithubWorkflow
}

// The jsii proxy struct for GitHub
type jsiiProxy_GitHub struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_GitHub) Actions() GitHubActionsProvider {
	var returns GitHubActionsProvider
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) DownloadLfs() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"downloadLfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) Mergify() Mergify {
	var returns Mergify
	_jsii_.Get(
		j,
		"mergify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) ProjenCredentials() GithubCredentials {
	var returns GithubCredentials
	_jsii_.Get(
		j,
		"projenCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) Workflows() *[]GithubWorkflow {
	var returns *[]GithubWorkflow
	_jsii_.Get(
		j,
		"workflows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) WorkflowsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"workflowsEnabled",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHub(project projen.Project, options *GitHubOptions) GitHub {
	_init_.Initialize()

	if err := validateNewGitHubParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHub{}

	_jsii_.Create(
		"projen.github.GitHub",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHub_Override(g GitHub, project projen.Project, options *GitHubOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.GitHub",
		[]interface{}{project, options},
		g,
	)
}

// Test whether the given construct is a component.
// Experimental.
func GitHub_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitHub_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.GitHub",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func GitHub_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitHub_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.GitHub",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the `GitHub` component of a project or `undefined` if the project does not have a GitHub component.
// Experimental.
func GitHub_Of(project projen.Project) GitHub {
	_init_.Initialize()

	if err := validateGitHub_OfParameters(project); err != nil {
		panic(err)
	}
	var returns GitHub

	_jsii_.StaticInvoke(
		"projen.github.GitHub",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHub) AddDependabot(options *DependabotOptions) Dependabot {
	if err := g.validateAddDependabotParameters(options); err != nil {
		panic(err)
	}
	var returns Dependabot

	_jsii_.Invoke(
		g,
		"addDependabot",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHub) AddPullRequestTemplate(content ...*string) PullRequestTemplate {
	args := []interface{}{}
	for _, a := range content {
		args = append(args, a)
	}

	var returns PullRequestTemplate

	_jsii_.Invoke(
		g,
		"addPullRequestTemplate",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHub) AddWorkflow(name *string) GithubWorkflow {
	if err := g.validateAddWorkflowParameters(name); err != nil {
		panic(err)
	}
	var returns GithubWorkflow

	_jsii_.Invoke(
		g,
		"addWorkflow",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHub) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHub) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHub) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHub) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHub) TryFindWorkflow(name *string) GithubWorkflow {
	if err := g.validateTryFindWorkflowParameters(name); err != nil {
		panic(err)
	}
	var returns GithubWorkflow

	_jsii_.Invoke(
		g,
		"tryFindWorkflow",
		[]interface{}{name},
		&returns,
	)

	return returns
}

