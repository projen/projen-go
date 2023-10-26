package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type Version interface {
	Component
	// Experimental.
	BumpTask() Task
	// The name of the changelog file (under `artifactsDirectory`).
	// Experimental.
	ChangelogFileName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// The name of the file that contains the release tag (under `artifactsDirectory`).
	// Experimental.
	ReleaseTagFileName() *string
	// Experimental.
	UnbumpTask() Task
	// The name of the file that contains the version (under `artifactsDirectory`).
	// Experimental.
	VersionFileName() *string
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

// The jsii proxy struct for Version
type jsiiProxy_Version struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Version) BumpTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"bumpTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) ChangelogFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changelogFileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) ReleaseTagFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTagFileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) UnbumpTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"unbumpTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) VersionFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionFileName",
		&returns,
	)
	return returns
}


// Experimental.
func NewVersion(project Project, options *VersionOptions) Version {
	_init_.Initialize()

	if err := validateNewVersionParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Version{}

	_jsii_.Create(
		"projen.Version",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewVersion_Override(v Version, project Project, options *VersionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Version",
		[]interface{}{project, options},
		v,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Version_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVersion_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Version",
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
func Version_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Version",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Version_STANDARD_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Version",
		"STANDARD_VERSION",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_Version) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Version) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Version) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Version) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

