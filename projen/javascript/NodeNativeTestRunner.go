package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Configures Node's built-in test runner (`node --test`).
//
// Owns the generated Node.js configuration file, the "test", "test:update"
// and "test:watch" tasks, and the reporters and test match patterns they use.
// Experimental.
type NodeNativeTestRunner interface {
	projen.Component
	// The generated Node.js configuration file.
	// Experimental.
	ConfigFile() NodeConfigFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	// Returns `NodeProject`, use interface conversion if needed.
	Project() projen.Project
	// Adds a reporter, or updates its destination if one with the same name is already configured.
	// Default: Destination.STDOUT
	//
	// Experimental.
	AddReporter(name *string, destination Destination)
	// Adds a test match pattern.
	// Experimental.
	AddTestMatch(pattern *string)
	// Lists the configured reporters, in the order they were added.
	// Experimental.
	ListReporters() *[]*NodeReporter
	// Called once, right after `postSynthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// It is also skipped when post-synthesis steps are disabled, e.g. `--no-post` or `PROJEN_DISABLE_POST`.
	// Use it for one-off setup that can be turned off by the user, like running a task to give the user immediate
	// feedback on their new project. Order across components is not guaranteed.
	// Experimental.
	PostProjectCreation(initProject *projen.InitProject)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Called once, right after `synthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// Use it for deterministic, one-off file generation. Order across components is not guaranteed.
	// Experimental.
	ProjectCreation(initProject *projen.InitProject)
	// Removes a reporter, if configured.
	// Experimental.
	RemoveReporter(name *string)
	// Removes a test match pattern, if configured.
	// Experimental.
	RemoveTestMatch(pattern *string)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for NodeNativeTestRunner
type jsiiProxy_NodeNativeTestRunner struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NodeNativeTestRunner) ConfigFile() NodeConfigFile {
	var returns NodeConfigFile
	_jsii_.Get(
		j,
		"configFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeNativeTestRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeNativeTestRunner) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodeNativeTestRunner(scope constructs.IConstruct, options *NodeNativeTestRunnerOptions) NodeNativeTestRunner {
	_init_.Initialize()

	if err := validateNewNodeNativeTestRunnerParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodeNativeTestRunner{}

	_jsii_.Create(
		"projen.javascript.NodeNativeTestRunner",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNodeNativeTestRunner_Override(n NodeNativeTestRunner, scope constructs.IConstruct, options *NodeNativeTestRunnerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NodeNativeTestRunner",
		[]interface{}{scope, options},
		n,
	)
}

// Test whether the given construct is a component.
// Experimental.
func NodeNativeTestRunner_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodeNativeTestRunner_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NodeNativeTestRunner",
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
func NodeNativeTestRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodeNativeTestRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NodeNativeTestRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the singleton NodeNativeTestRunner component of a project or undefined if there is none.
// Experimental.
func NodeNativeTestRunner_Of(project projen.Project) NodeNativeTestRunner {
	_init_.Initialize()

	if err := validateNodeNativeTestRunner_OfParameters(project); err != nil {
		panic(err)
	}
	var returns NodeNativeTestRunner

	_jsii_.StaticInvoke(
		"projen.javascript.NodeNativeTestRunner",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeNativeTestRunner) AddReporter(name *string, destination Destination) {
	if err := n.validateAddReporterParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addReporter",
		[]interface{}{name, destination},
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) AddTestMatch(pattern *string) {
	if err := n.validateAddTestMatchParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addTestMatch",
		[]interface{}{pattern},
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) ListReporters() *[]*NodeReporter {
	var returns *[]*NodeReporter

	_jsii_.Invoke(
		n,
		"listReporters",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeNativeTestRunner) PostProjectCreation(initProject *projen.InitProject) {
	if err := n.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) ProjectCreation(initProject *projen.InitProject) {
	if err := n.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) RemoveReporter(name *string) {
	if err := n.validateRemoveReporterParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"removeReporter",
		[]interface{}{name},
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) RemoveTestMatch(pattern *string) {
	if err := n.validateRemoveTestMatchParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"removeTestMatch",
		[]interface{}{pattern},
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeNativeTestRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeNativeTestRunner) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		n,
		"with",
		args,
		&returns,
	)

	return returns
}

