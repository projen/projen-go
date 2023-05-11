package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Deprecated: use `ProjenrcJson`.
type Projenrc interface {
	ProjenrcJson
	// The path of the projenrc file.
	// Deprecated: use `ProjenrcJson`.
	FilePath() *string
	// Deprecated: use `ProjenrcJson`.
	Project() Project
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Deprecated: use `ProjenrcJson`.
	PostSynthesize()
	// Called before synthesis.
	// Deprecated: use `ProjenrcJson`.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Deprecated: use `ProjenrcJson`.
	Synthesize()
}

// The jsii proxy struct for Projenrc
type jsiiProxy_Projenrc struct {
	jsiiProxy_ProjenrcJson
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

func (j *jsiiProxy_Projenrc) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Deprecated: use `ProjenrcJson`.
func NewProjenrc(project Project, options *ProjenrcJsonOptions) Projenrc {
	_init_.Initialize()

	if err := validateNewProjenrcParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.Projenrc",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Deprecated: use `ProjenrcJson`.
func NewProjenrc_Override(p Projenrc, project Project, options *ProjenrcJsonOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Projenrc",
		[]interface{}{project, options},
		p,
	)
}

// Returns the `Projenrc` instance associated with a project or `undefined` if there is no Projenrc.
//
// Returns: A Projenrc.
// Deprecated: use `ProjenrcJson`.
func Projenrc_Of(project Project) ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrc_OfParameters(project); err != nil {
		panic(err)
	}
	var returns ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.Projenrc",
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

