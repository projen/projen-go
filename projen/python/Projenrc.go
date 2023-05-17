package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Allows writing projenrc files in python.
//
// This will install `projen` as a Python dependency and will add a
// `synth` task which will run `.projenrc.py`.
// Experimental.
type Projenrc interface {
	projen.ProjenrcFile
	// The name of the projenrc file.
	// Experimental.
	FilePath() *string
	// Experimental.
	Project() projen.Project
	// Path to the python executable to use.
	// Experimental.
	PythonExec() *string
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

// The jsii proxy struct for Projenrc
type jsiiProxy_Projenrc struct {
	internal.Type__projenProjenrcFile
}

func (j *jsiiProxy_Projenrc) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Projenrc) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Projenrc) PythonExec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonExec",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrc(project projen.Project, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

	if err := validateNewProjenrcParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.python.Projenrc",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrc_Override(p Projenrc, project projen.Project, options *ProjenrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Projenrc",
		[]interface{}{project, options},
		p,
	)
}

// Returns the `Projenrc` instance associated with a project or `undefined` if there is no Projenrc.
//
// Returns: A Projenrc.
// Experimental.
func Projenrc_Of(project projen.Project) projen.ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrc_OfParameters(project); err != nil {
		panic(err)
	}
	var returns projen.ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.python.Projenrc",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

