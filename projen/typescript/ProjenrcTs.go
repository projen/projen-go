package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/typescript/internal"
)

// A projenrc file written in TypeScript.
//
// This component can be instantiated in any type of project
// and has no expectations around the project's main language.
//
// Requires that `npx` is available.
// Experimental.
type ProjenrcTs interface {
	projen.ProjenrcFile
	// The path of the projenrc file.
	// Experimental.
	FilePath() *string
	// Experimental.
	Project() projen.Project
	// TypeScript configuration file used to compile projen source files.
	// Experimental.
	Tsconfig() javascript.TypescriptConfig
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
}

// The jsii proxy struct for ProjenrcTs
type jsiiProxy_ProjenrcTs struct {
	internal.Type__projenProjenrcFile
}

func (j *jsiiProxy_ProjenrcTs) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcTs) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcTs) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrcTs(project projen.Project, options *ProjenrcTsOptions) ProjenrcTs {
	_init_.Initialize()

	if err := validateNewProjenrcTsParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjenrcTs{}

	_jsii_.Create(
		"projen.typescript.ProjenrcTs",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrcTs_Override(p ProjenrcTs, project projen.Project, options *ProjenrcTsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.typescript.ProjenrcTs",
		[]interface{}{project, options},
		p,
	)
}

// Returns the `Projenrc` instance associated with a project or `undefined` if there is no Projenrc.
//
// Returns: A Projenrc.
// Experimental.
func ProjenrcTs_Of(project projen.Project) projen.ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrcTs_OfParameters(project); err != nil {
		panic(err)
	}
	var returns projen.ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.typescript.ProjenrcTs",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjenrcTs) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcTs) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcTs) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

