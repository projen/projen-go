package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Represents a YAML file.
// Experimental.
type YamlFile interface {
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
	// Maximum line width (set to 0 to disable folding).
	// Experimental.
	LineWidth() *float64
	// Experimental.
	SetLineWidth(val *float64)
	// The projen marker, used to identify files as projen-generated.
	//
	// Value is undefined if the project is being ejected.
	// Experimental.
	Marker() *string
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty() *bool
	// The file path, relative to the project root.
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
}

// The jsii proxy struct for YamlFile
type jsiiProxy_YamlFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_YamlFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewYamlFile(project Project, filePath *string, options *YamlFileOptions) YamlFile {
	_init_.Initialize()

	if err := validateNewYamlFileParameters(project, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_YamlFile{}

	_jsii_.Create(
		"projen.YamlFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewYamlFile_Override(y YamlFile, project Project, filePath *string, options *YamlFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.YamlFile",
		[]interface{}{project, filePath, options},
		y,
	)
}

func (j *jsiiProxy_YamlFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_YamlFile)SetLineWidth(val *float64) {
	if err := j.validateSetLineWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

func (j *jsiiProxy_YamlFile)SetReadonly(val *bool) {
	if err := j.validateSetReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (y *jsiiProxy_YamlFile) AddDeletionOverride(path *string) {
	if err := y.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		y,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (y *jsiiProxy_YamlFile) AddOverride(path *string, value interface{}) {
	if err := y.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		y,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (y *jsiiProxy_YamlFile) AddToArray(path *string, values ...interface{}) {
	if err := y.validateAddToArrayParameters(path); err != nil {
		panic(err)
	}
	args := []interface{}{path}
	for _, a := range values {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		y,
		"addToArray",
		args,
	)
}

func (y *jsiiProxy_YamlFile) Patch(patches ...JsonPatch) {
	args := []interface{}{}
	for _, a := range patches {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		y,
		"patch",
		args,
	)
}

func (y *jsiiProxy_YamlFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		y,
		"postSynthesize",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YamlFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		y,
		"preSynthesize",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YamlFile) Synthesize() {
	_jsii_.InvokeVoid(
		y,
		"synthesize",
		nil, // no parameters
	)
}

func (y *jsiiProxy_YamlFile) SynthesizeContent(resolver IResolver) *string {
	if err := y.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		y,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

