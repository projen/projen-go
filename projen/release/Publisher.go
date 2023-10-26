package release

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/release/internal"
)

// Implements GitHub jobs for publishing modules to package managers.
//
// Under the hood, it uses https://github.com/aws/publib
// Experimental.
type Publisher interface {
	projen.Component
	// Experimental.
	ArtifactName() *string
	// Experimental.
	BuildJobId() *string
	// Experimental.
	Condition() *string
	// Deprecated: use `publibVersion`.
	JsiiReleaseVersion() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Experimental.
	PublibVersion() *string
	// Adds post publishing steps for the GitHub release job.
	// Experimental.
	AddGitHubPostPublishingSteps(steps ...*workflows.JobStep)
	// Adds pre publishing steps for the GitHub release job.
	// Experimental.
	AddGitHubPrePublishingSteps(steps ...*workflows.JobStep)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Publish to git.
	//
	// This includes generating a project-level changelog and release tags.
	// Experimental.
	PublishToGit(options *GitPublishOptions) projen.Task
	// Creates a GitHub Release.
	// Experimental.
	PublishToGitHubReleases(options *GitHubReleasesPublishOptions)
	// Adds a go publishing job.
	// Experimental.
	PublishToGo(options *GoPublishOptions)
	// Publishes artifacts from `java/**` to Maven.
	// Experimental.
	PublishToMaven(options *MavenPublishOptions)
	// Publishes artifacts from `js/**` to npm.
	// Experimental.
	PublishToNpm(options *NpmPublishOptions)
	// Publishes artifacts from `dotnet/**` to NuGet Gallery.
	// Experimental.
	PublishToNuget(options *NugetPublishOptions)
	// Publishes wheel artifacts from `python` to PyPI.
	// Experimental.
	PublishToPyPi(options *PyPiPublishOptions)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Publisher
type jsiiProxy_Publisher struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Publisher) ArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) BuildJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) JsiiReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsiiReleaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) PublibVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publibVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewPublisher(project projen.Project, options *PublisherOptions) Publisher {
	_init_.Initialize()

	if err := validateNewPublisherParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Publisher{}

	_jsii_.Create(
		"projen.release.Publisher",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPublisher_Override(p Publisher, project projen.Project, options *PublisherOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.release.Publisher",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Publisher_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublisher_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.release.Publisher",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Publisher_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublisher_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.release.Publisher",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Publisher_PUBLISH_GIT_TASK_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.release.Publisher",
		"PUBLISH_GIT_TASK_NAME",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Publisher) AddGitHubPostPublishingSteps(steps ...*workflows.JobStep) {
	if err := p.validateAddGitHubPostPublishingStepsParameters(&steps); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addGitHubPostPublishingSteps",
		args,
	)
}

func (p *jsiiProxy_Publisher) AddGitHubPrePublishingSteps(steps ...*workflows.JobStep) {
	if err := p.validateAddGitHubPrePublishingStepsParameters(&steps); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addGitHubPrePublishingSteps",
		args,
	)
}

func (p *jsiiProxy_Publisher) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publisher) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publisher) PublishToGit(options *GitPublishOptions) projen.Task {
	if err := p.validatePublishToGitParameters(options); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		p,
		"publishToGit",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publisher) PublishToGitHubReleases(options *GitHubReleasesPublishOptions) {
	if err := p.validatePublishToGitHubReleasesParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"publishToGitHubReleases",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_Publisher) PublishToGo(options *GoPublishOptions) {
	if err := p.validatePublishToGoParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"publishToGo",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_Publisher) PublishToMaven(options *MavenPublishOptions) {
	if err := p.validatePublishToMavenParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"publishToMaven",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_Publisher) PublishToNpm(options *NpmPublishOptions) {
	if err := p.validatePublishToNpmParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"publishToNpm",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_Publisher) PublishToNuget(options *NugetPublishOptions) {
	if err := p.validatePublishToNugetParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"publishToNuget",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_Publisher) PublishToPyPi(options *PyPiPublishOptions) {
	if err := p.validatePublishToPyPiParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"publishToPyPi",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_Publisher) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publisher) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

