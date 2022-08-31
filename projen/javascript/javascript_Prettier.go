package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Represents prettier configuration.
// Experimental.
type Prettier interface {
	projen.Component
	// The .prettierIgnore file.
	// Experimental.
	IgnoreFile() projen.IgnoreFile
	// Access to the Prettieroverrides to extend those.
	// Experimental.
	Overrides() *[]*PrettierOverride
	// Experimental.
	Project() projen.Project
	// Direct access to the prettier settings.
	// Experimental.
	Settings() *PrettierSettings
	// Defines Prettier ignore Patterns these patterns will be added to the file .prettierignore.
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Add a prettier override.
	// See: https://prettier.io/docs/en/configuration.html#configuration-overrides
	//
	// Experimental.
	AddOverride(override *PrettierOverride)
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

// The jsii proxy struct for Prettier
type jsiiProxy_Prettier struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Prettier) IgnoreFile() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"ignoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Overrides() *[]*PrettierOverride {
	var returns *[]*PrettierOverride
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Settings() *PrettierSettings {
	var returns *PrettierSettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrettier(project NodeProject, options *PrettierOptions) Prettier {
	_init_.Initialize()

	if err := validateNewPrettierParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Prettier{}

	_jsii_.Create(
		"projen.javascript.Prettier",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPrettier_Override(p Prettier, project NodeProject, options *PrettierOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Prettier",
		[]interface{}{project, options},
		p,
	)
}

// Experimental.
func Prettier_Of(project projen.Project) Prettier {
	_init_.Initialize()

	if err := validatePrettier_OfParameters(project); err != nil {
		panic(err)
	}
	var returns Prettier

	_jsii_.StaticInvoke(
		"projen.javascript.Prettier",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prettier) AddIgnorePattern(pattern *string) {
	if err := p.validateAddIgnorePatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (p *jsiiProxy_Prettier) AddOverride(override *PrettierOverride) {
	if err := p.validateAddOverrideParameters(override); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{override},
	)
}

func (p *jsiiProxy_Prettier) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

