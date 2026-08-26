package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Java `.properties` file.
//
// Properties files consist of key=value pairs, one per line,
// with `#` used for comments. This file format is commonly
// used by tools like SonarQube, Gradle, and other JVM-based
// ecosystem tools.
//
// The object passed as `obj` can be a nested structure. Nested keys are
// flattened into dot-separated property keys on synthesis. For example,
// `{ sonar: { sources: "src" } }` produces `sonar.sources=src`.
//
// This means `addOverride` uses dot notation as a path separator,
// consistent with ObjectFile/JsonPatch semantics.
//
// Example:
//   new PropertiesFile(project, 'sonar-project.properties', {
//     obj: {
//       sonar: {
//         projectKey: 'my-project',
//         sources: 'src',
//       },
//     },
//   });
//
//   // Produces:
//   // sonar.projectKey=my-project
//   // sonar.sources=src
//
// Experimental.
type PropertiesFile interface {
	ObjectFile
	// The absolute path of this file.
	// Experimental.
	AbsolutePath() *string
	// Indicates if the file has been changed during synthesis.
	//
	// This property is
	// only available in `postSynthesize()` hooks. If this is `undefined`, the
	// file has not been synthesized yet.
	// Experimental.
	Changed() *bool
	// Indicates if the file will be committed.
	// Experimental.
	Committed() *bool
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable() *bool
	// Experimental.
	SetExecutable(val *bool)
	// The projen marker, used to identify files as projen-generated.
	//
	// Value is undefined if the project is being ejected.
	// Experimental.
	Marker() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty() *bool
	// The file path, relative to the project's outdir.
	// Experimental.
	Path() *string
	// Experimental.
	Project() Project
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly() *bool
	// Experimental.
	SetReadonly(val *bool)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Adds an override to the synthesized object file.
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
	// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
	// ```
	// would add the overrides
	// ```json
	// "compilerOptions": {
	//   "alwaysStrict": true,
	//   "lib": [
	//     "dom",
	//     "dom.iterable",
	//     "esnext"
	//   ]
	//   ...
	// }
	// ...
	// ```.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds to an array in the synthesized object file.
	//
	// If the array is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example, with the following object file
	// ```json
	// "compilerOptions": {
	//   "exclude": ["node_modules"],
	//   "lib": ["es2020"]
	//   ...
	// }
	// ...
	// ```
	//
	// ```typescript
	// project.tsconfig.file.addToArray('compilerOptions.exclude', 'coverage');
	// project.tsconfig.file.addToArray('compilerOptions.lib', 'dom', 'dom.iterable', 'esnext');
	// ```
	// would result in the following object file
	// ```json
	// "compilerOptions": {
	//   "exclude": ["node_modules", "coverage"],
	//   "lib": ["es2020", "dom", "dom.iterable", "esnext"]
	//   ...
	// }
	// ...
	// ```.
	// Experimental.
	AddToArray(path *string, values ...interface{})
	// Returns a unified diff of the old and new file contents with context lines and hunk headers.
	//
	// Only available after synthesis.
	//
	// This is an expensive operation and should only be used on non time-critical
	// code paths, like debug output.
	//
	// Returns: the diff as an array of lines, or `undefined` if the file was
	// not changed or has not been synthesized yet.
	// Default: 3.
	//
	// Experimental.
	Diff(colorize *bool, contextLines *float64) *[]*string
	// Applies an RFC 6902 JSON-patch to the synthesized object file. See https://datatracker.ietf.org/doc/html/rfc6902 for more information.
	//
	// For example, with the following object file
	// ```json
	// "compilerOptions": {
	//   "exclude": ["node_modules"],
	//   "lib": ["es2020"]
	//   ...
	// }
	// ...
	// ```
	//
	// ```typescript
	// project.tsconfig.file.patch(JsonPatch.add("/compilerOptions/exclude/-", "coverage"));
	// project.tsconfig.file.patch(JsonPatch.replace("/compilerOptions/lib", ["dom", "dom.iterable", "esnext"]));
	// ```
	// would result in the following object file
	// ```json
	// "compilerOptions": {
	//   "exclude": ["node_modules", "coverage"],
	//   "lib": ["dom", "dom.iterable", "esnext"]
	//   ...
	// }
	// ...
	// ```.
	// Experimental.
	Patch(patches ...JsonPatch)
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
	// Writes the file to the project's output directory.
	// Experimental.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Experimental.
	SynthesizeContent(resolver IResolver) *string
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

// The jsii proxy struct for PropertiesFile
type jsiiProxy_PropertiesFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_PropertiesFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Committed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"committed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PropertiesFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewPropertiesFile(scope constructs.IConstruct, filePath *string, options *PropertiesFileOptions) PropertiesFile {
	_init_.Initialize()

	if err := validateNewPropertiesFileParameters(scope, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PropertiesFile{}

	_jsii_.Create(
		"projen.PropertiesFile",
		[]interface{}{scope, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPropertiesFile_Override(p PropertiesFile, scope constructs.IConstruct, filePath *string, options *PropertiesFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.PropertiesFile",
		[]interface{}{scope, filePath, options},
		p,
	)
}

func (j *jsiiProxy_PropertiesFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_PropertiesFile)SetReadonly(val *bool) {
	if err := j.validateSetReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

// Test whether the given construct is a component.
// Experimental.
func PropertiesFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePropertiesFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.PropertiesFile",
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
func PropertiesFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePropertiesFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.PropertiesFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PropertiesFile) AddDeletionOverride(path *string) {
	if err := p.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (p *jsiiProxy_PropertiesFile) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PropertiesFile) AddToArray(path *string, values ...interface{}) {
	if err := p.validateAddToArrayParameters(path); err != nil {
		panic(err)
	}
	args := []interface{}{path}
	for _, a := range values {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addToArray",
		args,
	)
}

func (p *jsiiProxy_PropertiesFile) Diff(colorize *bool, contextLines *float64) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"diff",
		[]interface{}{colorize, contextLines},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PropertiesFile) Patch(patches ...JsonPatch) {
	args := []interface{}{}
	for _, a := range patches {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"patch",
		args,
	)
}

func (p *jsiiProxy_PropertiesFile) PostProjectCreation(initProject *InitProject) {
	if err := p.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (p *jsiiProxy_PropertiesFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PropertiesFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PropertiesFile) ProjectCreation(initProject *InitProject) {
	if err := p.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (p *jsiiProxy_PropertiesFile) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PropertiesFile) SynthesizeContent(resolver IResolver) *string {
	if err := p.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PropertiesFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PropertiesFile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

