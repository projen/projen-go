package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/typescript/internal"
)

// The runner used to execute TypeScript files.
//
// A runner is a {@link FutureComponent }: create it standalone (e.g. via one of
// the static factories) and it is attached to a project by the component that
// consumes it. Use {@link TypeScriptRunner.copy} to derive a new runner with
// adjusted options.
// Experimental.
type TypeScriptRunner interface {
	projen.ScriptRunner
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
	Project() projen.Project
	// Attach the component to a scope. Only now does it become usable.
	//
	// Returns the real, unwrapped component (not the proxy). A component may be
	// attached exactly once; attaching an already-attached component throws (copy
	// it first to attach a variant elsewhere). Use `tryAttach()` if you don't care
	// whether it has already been attached.
	// Experimental.
	Attach(scope constructs.IConstruct, id *string) projen.FutureComponent
	// Produce the {@link RunScriptConfig} to run the given entrypoint with this runner.
	// Experimental.
	ConfigFor(entrypoint *string) *projen.RunScriptConfig
	// Returns a new (detached) runner of the same kind, with `overrides` merged over the current options.
	//
	// Safe to call while detached.
	// Experimental.
	Copy(overrides *TypeScriptRunnerOptions) TypeScriptRunner
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
	// Attach the component if it isn't already, without caring *where*.
	//
	// Unlike `attach()`, never throws on an already-attached component: if attached
	// anywhere at all, the existing instance is returned and `scope` is ignored.
	// Use `attach()` when attaching to a specific scope is part of your contract
	// and a pre-existing attachment elsewhere would be a bug.
	// Experimental.
	TryAttach(scope constructs.IConstruct, id *string) projen.FutureComponent
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

// The jsii proxy struct for TypeScriptRunner
type jsiiProxy_TypeScriptRunner struct {
	internal.Type__projenScriptRunner
}

func (j *jsiiProxy_TypeScriptRunner) Attached() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"attached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptRunner) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Test whether the given construct is a component.
// Experimental.
func TypeScriptRunner_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTypeScriptRunner_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.typescript.TypeScriptRunner",
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
func TypeScriptRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTypeScriptRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.typescript.TypeScriptRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Use the native Node.js TypeScript support.
//
// Requires Node.js 22.18.0 or later.
// The script must use ESM style imports (no directories, include file endings, etc.).
//
// Named `nodejs` (not `node`) because a runner is a construct, and `node` is
// reserved by `constructs.Construct` for the construct's tree node.
// Experimental.
func TypeScriptRunner_Nodejs(options *NodeRunnerOptions) TypeScriptRunner {
	_init_.Initialize()

	if err := validateTypeScriptRunner_NodejsParameters(options); err != nil {
		panic(err)
	}
	var returns TypeScriptRunner

	_jsii_.StaticInvoke(
		"projen.typescript.TypeScriptRunner",
		"nodejs",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use ts-node to execute TypeScript files.
// Experimental.
func TypeScriptRunner_TsNode(options *TsNodeRunnerOptions) TypeScriptRunner {
	_init_.Initialize()

	if err := validateTypeScriptRunner_TsNodeParameters(options); err != nil {
		panic(err)
	}
	var returns TypeScriptRunner

	_jsii_.StaticInvoke(
		"projen.typescript.TypeScriptRunner",
		"tsNode",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use tsx to execute TypeScript files.
//
// tsx is a fast TypeScript runtime that does not perform type-checking.
// You may opt-in to an explicit type-checking step before the script is run.
// Experimental.
func TypeScriptRunner_Tsx(options *TsxRunnerOptions) TypeScriptRunner {
	_init_.Initialize()

	if err := validateTypeScriptRunner_TsxParameters(options); err != nil {
		panic(err)
	}
	var returns TypeScriptRunner

	_jsii_.StaticInvoke(
		"projen.typescript.TypeScriptRunner",
		"tsx",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptRunner) Attach(scope constructs.IConstruct, id *string) projen.FutureComponent {
	if err := t.validateAttachParameters(scope); err != nil {
		panic(err)
	}
	var returns projen.FutureComponent

	_jsii_.Invoke(
		t,
		"attach",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptRunner) ConfigFor(entrypoint *string) *projen.RunScriptConfig {
	if err := t.validateConfigForParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns *projen.RunScriptConfig

	_jsii_.Invoke(
		t,
		"configFor",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptRunner) Copy(overrides *TypeScriptRunnerOptions) TypeScriptRunner {
	if err := t.validateCopyParameters(overrides); err != nil {
		panic(err)
	}
	var returns TypeScriptRunner

	_jsii_.Invoke(
		t,
		"copy",
		[]interface{}{overrides},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptRunner) Init() {
	_jsii_.InvokeVoid(
		t,
		"init",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptRunner) PostProjectCreation(initProject *projen.InitProject) {
	if err := t.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (t *jsiiProxy_TypeScriptRunner) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptRunner) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptRunner) ProjectCreation(initProject *projen.InitProject) {
	if err := t.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (t *jsiiProxy_TypeScriptRunner) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptRunner) TryAttach(scope constructs.IConstruct, id *string) projen.FutureComponent {
	if err := t.validateTryAttachParameters(scope); err != nil {
		panic(err)
	}
	var returns projen.FutureComponent

	_jsii_.Invoke(
		t,
		"tryAttach",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptRunner) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

