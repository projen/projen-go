package release

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/release/internal"
)

// Manages releases (currently through GitHub workflows).
//
// By default, no branches are released. To add branches, call `addBranch()`.
// Experimental.
type Release interface {
	projen.Component
	// Location of build artifacts.
	// Experimental.
	ArtifactsDirectory() *string
	// Retrieve all release branch names.
	// Experimental.
	Branches() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Package publisher.
	// Experimental.
	Publisher() Publisher
	// Adds a release branch.
	//
	// It is a git branch from which releases are published. If a project has more than one release
	// branch, we require that `majorVersion` is also specified for the primary branch in order to
	// ensure branches always release the correct version.
	// Experimental.
	AddBranch(branch *string, options *BranchOptions)
	// Adds jobs to all release workflows.
	// Experimental.
	AddJobs(jobs *map[string]*workflows.Job)
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
}

// The jsii proxy struct for Release
type jsiiProxy_Release struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Release) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Branches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Publisher() Publisher {
	var returns Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}


// Experimental.
func NewRelease(project github.GitHubProject, options *ReleaseOptions) Release {
	_init_.Initialize()

	if err := validateNewReleaseParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Release{}

	_jsii_.Create(
		"projen.release.Release",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRelease_Override(r Release, project github.GitHubProject, options *ReleaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.release.Release",
		[]interface{}{project, options},
		r,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Release_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRelease_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.release.Release",
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
func Release_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRelease_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.release.Release",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the `Release` component of a project or `undefined` if the project does not have a Release component.
// Experimental.
func Release_Of(project github.GitHubProject) Release {
	_init_.Initialize()

	if err := validateRelease_OfParameters(project); err != nil {
		panic(err)
	}
	var returns Release

	_jsii_.StaticInvoke(
		"projen.release.Release",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func Release_ANTI_TAMPER_CMD() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.release.Release",
		"ANTI_TAMPER_CMD",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Release) AddBranch(branch *string, options *BranchOptions) {
	if err := r.validateAddBranchParameters(branch, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addBranch",
		[]interface{}{branch, options},
	)
}

func (r *jsiiProxy_Release) AddJobs(jobs *map[string]*workflows.Job) {
	if err := r.validateAddJobsParameters(jobs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (r *jsiiProxy_Release) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

