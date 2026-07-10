package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Automatically merge Pull Requests using Mergify.
//
// > [!NOTE]
// > GitHub now natively provides the same features, so you don't need Mergify
// > anymore. See `GitHubOptions.mergeQueue` and `MergeQueueOptions.autoQueue`.
//
// If `buildJob` is specified, the specified GitHub workflow job ID is required
// to succeed in order for the PR to be merged.
//
// `approvedReviews` specified the number of code review approvals required for
// the PR to be merged.
// See: https://mergify.com/
//
// Experimental.
type AutoMerge interface {
	projen.Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Adds conditions to the auto merge rule.
	// Experimental.
	AddConditions(conditions ...*string)
	// Adds conditions that will be rendered only during synthesis.
	// Experimental.
	AddConditionsLater(later IAddConditionsLater)
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

// The jsii proxy struct for AutoMerge
type jsiiProxy_AutoMerge struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoMerge) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoMerge) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoMerge(github GitHub, options *AutoMergeOptions) AutoMerge {
	_init_.Initialize()

	if err := validateNewAutoMergeParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoMerge{}

	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoMerge_Override(a AutoMerge, github GitHub, options *AutoMergeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{github, options},
		a,
	)
}

// Test whether the given construct is a component.
// Experimental.
func AutoMerge_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoMerge_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.AutoMerge",
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
func AutoMerge_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoMerge_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.AutoMerge",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoMerge) AddConditions(conditions ...*string) {
	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addConditions",
		args,
	)
}

func (a *jsiiProxy_AutoMerge) AddConditionsLater(later IAddConditionsLater) {
	if err := a.validateAddConditionsLaterParameters(later); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addConditionsLater",
		[]interface{}{later},
	)
}

func (a *jsiiProxy_AutoMerge) PostProjectCreation(initProject *projen.InitProject) {
	if err := a.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (a *jsiiProxy_AutoMerge) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) ProjectCreation(initProject *projen.InitProject) {
	if err := a.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (a *jsiiProxy_AutoMerge) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoMerge) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

