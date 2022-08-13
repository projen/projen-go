// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Represents a JSON file.
// Experimental.
type JsonFile interface {
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
	//    "alwaysStrict": true,
	//    "lib": [
	//      "dom",
	//      "dom.iterable",
	//      "esnext"
	//    ]
	//    ...
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
	//    "exclude": ["node_modules"],
	//    "lib": ["es2019"]
	//    ...
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
	//    "exclude": ["node_modules", "coverage"],
	//    "lib": ["es2019", "dom", "dom.iterable", "esnext"]
	//    ...
	// }
	// ...
	// ```.
	// Experimental.
	AddToArray(path *string, values ...interface{})
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

// The jsii proxy struct for JsonFile
type jsiiProxy_JsonFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_JsonFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewJsonFile(project Project, filePath *string, options *JsonFileOptions) JsonFile {
	_init_.Initialize()

	j := jsiiProxy_JsonFile{}

	_jsii_.Create(
		"projen.JsonFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJsonFile_Override(j JsonFile, project Project, filePath *string, options *JsonFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.JsonFile",
		[]interface{}{project, filePath, options},
		j,
	)
}

func (j *jsiiProxy_JsonFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_JsonFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (j *jsiiProxy_JsonFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		j,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (j *jsiiProxy_JsonFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_JsonFile) AddToArray(path *string, values ...interface{}) {
	args := []interface{}{path}
	for _, a := range values {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		j,
		"addToArray",
		args,
	)
}

func (j *jsiiProxy_JsonFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JsonFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JsonFile) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JsonFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

