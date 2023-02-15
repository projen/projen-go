package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// A Project Object Model or POM is the fundamental unit of work in Maven.
//
// It is
// an XML file that contains information about the project and configuration
// details used by Maven to build the project.
// Experimental.
type Pom interface {
	projen.Component
	// Maven artifact ID.
	// Experimental.
	ArtifactId() *string
	// Project description.
	// Experimental.
	Description() *string
	// The name of the pom file.
	// Experimental.
	FileName() *string
	// Maven group ID.
	// Experimental.
	GroupId() *string
	// Project display name.
	// Experimental.
	Name() *string
	// Maven packaging format.
	// Experimental.
	Packaging() *string
	// Experimental.
	Project() projen.Project
	// Project URL.
	// Experimental.
	Url() *string
	// Project version.
	// Experimental.
	Version() *string
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a build plugin to the pom.
	//
	// The plug in is also added as a BUILD dep to the project.
	// Experimental.
	AddPlugin(spec *string, options *PluginOptions) *projen.Dependency
	// Adds a key/value property to the pom.
	// Experimental.
	AddProperty(key *string, value *string)
	// Adds a repository to the pom.
	// Experimental.
	AddRepository(repository *MavenRepository)
	// Adds a test dependency.
	// Experimental.
	AddTestDependency(spec *string)
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

// The jsii proxy struct for Pom
type jsiiProxy_Pom struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Pom) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Packaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewPom(project projen.Project, options *PomOptions) Pom {
	_init_.Initialize()

	if err := validateNewPomParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pom{}

	_jsii_.Create(
		"projen.java.Pom",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPom_Override(p Pom, project projen.Project, options *PomOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Pom",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_Pom) AddDependency(spec *string) {
	if err := p.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Pom) AddPlugin(spec *string, options *PluginOptions) *projen.Dependency {
	if err := p.validateAddPluginParameters(spec, options); err != nil {
		panic(err)
	}
	var returns *projen.Dependency

	_jsii_.Invoke(
		p,
		"addPlugin",
		[]interface{}{spec, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pom) AddProperty(key *string, value *string) {
	if err := p.validateAddPropertyParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addProperty",
		[]interface{}{key, value},
	)
}

func (p *jsiiProxy_Pom) AddRepository(repository *MavenRepository) {
	if err := p.validateAddRepositoryParameters(repository); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addRepository",
		[]interface{}{repository},
	)
}

func (p *jsiiProxy_Pom) AddTestDependency(spec *string) {
	if err := p.validateAddTestDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addTestDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Pom) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pom) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pom) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

