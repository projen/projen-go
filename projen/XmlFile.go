package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an XML file.
//
// Objects passed in will be synthesized using the npm "xml" library.
// See: https://www.npmjs.com/package/xml
//
// Experimental.
type XmlFile interface {
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
	//   "lib": ["es2019"]
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
	//   "lib": ["es2019", "dom", "dom.iterable", "esnext"]
	//   ...
	// }
	// ...
	// ```.
	// Experimental.
	AddToArray(path *string, values ...interface{})
	// Applies an RFC 6902 JSON-patch to the synthesized object file. See https://datatracker.ietf.org/doc/html/rfc6902 for more information.
	//
	// For example, with the following object file
	// ```json
	// "compilerOptions": {
	//   "exclude": ["node_modules"],
	//   "lib": ["es2019"]
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
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Writes the file to the project's output directory.
	// Experimental.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Experimental.
	SynthesizeContent(resolver IResolver) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for XmlFile
type jsiiProxy_XmlFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_XmlFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewXmlFile(project Project, filePath *string, options *XmlFileOptions) XmlFile {
	_init_.Initialize()

	if err := validateNewXmlFileParameters(project, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_XmlFile{}

	_jsii_.Create(
		"projen.XmlFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewXmlFile_Override(x XmlFile, project Project, filePath *string, options *XmlFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.XmlFile",
		[]interface{}{project, filePath, options},
		x,
	)
}

func (j *jsiiProxy_XmlFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_XmlFile)SetReadonly(val *bool) {
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
func XmlFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXmlFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.XmlFile",
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
func XmlFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXmlFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.XmlFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XmlFile) AddDeletionOverride(path *string) {
	if err := x.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (x *jsiiProxy_XmlFile) AddOverride(path *string, value interface{}) {
	if err := x.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (x *jsiiProxy_XmlFile) AddToArray(path *string, values ...interface{}) {
	if err := x.validateAddToArrayParameters(path); err != nil {
		panic(err)
	}
	args := []interface{}{path}
	for _, a := range values {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		x,
		"addToArray",
		args,
	)
}

func (x *jsiiProxy_XmlFile) Patch(patches ...JsonPatch) {
	args := []interface{}{}
	for _, a := range patches {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		x,
		"patch",
		args,
	)
}

func (x *jsiiProxy_XmlFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		x,
		"postSynthesize",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XmlFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		x,
		"preSynthesize",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XmlFile) Synthesize() {
	_jsii_.InvokeVoid(
		x,
		"synthesize",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XmlFile) SynthesizeContent(resolver IResolver) *string {
	if err := x.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		x,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XmlFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

