package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Represents a Node.js configuration file, loaded via `--experimental-config-file`.
// See: https://nodejs.org/api/cli.html#configuration-via-nodeconfig
//
// Experimental.
type NodeConfigFile interface {
	projen.Component
	// Escape hatch for the generated configuration file.
	// Experimental.
	Config() *NodeConfigSchema
	// The underlying Node.js configuration file.
	// Experimental.
	File() projen.JsonFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for NodeConfigFile
type jsiiProxy_NodeConfigFile struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NodeConfigFile) Config() *NodeConfigSchema {
	var returns *NodeConfigSchema
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeConfigFile) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeConfigFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeConfigFile) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodeConfigFile(project projen.Project, options *NodeConfigFileOptions) NodeConfigFile {
	_init_.Initialize()

	if err := validateNewNodeConfigFileParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodeConfigFile{}

	_jsii_.Create(
		"projen.javascript.NodeConfigFile",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNodeConfigFile_Override(n NodeConfigFile, project projen.Project, options *NodeConfigFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NodeConfigFile",
		[]interface{}{project, options},
		n,
	)
}

// Test whether the given construct is a component.
// Experimental.
func NodeConfigFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodeConfigFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NodeConfigFile",
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
func NodeConfigFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodeConfigFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NodeConfigFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the `NodeConfigFile` instance associated with a project or `undefined` if there is none.
// Experimental.
func NodeConfigFile_Of(project projen.Project) NodeConfigFile {
	_init_.Initialize()

	if err := validateNodeConfigFile_OfParameters(project); err != nil {
		panic(err)
	}
	var returns NodeConfigFile

	_jsii_.StaticInvoke(
		"projen.javascript.NodeConfigFile",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeConfigFile) PostProjectCreation(initProject *projen.InitProject) {
	if err := n.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (n *jsiiProxy_NodeConfigFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeConfigFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeConfigFile) ProjectCreation(initProject *projen.InitProject) {
	if err := n.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (n *jsiiProxy_NodeConfigFile) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeConfigFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeConfigFile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

