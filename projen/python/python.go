package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/python/internal"
	"github.com/projen/projen-go/projen/vscode"
)

// Experimental.
type IPackageProvider interface {
	// An array of packages (may be dynamically generated).
	// Experimental.
	Packages() *[]*projen.Dependency
}

// The jsii proxy for IPackageProvider
type jsiiProxy_IPackageProvider struct {
	_ byte // padding
}

func (j *jsiiProxy_IPackageProvider) Packages() *[]*projen.Dependency {
	var returns *[]*projen.Dependency
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

// Experimental.
type IPythonDeps interface {
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Installs dependencies (called during post-synthesis).
	// Experimental.
	InstallDependencies()
	// A task that installs and updates dependencies.
	// Experimental.
	InstallTask() projen.Task
}

// The jsii proxy for IPythonDeps
type jsiiProxy_IPythonDeps struct {
	_ byte // padding
}

func (i *jsiiProxy_IPythonDeps) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		i,
		"addDependency",
		[]interface{}{spec},
	)
}

func (i *jsiiProxy_IPythonDeps) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		i,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (i *jsiiProxy_IPythonDeps) InstallDependencies() {
	_jsii_.InvokeVoid(
		i,
		"installDependencies",
		nil, // no parameters
	)
}

func (j *jsiiProxy_IPythonDeps) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

// Experimental.
type IPythonEnv interface {
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
}

// The jsii proxy for IPythonEnv
type jsiiProxy_IPythonEnv struct {
	_ byte // padding
}

func (i *jsiiProxy_IPythonEnv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		i,
		"setupEnvironment",
		nil, // no parameters
	)
}

// Experimental.
type IPythonPackaging interface {
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
}

// The jsii proxy for IPythonPackaging
type jsiiProxy_IPythonPackaging struct {
	_ byte // padding
}

func (j *jsiiProxy_IPythonPackaging) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

// Manages dependencies using a requirements.txt file and the pip CLI tool.
// Experimental.
type Pip interface {
	projen.Component
	IPythonDeps
	// A task that installs and updates dependencies.
	// Experimental.
	InstallTask() projen.Task
	// Experimental.
	Project() projen.Project
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Installs dependencies (called during post-synthesis).
	// Experimental.
	InstallDependencies()
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

// The jsii proxy struct for Pip
type jsiiProxy_Pip struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonDeps
}

func (j *jsiiProxy_Pip) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pip) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPip(project PythonProject, _options *PipOptions) Pip {
	_init_.Initialize()

	j := jsiiProxy_Pip{}

	_jsii_.Create(
		"projen.python.Pip",
		[]interface{}{project, _options},
		&j,
	)

	return &j
}

// Experimental.
func NewPip_Override(p Pip, project PythonProject, _options *PipOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Pip",
		[]interface{}{project, _options},
		p,
	)
}

func (p *jsiiProxy_Pip) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Pip) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Pip) InstallDependencies() {
	_jsii_.InvokeVoid(
		p,
		"installDependencies",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for pip.
// Experimental.
type PipOptions struct {
}

// Manage project dependencies, virtual environments, and packaging through the poetry CLI tool.
// Experimental.
type Poetry interface {
	projen.Component
	IPythonDeps
	IPythonEnv
	IPythonPackaging
	// A task that installs and updates dependencies.
	// Experimental.
	InstallTask() projen.Task
	// Experimental.
	Project() projen.Project
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
	// A task that uploads the package to the Test PyPI repository.
	// Experimental.
	PublishTestTask() projen.Task
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Installs dependencies (called during post-synthesis).
	// Experimental.
	InstallDependencies()
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Poetry
type jsiiProxy_Poetry struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonDeps
	jsiiProxy_IPythonEnv
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Poetry) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poetry) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poetry) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poetry) PublishTestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTestTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewPoetry(project PythonProject, options *PythonPackagingOptions) Poetry {
	_init_.Initialize()

	j := jsiiProxy_Poetry{}

	_jsii_.Create(
		"projen.python.Poetry",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPoetry_Override(p Poetry, project PythonProject, options *PythonPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Poetry",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_Poetry) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Poetry) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Poetry) InstallDependencies() {
	_jsii_.InvokeVoid(
		p,
		"installDependencies",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) SetupEnvironment() {
	_jsii_.InvokeVoid(
		p,
		"setupEnvironment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Represents configuration of a pyproject.toml file for a Poetry project.
// See: https://python-poetry.org/docs/pyproject/
//
// Experimental.
type PoetryPyproject interface {
	projen.Component
	// Experimental.
	File() projen.TomlFile
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for PoetryPyproject
type jsiiProxy_PoetryPyproject struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PoetryPyproject) File() projen.TomlFile {
	var returns projen.TomlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PoetryPyproject) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPoetryPyproject(project PythonProject, options *PoetryPyprojectOptions) PoetryPyproject {
	_init_.Initialize()

	j := jsiiProxy_PoetryPyproject{}

	_jsii_.Create(
		"projen.python.PoetryPyproject",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPoetryPyproject_Override(p PoetryPyproject, project PythonProject, options *PoetryPyprojectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PoetryPyproject",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_PoetryPyproject) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PoetryPyproject) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PoetryPyproject) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Poetry-specific options.
// See: https://python-poetry.org/docs/pyproject/
//
// Experimental.
type PoetryPyprojectOptions struct {
	// The authors of the package.
	//
	// Must be in the form "name <email>".
	// Experimental.
	Authors *[]*string `json:"authors" yaml:"authors"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers" yaml:"classifiers"`
	// A short description of the package (required).
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// A URL to the documentation of the project.
	// Experimental.
	Documentation *string `json:"documentation" yaml:"documentation"`
	// A list of patterns that will be excluded in the final package.
	//
	// If a VCS is being used for a package, the exclude field will be seeded with
	// the VCS’ ignore settings (.gitignore for git for example).
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// Package extras.
	// Experimental.
	Extras *map[string]*[]*string `json:"extras" yaml:"extras"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// A list of patterns that will be included in the final package.
	// Experimental.
	Include *[]*string `json:"include" yaml:"include"`
	// A list of keywords (max: 5) that the package is related to.
	// Experimental.
	Keywords *[]*string `json:"keywords" yaml:"keywords"`
	// License of this package as an SPDX identifier.
	//
	// If the project is proprietary and does not use a specific license, you
	// can set this value as "Proprietary".
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// the maintainers of the package.
	//
	// Must be in the form "name <email>".
	// Experimental.
	Maintainers *[]*string `json:"maintainers" yaml:"maintainers"`
	// Name of the package (required).
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// A list of packages and modules to include in the final distribution.
	// Experimental.
	Packages *[]interface{} `json:"packages" yaml:"packages"`
	// Plugins.
	//
	// Must be specified as a table.
	// See: https://toml.io/en/v1.0.0#table
	//
	// Experimental.
	Plugins interface{} `json:"plugins" yaml:"plugins"`
	// The name of the readme file of the package.
	// Experimental.
	Readme *string `json:"readme" yaml:"readme"`
	// A URL to the repository of the project.
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// The scripts or executables that will be installed when installing the package.
	// Experimental.
	Scripts *map[string]interface{} `json:"scripts" yaml:"scripts"`
	// Source registries from which packages are retrieved.
	// Experimental.
	Source *[]interface{} `json:"source" yaml:"source"`
	// Project custom URLs, in addition to homepage, repository and documentation.
	//
	// E.g. "Bug Tracker"
	// Experimental.
	Urls *map[string]*string `json:"urls" yaml:"urls"`
	// Version of the package (required).
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// A list of dependencies for the project.
	//
	// The python version for which your package is compatible is also required.
	//
	// Example:
	//   { requests: "^2.13.0" }
	//
	// Experimental.
	Dependencies *map[string]interface{} `json:"dependencies" yaml:"dependencies"`
	// A list of development dependencies for the project.
	//
	// Example:
	//   { requests: "^2.13.0" }
	//
	// Experimental.
	DevDependencies *map[string]interface{} `json:"devDependencies" yaml:"devDependencies"`
}

// Poetry-specific options.
// See: https://python-poetry.org/docs/pyproject/
//
// Experimental.
type PoetryPyprojectOptionsWithoutDeps struct {
	// The authors of the package.
	//
	// Must be in the form "name <email>".
	// Experimental.
	Authors *[]*string `json:"authors" yaml:"authors"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers" yaml:"classifiers"`
	// A short description of the package (required).
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// A URL to the documentation of the project.
	// Experimental.
	Documentation *string `json:"documentation" yaml:"documentation"`
	// A list of patterns that will be excluded in the final package.
	//
	// If a VCS is being used for a package, the exclude field will be seeded with
	// the VCS’ ignore settings (.gitignore for git for example).
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// Package extras.
	// Experimental.
	Extras *map[string]*[]*string `json:"extras" yaml:"extras"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// A list of patterns that will be included in the final package.
	// Experimental.
	Include *[]*string `json:"include" yaml:"include"`
	// A list of keywords (max: 5) that the package is related to.
	// Experimental.
	Keywords *[]*string `json:"keywords" yaml:"keywords"`
	// License of this package as an SPDX identifier.
	//
	// If the project is proprietary and does not use a specific license, you
	// can set this value as "Proprietary".
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// the maintainers of the package.
	//
	// Must be in the form "name <email>".
	// Experimental.
	Maintainers *[]*string `json:"maintainers" yaml:"maintainers"`
	// Name of the package (required).
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// A list of packages and modules to include in the final distribution.
	// Experimental.
	Packages *[]interface{} `json:"packages" yaml:"packages"`
	// Plugins.
	//
	// Must be specified as a table.
	// See: https://toml.io/en/v1.0.0#table
	//
	// Experimental.
	Plugins interface{} `json:"plugins" yaml:"plugins"`
	// The name of the readme file of the package.
	// Experimental.
	Readme *string `json:"readme" yaml:"readme"`
	// A URL to the repository of the project.
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// The scripts or executables that will be installed when installing the package.
	// Experimental.
	Scripts *map[string]interface{} `json:"scripts" yaml:"scripts"`
	// Source registries from which packages are retrieved.
	// Experimental.
	Source *[]interface{} `json:"source" yaml:"source"`
	// Project custom URLs, in addition to homepage, repository and documentation.
	//
	// E.g. "Bug Tracker"
	// Experimental.
	Urls *map[string]*string `json:"urls" yaml:"urls"`
	// Version of the package (required).
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// Allows writing projenrc files in python.
//
// This will install `projen` as a Python dependency and will add a
// `synth` task which will run `.projenrc.py`.
// Experimental.
type Projenrc interface {
	projen.Component
	// Experimental.
	Project() projen.Project
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
	internal.Type__projenComponent
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


// Experimental.
func NewProjenrc(project projen.Project, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

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

// Options for `Projenrc`.
// Experimental.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Experimental.
	Filename *string `json:"filename" yaml:"filename"`
	// The projen version to use.
	// Experimental.
	ProjenVersion *string `json:"projenVersion" yaml:"projenVersion"`
}

// Experimental.
type Pytest interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// Experimental.
	Testdir() *string
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

// The jsii proxy struct for Pytest
type jsiiProxy_Pytest struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Pytest) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pytest) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}


// Experimental.
func NewPytest(project PythonProject, options *PytestOptions) Pytest {
	_init_.Initialize()

	j := jsiiProxy_Pytest{}

	_jsii_.Create(
		"projen.python.Pytest",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPytest_Override(p Pytest, project PythonProject, options *PytestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Pytest",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_Pytest) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pytest) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pytest) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type PytestOptions struct {
	// Stop the testing process after the first N failures.
	// Experimental.
	MaxFailures *float64 `json:"maxFailures" yaml:"maxFailures"`
	// Directory with tests.
	// Experimental.
	Testdir *string `json:"testdir" yaml:"testdir"`
	// Pytest version.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// Experimental.
type PytestSample interface {
	projen.Component
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for PytestSample
type jsiiProxy_PytestSample struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PytestSample) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPytestSample(project PythonProject, testdir *string) PytestSample {
	_init_.Initialize()

	j := jsiiProxy_PytestSample{}

	_jsii_.Create(
		"projen.python.PytestSample",
		[]interface{}{project, testdir},
		&j,
	)

	return &j
}

// Experimental.
func NewPytestSample_Override(p PytestSample, project PythonProject, testdir *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PytestSample",
		[]interface{}{project, testdir},
		p,
	)
}

func (p *jsiiProxy_PytestSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PytestSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PytestSample) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type PythonPackagingOptions struct {
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName" yaml:"authorName"`
	// Version of the package.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers" yaml:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *PoetryPyprojectOptionsWithoutDeps `json:"poetryOptions" yaml:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `json:"setupConfig" yaml:"setupConfig"`
}

// Python project.
// Experimental.
type PythonProject interface {
	github.GitHubProject
	// Auto approve set up for this project.
	// Deprecated.
	AutoApprove() github.AutoApprove
	// Experimental.
	BuildTask() projen.Task
	// Experimental.
	CompileTask() projen.Task
	// Returns all the components within this project.
	// Experimental.
	Components() *[]projen.Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Experimental.
	DefaultTask() projen.Task
	// Project dependencies.
	// Experimental.
	Deps() projen.Dependencies
	// API for managing dependencies.
	// Experimental.
	DepsManager() IPythonDeps
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated.
	DevContainer() vscode.DevContainer
	// Whether or not the project is being ejected.
	// Experimental.
	Ejected() *bool
	// API for mangaging the Python runtime environment.
	// Experimental.
	EnvManager() IPythonEnv
	// All files in this project.
	// Experimental.
	Files() *[]projen.FileBase
	// The .gitattributes file for this repository.
	// Experimental.
	Gitattributes() projen.GitAttributesFile
	// Access all github components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Github() github.GitHub
	// .gitignore.
	// Experimental.
	Gitignore() projen.IgnoreFile
	// Access for Gitpod.
	//
	// This will be `undefined` if gitpod boolean is false.
	// Deprecated.
	Gitpod() projen.Gitpod
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Experimental.
	InitProject() *projen.InitProject
	// Logging utilities.
	// Experimental.
	Logger() projen.Logger
	// Python module name (the project name, with any hyphens or periods replaced with underscores).
	// Experimental.
	ModuleName() *string
	// Project name.
	// Experimental.
	Name() *string
	// Absolute output directory of this project.
	// Experimental.
	Outdir() *string
	// Experimental.
	PackageTask() projen.Task
	// API for managing packaging the project as a library.
	//
	// Only applies when the `projectType` is LIB.
	// Experimental.
	PackagingManager() IPythonPackaging
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() projen.Project
	// Experimental.
	PostCompileTask() projen.Task
	// Experimental.
	PreCompileTask() projen.Task
	// Manages the build process of the project.
	// Experimental.
	ProjectBuild() projen.ProjectBuild
	// Deprecated.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand() *string
	// Pytest component.
	// Experimental.
	Pytest() Pytest
	// Experimental.
	SetPytest(val Pytest)
	// The root project.
	// Experimental.
	Root() projen.Project
	// Project tasks.
	// Experimental.
	Tasks() projen.Tasks
	// Experimental.
	TestTask() projen.Task
	// Version of the package for distribution (should follow semver).
	// Experimental.
	Version() *string
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Vscode() vscode.VsCode
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Experimental.
	AddExcludeFromCleanup(globs ...*string)
	// Adds a .gitignore pattern.
	// Experimental.
	AddGitIgnore(pattern *string)
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Experimental.
	AddPackageIgnore(_pattern *string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Experimental.
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	// Prints a "tip" message during synthesis.
	// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
	AddTip(message *string)
	// Marks the provided file(s) as being generated.
	//
	// This is achieved using the
	// github-linguist attributes. Generated files do not count against the
	// repository statistics and language breakdown.
	// See: https://github.com/github/linguist/blob/master/docs/overrides.md
	//
	// Deprecated.
	AnnotateGenerated(glob *string)
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before all components are synthesized.
	// Experimental.
	PreSynthesize()
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Experimental.
	RemoveTask(name *string) projen.Task
	// Returns the shell command to execute in order to run a task.
	//
	// By default, this is `npx projen@<version> <task>`.
	// Experimental.
	RunTaskCommand(task projen.Task) *string
	// Synthesize all project files into `outdir`.
	//
	// 1. Call "this.preSynthesize()"
	// 2. Delete all generated files
	// 3. Synthesize all sub-projects
	// 4. Synthesize all components of this project
	// 5. Call "postSynthesize()" for all components of this project
	// 6. Call "this.postSynthesize()"
	// Experimental.
	Synth()
	// Finds a file at the specified relative path within this project and all its subprojects.
	//
	// Returns: a `FileBase` or undefined if there is no file in that path.
	// Experimental.
	TryFindFile(filePath *string) projen.FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) projen.JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Experimental.
	TryFindObjectFile(filePath *string) projen.ObjectFile
}

// The jsii proxy struct for PythonProject
type jsiiProxy_PythonProject struct {
	internal.Type__githubGitHubProject
}

func (j *jsiiProxy_PythonProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) DepsManager() IPythonDeps {
	var returns IPythonDeps
	_jsii_.Get(
		j,
		"depsManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) EnvManager() IPythonEnv {
	var returns IPythonEnv
	_jsii_.Get(
		j,
		"envManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) ModuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) PackagingManager() IPythonPackaging {
	var returns IPythonPackaging
	_jsii_.Get(
		j,
		"packagingManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Pytest() Pytest {
	var returns Pytest
	_jsii_.Get(
		j,
		"pytest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonProject(options *PythonProjectOptions) PythonProject {
	_init_.Initialize()

	j := jsiiProxy_PythonProject{}

	_jsii_.Create(
		"projen.python.PythonProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonProject_Override(p PythonProject, options *PythonProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PythonProject",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PythonProject) SetPytest(val Pytest) {
	_jsii_.Set(
		j,
		"pytest",
		val,
	)
}

func PythonProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.python.PythonProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PythonProject) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_PythonProject) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_PythonProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addExcludeFromCleanup",
		args,
	)
}

func (p *jsiiProxy_PythonProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (p *jsiiProxy_PythonProject) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (p *jsiiProxy_PythonProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		p,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		p,
		"addTip",
		[]interface{}{message},
	)
}

func (p *jsiiProxy_PythonProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		p,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (p *jsiiProxy_PythonProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		p,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonProject) Synth() {
	_jsii_.InvokeVoid(
		p,
		"synth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		p,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		p,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		p,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Options for `PythonProject`.
// Experimental.
type PythonProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials github.GithubCredentials `json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode" yaml:"vscode"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName" yaml:"authorName"`
	// Version of the package.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers" yaml:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *PoetryPyprojectOptionsWithoutDeps `json:"poetryOptions" yaml:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `json:"setupConfig" yaml:"setupConfig"`
	// Name of the python package as used in imports and filenames.
	//
	// Must only consist of alphanumeric characters and underscores.
	// Experimental.
	ModuleName *string `json:"moduleName" yaml:"moduleName"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `json:"deps" yaml:"deps"`
	// List of dev dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDevDependency()`.
	// Experimental.
	DevDeps *[]*string `json:"devDeps" yaml:"devDeps"`
	// Use pip with a requirements.txt file to track project dependencies.
	// Experimental.
	Pip *bool `json:"pip" yaml:"pip"`
	// Use poetry to manage your project dependencies, virtual environment, and (optional) packaging/publishing.
	// Experimental.
	Poetry *bool `json:"poetry" yaml:"poetry"`
	// Use projenrc in python.
	//
	// This will install `projen` as a python dependency and will add a `synth`
	// task which will run `.projenrc.py`.
	// Experimental.
	ProjenrcPython *bool `json:"projenrcPython" yaml:"projenrcPython"`
	// Options related to projenrc in python.
	// Experimental.
	ProjenrcPythonOptions *ProjenrcOptions `json:"projenrcPythonOptions" yaml:"projenrcPythonOptions"`
	// Include pytest tests.
	// Experimental.
	Pytest *bool `json:"pytest" yaml:"pytest"`
	// pytest options.
	// Experimental.
	PytestOptions *PytestOptions `json:"pytestOptions" yaml:"pytestOptions"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample *bool `json:"sample" yaml:"sample"`
	// Use setuptools with a setup.py script for packaging and publishing.
	// Experimental.
	Setuptools *bool `json:"setuptools" yaml:"setuptools"`
	// Use venv to manage a virtual environment for installing dependencies inside.
	// Experimental.
	Venv *bool `json:"venv" yaml:"venv"`
	// Venv options.
	// Experimental.
	VenvOptions *VenvOptions `json:"venvOptions" yaml:"venvOptions"`
}

// Python code sample.
// Experimental.
type PythonSample interface {
	projen.Component
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for PythonSample
type jsiiProxy_PythonSample struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PythonSample) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonSample(project PythonProject, _options *PythonSampleOptions) PythonSample {
	_init_.Initialize()

	j := jsiiProxy_PythonSample{}

	_jsii_.Create(
		"projen.python.PythonSample",
		[]interface{}{project, _options},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonSample_Override(p PythonSample, project PythonProject, _options *PythonSampleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PythonSample",
		[]interface{}{project, _options},
		p,
	)
}

func (p *jsiiProxy_PythonSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonSample) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for python sample code.
// Experimental.
type PythonSampleOptions struct {
}

// Specifies a list of packages to be installed using pip.
// See: https://pip.pypa.io/en/stable/reference/pip_install/#requirements-file-format
//
// Experimental.
type RequirementsFile interface {
	projen.FileBase
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
	// The file path, relative to the project root.
	// Experimental.
	Path() *string
	// Experimental.
	Project() projen.Project
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly() *bool
	// Experimental.
	SetReadonly(val *bool)
	// Adds the specified packages provided in semver format.
	//
	// Comment lines (start with `#`) are ignored.
	// Experimental.
	AddPackages(packages ...*string)
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
	SynthesizeContent(resolver projen.IResolver) *string
}

// The jsii proxy struct for RequirementsFile
type jsiiProxy_RequirementsFile struct {
	internal.Type__projenFileBase
}

func (j *jsiiProxy_RequirementsFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewRequirementsFile(project projen.Project, filePath *string, options *RequirementsFileOptions) RequirementsFile {
	_init_.Initialize()

	j := jsiiProxy_RequirementsFile{}

	_jsii_.Create(
		"projen.python.RequirementsFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRequirementsFile_Override(r RequirementsFile, project projen.Project, filePath *string, options *RequirementsFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.RequirementsFile",
		[]interface{}{project, filePath, options},
		r,
	)
}

func (j *jsiiProxy_RequirementsFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_RequirementsFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (r *jsiiProxy_RequirementsFile) AddPackages(packages ...*string) {
	args := []interface{}{}
	for _, a := range packages {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addPackages",
		args,
	)
}

func (r *jsiiProxy_RequirementsFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequirementsFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequirementsFile) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequirementsFile) SynthesizeContent(resolver projen.IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Experimental.
type RequirementsFileOptions struct {
	// Provide a list of packages that can be dynamically updated.
	// Experimental.
	PackageProvider IPackageProvider `json:"packageProvider" yaml:"packageProvider"`
}

// Python packaging script where package metadata can be placed.
// Experimental.
type SetupPy interface {
	projen.FileBase
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
	// The file path, relative to the project root.
	// Experimental.
	Path() *string
	// Experimental.
	Project() projen.Project
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly() *bool
	// Experimental.
	SetReadonly(val *bool)
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
	SynthesizeContent(resolver projen.IResolver) *string
}

// The jsii proxy struct for SetupPy
type jsiiProxy_SetupPy struct {
	internal.Type__projenFileBase
}

func (j *jsiiProxy_SetupPy) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewSetupPy(project PythonProject, options *SetupPyOptions) SetupPy {
	_init_.Initialize()

	j := jsiiProxy_SetupPy{}

	_jsii_.Create(
		"projen.python.SetupPy",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSetupPy_Override(s SetupPy, project PythonProject, options *SetupPyOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.SetupPy",
		[]interface{}{project, options},
		s,
	)
}

func (j *jsiiProxy_SetupPy) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_SetupPy) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (s *jsiiProxy_SetupPy) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SetupPy) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SetupPy) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SetupPy) SynthesizeContent(resolver projen.IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Fields to pass in the setup() function of setup.py.
// See: https://docs.python.org/3/distutils/setupscript.html
//
// Experimental.
type SetupPyOptions struct {
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName" yaml:"authorName"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers" yaml:"classifiers"`
	// A short project description.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// The project license.
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// Name of the package.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// List of submodules to be packaged.
	// Experimental.
	Packages *[]*string `json:"packages" yaml:"packages"`
	// Manually specify package version.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// Manages packaging through setuptools with a setup.py script.
// Experimental.
type Setuptools interface {
	projen.Component
	IPythonPackaging
	// Experimental.
	Project() projen.Project
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
	// A task that uploads the package to the Test PyPI repository.
	// Experimental.
	PublishTestTask() projen.Task
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

// The jsii proxy struct for Setuptools
type jsiiProxy_Setuptools struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Setuptools) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) PublishTestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTestTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewSetuptools(project PythonProject, options *PythonPackagingOptions) Setuptools {
	_init_.Initialize()

	j := jsiiProxy_Setuptools{}

	_jsii_.Create(
		"projen.python.Setuptools",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSetuptools_Override(s Setuptools, project PythonProject, options *PythonPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Setuptools",
		[]interface{}{project, options},
		s,
	)
}

func (s *jsiiProxy_Setuptools) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// Manages a virtual environment through the Python venv module.
// Experimental.
type Venv interface {
	projen.Component
	IPythonEnv
	// Experimental.
	Project() projen.Project
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Venv
type jsiiProxy_Venv struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonEnv
}

func (j *jsiiProxy_Venv) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVenv(project PythonProject, options *VenvOptions) Venv {
	_init_.Initialize()

	j := jsiiProxy_Venv{}

	_jsii_.Create(
		"projen.python.Venv",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewVenv_Override(v Venv, project PythonProject, options *VenvOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Venv",
		[]interface{}{project, options},
		v,
	)
}

func (v *jsiiProxy_Venv) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		v,
		"setupEnvironment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

// Options for venv.
// Experimental.
type VenvOptions struct {
	// Name of directory to store the environment in.
	// Experimental.
	Envdir *string `json:"envdir" yaml:"envdir"`
}

