package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/cdk/internal"
)

// Experimental.
type IntegrationTestBase interface {
	projen.Component
	// Synthesizes the integration test and compares against a local copy (runs during build).
	// Experimental.
	AssertTask() projen.Task
	// Deploy the integration test and update the snapshot upon success.
	// Experimental.
	DeployTask() projen.Task
	// Integration test name.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Snapshot output directory.
	// Experimental.
	SnapshotDir() *string
	// Just update snapshot (without deployment).
	// Experimental.
	SnapshotTask() projen.Task
	// Temporary directory for each integration test.
	// Experimental.
	TmpDir() *string
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

// The jsii proxy struct for IntegrationTestBase
type jsiiProxy_IntegrationTestBase struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_IntegrationTestBase) AssertTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"assertTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) DeployTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"deployTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) SnapshotDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) SnapshotTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"snapshotTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestBase) TmpDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tmpDir",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTestBase_Override(i IntegrationTestBase, project projen.Project, options *IntegrationTestBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.IntegrationTestBase",
		[]interface{}{project, options},
		i,
	)
}

// Test whether the given construct is a component.
// Experimental.
func IntegrationTestBase_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTestBase_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk.IntegrationTestBase",
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
func IntegrationTestBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTestBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk.IntegrationTestBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationTestBase) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestBase) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestBase) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

