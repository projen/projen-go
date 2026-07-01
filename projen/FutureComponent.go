package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A {@link Component} that is created *detached* from any project and attached to one later via {@link FutureComponent.attach}.
//
// Like a regular component, but constructed without a project. It improves on a
// naive deferred component in three ways:
//
// - Use-before-attach is an error, not a silent footgun. The constructor hands
//   the caller a guard proxy; touching `project`, `node`, `synthesize()` or any
//   subclass feature before `attach()` throws.
// - No global shadow-tree leak. Each instance gets its own throwaway shadow
//   root, so detached components never share an id counter and the root becomes
//   garbage once the component is reparented on attach.
// - `attach()` returns the unwrapped component, so callers can opt out of the
//   proxy entirely.
//
// The constructor takes no arguments (`super()`). A subclass that needs options
// captures them itself, reading the local parameter inside its constructor - NOT
// `this.options`, which the guard blocks until attach.
//
// ```ts
// class Worker extends FutureComponent {
//   private readonly options: WorkerOptions;
//   constructor(options: WorkerOptions = {}) {
//     super();
//     this.options = options; // set: allowed
//   }
//   protected init() {
//     // this.project is available here
//   }
// }
//
// const w = new Worker({ retries: 3 });
// // w.project;                  // throws: not attached yet
// const real = w.attach(project); // reparents, runs init(), returns the bare instance
// ```.
// Experimental.
type FutureComponent interface {
	Component
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
	// Project-dependent setup.
	//
	// Runs once, from `attach()`, when `this.project` is
	// finally available.
	// Experimental.
	Init()
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

// The jsii proxy struct for FutureComponent
type jsiiProxy_FutureComponent struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_FutureComponent) Attached() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"attached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FutureComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FutureComponent) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewFutureComponent_Override(f FutureComponent) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.FutureComponent",
		nil, // no parameters
		f,
	)
}

// Test whether the given construct is a component.
// Experimental.
func FutureComponent_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFutureComponent_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.FutureComponent",
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
func FutureComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFutureComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.FutureComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FutureComponent) Attach(scope constructs.IConstruct, id *string) FutureComponent {
	if err := f.validateAttachParameters(scope); err != nil {
		panic(err)
	}
	var returns FutureComponent

	_jsii_.Invoke(
		f,
		"attach",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FutureComponent) Init() {
	_jsii_.InvokeVoid(
		f,
		"init",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FutureComponent) PostSynthesize() {
	_jsii_.InvokeVoid(
		f,
		"postSynthesize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FutureComponent) PreSynthesize() {
	_jsii_.InvokeVoid(
		f,
		"preSynthesize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FutureComponent) Synthesize() {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FutureComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FutureComponent) TryAttach(scope constructs.IConstruct, id *string) FutureComponent {
	if err := f.validateTryAttachParameters(scope); err != nil {
		panic(err)
	}
	var returns FutureComponent

	_jsii_.Invoke(
		f,
		"tryAttach",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FutureComponent) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		f,
		"with",
		args,
		&returns,
	)

	return returns
}

