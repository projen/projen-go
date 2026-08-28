package polaris

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// A Coverity on Polaris configuration preset for JavaScript/TypeScript projects.
//
// Provides sensible defaults for JavaScript/TypeScript analysis:
// - `capture.languages.include` = `[javascript]`
// - `capture.files.excludeRegex` excludes `node_modules`, `lib`, `dist`,
//   `coverage` and other build artifacts, based on the paths projen's
//   `TypeScriptProject` excludes from git by default
//
// All defaults can be overridden via options. Nested options (e.g.
// `capture`) are deep-merged with the defaults, so overriding one nested
// field does not drop the other defaults in that subtree.
//
// Example:
//   new PolarisJavascriptCoverity(project, {
//     commit: {},
//   });
//
// Experimental.
type PolarisJavascriptCoverity interface {
	PolarisCoverity
	// The YAML file for the Coverity on Polaris configuration.
	// Experimental.
	File() projen.YamlFile
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

// The jsii proxy struct for PolarisJavascriptCoverity
type jsiiProxy_PolarisJavascriptCoverity struct {
	jsiiProxy_PolarisCoverity
}

func (j *jsiiProxy_PolarisJavascriptCoverity) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolarisJavascriptCoverity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolarisJavascriptCoverity) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPolarisJavascriptCoverity(project projen.Project, options *PolarisCoverityJavascriptOptions) PolarisJavascriptCoverity {
	_init_.Initialize()

	if err := validateNewPolarisJavascriptCoverityParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolarisJavascriptCoverity{}

	_jsii_.Create(
		"projen.polaris.PolarisJavascriptCoverity",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPolarisJavascriptCoverity_Override(p PolarisJavascriptCoverity, project projen.Project, options *PolarisCoverityJavascriptOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.polaris.PolarisJavascriptCoverity",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func PolarisJavascriptCoverity_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolarisJavascriptCoverity_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.polaris.PolarisJavascriptCoverity",
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
func PolarisJavascriptCoverity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolarisJavascriptCoverity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.polaris.PolarisJavascriptCoverity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolarisJavascriptCoverity) PostProjectCreation(initProject *projen.InitProject) {
	if err := p.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (p *jsiiProxy_PolarisJavascriptCoverity) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolarisJavascriptCoverity) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolarisJavascriptCoverity) ProjectCreation(initProject *projen.InitProject) {
	if err := p.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (p *jsiiProxy_PolarisJavascriptCoverity) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolarisJavascriptCoverity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolarisJavascriptCoverity) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}

