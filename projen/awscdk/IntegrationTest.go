package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/cdk"
)

// Cloud integration tests.
// Experimental.
type IntegrationTest interface {
	cdk.IntegrationTestBase
	// Synthesizes the integration test and compares against a local copy (runs during build).
	// Experimental.
	AssertTask() projen.Task
	// Deploy the integration test and update the snapshot upon success.
	// Experimental.
	DeployTask() projen.Task
	// Destroy the integration test resources.
	// Experimental.
	DestroyTask() projen.Task
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
	// The watch task.
	// Experimental.
	WatchTask() projen.Task
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

// The jsii proxy struct for IntegrationTest
type jsiiProxy_IntegrationTest struct {
	internal.Type__cdkIntegrationTestBase
}

func (j *jsiiProxy_IntegrationTest) AssertTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"assertTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) DeployTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"deployTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) DestroyTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"destroyTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) SnapshotDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) SnapshotTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"snapshotTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) TmpDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tmpDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTest(project projen.Project, options *IntegrationTestOptions) IntegrationTest {
	_init_.Initialize()

	if err := validateNewIntegrationTestParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationTest{}

	_jsii_.Create(
		"projen.awscdk.IntegrationTest",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTest_Override(i IntegrationTest, project projen.Project, options *IntegrationTestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.IntegrationTest",
		[]interface{}{project, options},
		i,
	)
}

// Test whether the given construct is a component.
// Experimental.
func IntegrationTest_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTest_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.IntegrationTest",
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
func IntegrationTest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.IntegrationTest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationTest) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTest) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTest) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

