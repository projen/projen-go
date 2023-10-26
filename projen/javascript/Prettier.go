package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Represents prettier configuration.
// Experimental.
type Prettier interface {
	projen.Component
	// The .prettierIgnore file.
	// Experimental.
	IgnoreFile() projen.IgnoreFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns all Prettier overrides.
	// Experimental.
	Overrides() *[]*PrettierOverride
	// Experimental.
	Project() projen.Project
	// Direct access to the prettier settings.
	// Experimental.
	Settings() *PrettierSettings
	// Defines Prettier ignore Patterns these patterns will be added to the file .prettierignore.
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Add a prettier override.
	// See: https://prettier.io/docs/en/configuration.html#configuration-overrides
	//
	// Experimental.
	AddOverride(override *PrettierOverride)
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

// The jsii proxy struct for Prettier
type jsiiProxy_Prettier struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Prettier) IgnoreFile() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"ignoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Overrides() *[]*PrettierOverride {
	var returns *[]*PrettierOverride
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Settings() *PrettierSettings {
	var returns *PrettierSettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrettier(project NodeProject, options *PrettierOptions) Prettier {
	_init_.Initialize()

	if err := validateNewPrettierParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Prettier{}

	_jsii_.Create(
		"projen.javascript.Prettier",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPrettier_Override(p Prettier, project NodeProject, options *PrettierOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Prettier",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Prettier_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrettier_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Prettier",
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
func Prettier_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrettier_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Prettier",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Prettier_Of(project projen.Project) Prettier {
	_init_.Initialize()

	if err := validatePrettier_OfParameters(project); err != nil {
		panic(err)
	}
	var returns Prettier

	_jsii_.StaticInvoke(
		"projen.javascript.Prettier",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prettier) AddIgnorePattern(pattern *string) {
	if err := p.validateAddIgnorePatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (p *jsiiProxy_Prettier) AddOverride(override *PrettierOverride) {
	if err := p.validateAddOverrideParameters(override); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{override},
	)
}

func (p *jsiiProxy_Prettier) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

