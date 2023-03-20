// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Sets up a project to use JSON for projenrc.
// Experimental.
type ProjenrcJson interface {
	ProjenrcFile
	// The path of the projenrc file.
	// Experimental.
	FilePath() *string
	// Experimental.
	Project() Project
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

// The jsii proxy struct for ProjenrcJson
type jsiiProxy_ProjenrcJson struct {
	jsiiProxy_ProjenrcFile
}

func (j *jsiiProxy_ProjenrcJson) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcJson) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrcJson(project Project, options *ProjenrcJsonOptions) ProjenrcJson {
	_init_.Initialize()

	if err := validateNewProjenrcJsonParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjenrcJson{}

	_jsii_.Create(
		"projen.ProjenrcJson",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrcJson_Override(p ProjenrcJson, project Project, options *ProjenrcJsonOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ProjenrcJson",
		[]interface{}{project, options},
		p,
	)
}

// Returns the `Projenrc` instance associated with a project or `undefined` if there is no Projenrc.
//
// Returns: A Projenrc.
// Experimental.
func ProjenrcJson_Of(project Project) ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrcJson_OfParameters(project); err != nil {
		panic(err)
	}
	var returns ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.ProjenrcJson",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjenrcJson) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcJson) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcJson) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

