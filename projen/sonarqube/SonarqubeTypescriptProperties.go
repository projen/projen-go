package sonarqube

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// A SonarQube configuration preset for TypeScript projects.
//
// Provides sensible defaults for TypeScript analysis:
// - `sonar.language` = `ts`
// - `sonar.sources` = `src`
// - `sonar.tests` = `test`
// - `sonar.sourceEncoding` = `UTF-8`
// - `sonar.profile` = `Sonar Way`
// - `sonar.scm.provider` = `git`
// - `sonar.typescript.tsconfigPath` = `tsconfig.json`
// - Typical exclusions for `node_modules`, `coverage`, test files
// - `sonar.javascript.lcov.reportPaths` = `coverage/lcov.info`
//
// All defaults can be overridden via options. Nested options (e.g. `coverage`,
// `javascript`) are deep-merged with the defaults, so overriding one nested
// field does not drop the other defaults in that subtree.
//
// Example:
//   new SonarqubeTypescriptProperties(project, {
//     projectKey: 'my-org_my-ts-project',
//   });
//
// Experimental.
type SonarqubeTypescriptProperties interface {
	SonarqubeProperties
	// The underlying properties file.
	// Experimental.
	File() projen.PropertiesFile
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

// The jsii proxy struct for SonarqubeTypescriptProperties
type jsiiProxy_SonarqubeTypescriptProperties struct {
	jsiiProxy_SonarqubeProperties
}

func (j *jsiiProxy_SonarqubeTypescriptProperties) File() projen.PropertiesFile {
	var returns projen.PropertiesFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SonarqubeTypescriptProperties) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SonarqubeTypescriptProperties) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewSonarqubeTypescriptProperties(scope constructs.IConstruct, options *SonarqubeTypescriptPropertiesOptions) SonarqubeTypescriptProperties {
	_init_.Initialize()

	if err := validateNewSonarqubeTypescriptPropertiesParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SonarqubeTypescriptProperties{}

	_jsii_.Create(
		"projen.sonarqube.SonarqubeTypescriptProperties",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSonarqubeTypescriptProperties_Override(s SonarqubeTypescriptProperties, scope constructs.IConstruct, options *SonarqubeTypescriptPropertiesOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.sonarqube.SonarqubeTypescriptProperties",
		[]interface{}{scope, options},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func SonarqubeTypescriptProperties_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSonarqubeTypescriptProperties_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.sonarqube.SonarqubeTypescriptProperties",
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
func SonarqubeTypescriptProperties_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSonarqubeTypescriptProperties_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.sonarqube.SonarqubeTypescriptProperties",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) PostProjectCreation(initProject *projen.InitProject) {
	if err := s.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) ProjectCreation(initProject *projen.InitProject) {
	if err := s.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

