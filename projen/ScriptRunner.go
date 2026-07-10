package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A script runner that executes the entrypoint file directly.
//
// A runner is a {@link FutureComponent}: it can be created standalone (e.g. in
// `.projenrc.ts`) and is attached to a project by whoever consumes it.
// Experimental.
type ScriptRunner interface {
	FutureComponent
	IScriptRunner
	// Whether `attach()` has been called.
	//
	// A convenience for tests/introspection;
	// prefer `tryAttach()` over reading this and branching.
	// Experimental.
	Attached() *bool
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// Attach the component to a scope. Only now does it become usable.
	//
	// Returns the real, unwrapped component (not the proxy). A component may be
	// attached exactly once; attaching an already-attached component throws (copy
	// it first to attach a variant elsewhere). Use `tryAttach()` if you don't care
	// whether it has already been attached.
	// Experimental.
	Attach(scope constructs.IConstruct, id *string) FutureComponent
	// Produce the configuration to run the given entrypoint.
	// Experimental.
	ConfigFor(entrypoint *string) *RunScriptConfig
	// Project-dependent setup.
	//
	// Runs once, from `attach()`, when `this.project` is
	// finally available.
	// Experimental.
	Init()
	// Called once, right after `postSynthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// It is also skipped when post-synthesis steps are disabled, e.g. `--no-post` or `PROJEN_DISABLE_POST`.
	// Use it for one-off setup that can be turned off by the user, like running a task to give the user immediate
	// feedback on their new project. Order across components is not guaranteed.
	// Experimental.
	PostProjectCreation(initProject *InitProject)
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
	ProjectCreation(initProject *InitProject)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Attach the component if it isn't already, without caring *where*.
	//
	// Unlike `attach()`, never throws on an already-attached component: if attached
	// anywhere at all, the existing instance is returned and `scope` is ignored.
	// Use `attach()` when attaching to a specific scope is part of your contract
	// and a pre-existing attachment elsewhere would be a bug.
	// Experimental.
	TryAttach(scope constructs.IConstruct, id *string) FutureComponent
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

// The jsii proxy struct for ScriptRunner
type jsiiProxy_ScriptRunner struct {
	jsiiProxy_FutureComponent
	jsiiProxy_IScriptRunner
}

func (j *jsiiProxy_ScriptRunner) Attached() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"attached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScriptRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScriptRunner) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewScriptRunner() ScriptRunner {
	_init_.Initialize()

	j := jsiiProxy_ScriptRunner{}

	_jsii_.Create(
		"projen.ScriptRunner",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewScriptRunner_Override(s ScriptRunner) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ScriptRunner",
		nil, // no parameters
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func ScriptRunner_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScriptRunner_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ScriptRunner",
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
func ScriptRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScriptRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ScriptRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScriptRunner) Attach(scope constructs.IConstruct, id *string) FutureComponent {
	if err := s.validateAttachParameters(scope); err != nil {
		panic(err)
	}
	var returns FutureComponent

	_jsii_.Invoke(
		s,
		"attach",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScriptRunner) ConfigFor(entrypoint *string) *RunScriptConfig {
	if err := s.validateConfigForParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns *RunScriptConfig

	_jsii_.Invoke(
		s,
		"configFor",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScriptRunner) Init() {
	_jsii_.InvokeVoid(
		s,
		"init",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScriptRunner) PostProjectCreation(initProject *InitProject) {
	if err := s.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (s *jsiiProxy_ScriptRunner) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScriptRunner) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScriptRunner) ProjectCreation(initProject *InitProject) {
	if err := s.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (s *jsiiProxy_ScriptRunner) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScriptRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScriptRunner) TryAttach(scope constructs.IConstruct, id *string) FutureComponent {
	if err := s.validateTryAttachParameters(scope); err != nil {
		panic(err)
	}
	var returns FutureComponent

	_jsii_.Invoke(
		s,
		"tryAttach",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScriptRunner) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

