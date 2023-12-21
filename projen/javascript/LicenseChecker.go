package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Enforces allowed licenses used by dependencies.
// Experimental.
type LicenseChecker interface {
	projen.Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Experimental.
	Task() projen.Task
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

// The jsii proxy struct for LicenseChecker
type jsiiProxy_LicenseChecker struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_LicenseChecker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicenseChecker) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicenseChecker) Task() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}


// Experimental.
func NewLicenseChecker(scope constructs.Construct, options *LicenseCheckerOptions) LicenseChecker {
	_init_.Initialize()

	if err := validateNewLicenseCheckerParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicenseChecker{}

	_jsii_.Create(
		"projen.javascript.LicenseChecker",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLicenseChecker_Override(l LicenseChecker, scope constructs.Construct, options *LicenseCheckerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.LicenseChecker",
		[]interface{}{scope, options},
		l,
	)
}

// Test whether the given construct is a component.
// Experimental.
func LicenseChecker_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLicenseChecker_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.LicenseChecker",
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
func LicenseChecker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLicenseChecker_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.LicenseChecker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicenseChecker) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicenseChecker) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicenseChecker) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicenseChecker) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

