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
	InstallTask() projen.Task
	Project() projen.Project
	AddDependency(spec *string)
	AddDevDependency(spec *string)
	InstallDependencies()
	PostSynthesize()
	PreSynthesize()
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

// Adds a runtime dependency.
// Experimental.
func (p *jsiiProxy_Pip) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

// Adds a dev dependency.
// Experimental.
func (p *jsiiProxy_Pip) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

// Installs dependencies (called during post-synthesis).
// Experimental.
func (p *jsiiProxy_Pip) InstallDependencies() {
	_jsii_.InvokeVoid(
		p,
		"installDependencies",
		nil, // no parameters
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Pip) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Pip) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	InstallTask() projen.Task
	Project() projen.Project
	PublishTask() projen.Task
	PublishTestTask() projen.Task
	AddDependency(spec *string)
	AddDevDependency(spec *string)
	InstallDependencies()
	PostSynthesize()
	PreSynthesize()
	SetupEnvironment()
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

// Adds a runtime dependency.
// Experimental.
func (p *jsiiProxy_Poetry) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

// Adds a dev dependency.
// Experimental.
func (p *jsiiProxy_Poetry) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

// Installs dependencies (called during post-synthesis).
// Experimental.
func (p *jsiiProxy_Poetry) InstallDependencies() {
	_jsii_.InvokeVoid(
		p,
		"installDependencies",
		nil, // no parameters
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Poetry) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Poetry) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
// Experimental.
func (p *jsiiProxy_Poetry) SetupEnvironment() {
	_jsii_.InvokeVoid(
		p,
		"setupEnvironment",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	File() projen.TomlFile
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_PoetryPyproject) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_PoetryPyproject) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	// Must be in the form "name <email>"
	// Experimental.
	Authors *[]*string `json:"authors"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers"`
	// A short description of the package (required).
	// Experimental.
	Description *string `json:"description"`
	// A URL to the documentation of the project.
	// Experimental.
	Documentation *string `json:"documentation"`
	// A list of patterns that will be excluded in the final package.
	//
	// If a VCS is being used for a package, the exclude field will be seeded with
	// the VCS’ ignore settings (.gitignore for git for example).
	// Experimental.
	Exclude *[]*string `json:"exclude"`
	// Package extras.
	// Experimental.
	Extras *map[string]*[]*string `json:"extras"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage"`
	// A list of patterns that will be included in the final package.
	// Experimental.
	Include *[]*string `json:"include"`
	// A list of keywords (max: 5) that the package is related to.
	// Experimental.
	Keywords *[]*string `json:"keywords"`
	// License of this package as an SPDX identifier.
	//
	// If the project is proprietary and does not use a specific license, you
	// can set this value as "Proprietary".
	// Experimental.
	License *string `json:"license"`
	// the maintainers of the package.
	//
	// Must be in the form "name <email>"
	// Experimental.
	Maintainers *[]*string `json:"maintainers"`
	// Name of the package (required).
	// Experimental.
	Name *string `json:"name"`
	// A list of packages and modules to include in the final distribution.
	// Experimental.
	Packages *[]interface{} `json:"packages"`
	// Plugins.
	//
	// Must be specified as a table.
	// See: https://toml.io/en/v1.0.0#table
	//
	// Experimental.
	Plugins interface{} `json:"plugins"`
	// The name of the readme file of the package.
	// Experimental.
	Readme *string `json:"readme"`
	// A URL to the repository of the project.
	// Experimental.
	Repository *string `json:"repository"`
	// The scripts or executables that will be installed when installing the package.
	// Experimental.
	Scripts *map[string]interface{} `json:"scripts"`
	// Source registries from which packages are retrieved.
	// Experimental.
	Source *[]interface{} `json:"source"`
	// Project custom URLs, in addition to homepage, repository and documentation.
	//
	// E.g. "Bug Tracker"
	// Experimental.
	Urls *map[string]*string `json:"urls"`
	// Version of the package (required).
	// Experimental.
	Version *string `json:"version"`
	// A list of dependencies for the project.
	//
	// The python version for which your package is compatible is also required.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Dependencies *map[string]interface{} `json:"dependencies"`
	// A list of development dependencies for the project.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDependencies *map[string]interface{} `json:"devDependencies"`
}

// Poetry-specific options.
// See: https://python-poetry.org/docs/pyproject/
//
// Experimental.
type PoetryPyprojectOptionsWithoutDeps struct {
	// The authors of the package.
	//
	// Must be in the form "name <email>"
	// Experimental.
	Authors *[]*string `json:"authors"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers"`
	// A short description of the package (required).
	// Experimental.
	Description *string `json:"description"`
	// A URL to the documentation of the project.
	// Experimental.
	Documentation *string `json:"documentation"`
	// A list of patterns that will be excluded in the final package.
	//
	// If a VCS is being used for a package, the exclude field will be seeded with
	// the VCS’ ignore settings (.gitignore for git for example).
	// Experimental.
	Exclude *[]*string `json:"exclude"`
	// Package extras.
	// Experimental.
	Extras *map[string]*[]*string `json:"extras"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage"`
	// A list of patterns that will be included in the final package.
	// Experimental.
	Include *[]*string `json:"include"`
	// A list of keywords (max: 5) that the package is related to.
	// Experimental.
	Keywords *[]*string `json:"keywords"`
	// License of this package as an SPDX identifier.
	//
	// If the project is proprietary and does not use a specific license, you
	// can set this value as "Proprietary".
	// Experimental.
	License *string `json:"license"`
	// the maintainers of the package.
	//
	// Must be in the form "name <email>"
	// Experimental.
	Maintainers *[]*string `json:"maintainers"`
	// Name of the package (required).
	// Experimental.
	Name *string `json:"name"`
	// A list of packages and modules to include in the final distribution.
	// Experimental.
	Packages *[]interface{} `json:"packages"`
	// Plugins.
	//
	// Must be specified as a table.
	// See: https://toml.io/en/v1.0.0#table
	//
	// Experimental.
	Plugins interface{} `json:"plugins"`
	// The name of the readme file of the package.
	// Experimental.
	Readme *string `json:"readme"`
	// A URL to the repository of the project.
	// Experimental.
	Repository *string `json:"repository"`
	// The scripts or executables that will be installed when installing the package.
	// Experimental.
	Scripts *map[string]interface{} `json:"scripts"`
	// Source registries from which packages are retrieved.
	// Experimental.
	Source *[]interface{} `json:"source"`
	// Project custom URLs, in addition to homepage, repository and documentation.
	//
	// E.g. "Bug Tracker"
	// Experimental.
	Urls *map[string]*string `json:"urls"`
	// Version of the package (required).
	// Experimental.
	Version *string `json:"version"`
}

// Allows writing projenrc files in python.
//
// This will install `projen` as a Python dependency and will add a
// `synth` task which will run `.projenrc.py`.
// Experimental.
type Projenrc interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Filename *string `json:"filename"`
	// The projen version to use.
	// Experimental.
	ProjenVersion *string `json:"projenVersion"`
}

// Experimental.
type Pytest interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Pytest) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Pytest) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	MaxFailures *float64 `json:"maxFailures"`
	// Directory with tests.
	// Experimental.
	Testdir *string `json:"testdir"`
	// Pytest version.
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type PythonPackagingOptions struct {
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName"`
	// Version of the package.
	// Experimental.
	Version *string `json:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `json:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `json:"license"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *PoetryPyprojectOptionsWithoutDeps `json:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `json:"setupConfig"`
}

// Python project.
// Experimental.
type PythonProject interface {
	github.GitHubProject
	AutoApprove() github.AutoApprove
	BuildTask() projen.Task
	CompileTask() projen.Task
	Components() *[]projen.Component
	DefaultTask() projen.Task
	Deps() projen.Dependencies
	DepsManager() IPythonDeps
	DevContainer() vscode.DevContainer
	EnvManager() IPythonEnv
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	Logger() projen.Logger
	ModuleName() *string
	Name() *string
	Outdir() *string
	PackageTask() projen.Task
	PackagingManager() IPythonPackaging
	Parent() projen.Project
	PostCompileTask() projen.Task
	PreCompileTask() projen.Task
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Pytest() Pytest
	Root() projen.Project
	Tasks() projen.Tasks
	TestTask() projen.Task
	Version() *string
	Vscode() vscode.VsCode
	AddDependency(spec *string)
	AddDevDependency(spec *string)
	AddExcludeFromCleanup(globs ...*string)
	AddGitIgnore(pattern *string)
	AddPackageIgnore(_pattern *string)
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	AddTip(message *string)
	AnnotateGenerated(glob *string)
	PostSynthesize()
	PreSynthesize()
	RemoveTask(name *string) projen.Task
	RunTaskCommand(task projen.Task) *string
	Synth()
	TryFindFile(filePath *string) projen.FileBase
	TryFindJsonFile(filePath *string) projen.JsonFile
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

// Adds a runtime dependency.
// Experimental.
func (p *jsiiProxy_PythonProject) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

// Adds a dev dependency.
// Experimental.
func (p *jsiiProxy_PythonProject) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
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

// Adds a .gitignore pattern.
// Experimental.
func (p *jsiiProxy_PythonProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (p *jsiiProxy_PythonProject) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
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

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (p *jsiiProxy_PythonProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		p,
		"addTip",
		[]interface{}{message},
	)
}

// Marks the provided file(s) as being generated.
//
// This is achieved using the
// github-linguist attributes. Generated files do not count against the
// repository statistics and language breakdown.
// See: https://github.com/github/linguist/blob/master/docs/overrides.md
//
// Deprecated.
func (p *jsiiProxy_PythonProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		p,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_PythonProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (p *jsiiProxy_PythonProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
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

// Returns the shell command to execute in order to run a task.
//
// By default, this is `npx projen@<version> <task>`
// Experimental.
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

// Synthesize all project files into `outdir`.
//
// 1. Call "this.preSynthesize()"
// 2. Delete all generated files
// 3. Synthesize all sub-projects
// 4. Synthesize all components of this project
// 5. Call "postSynthesize()" for all components of this project
// 6. Call "this.postSynthesize()"
// Experimental.
func (p *jsiiProxy_PythonProject) Synth() {
	_jsii_.InvokeVoid(
		p,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
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

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
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

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
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
	Name *string `json:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level
	ProjectType projen.ProjectType `json:"projectType"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Experimental.
	ProjenTokenSecret *string `json:"projenTokenSecret"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName"`
	// Version of the package.
	// Experimental.
	Version *string `json:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `json:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `json:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `json:"license"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *PoetryPyprojectOptionsWithoutDeps `json:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `json:"setupConfig"`
	// Name of the python package as used in imports and filenames.
	//
	// Must only consist of alphanumeric characters and underscores.
	// Experimental.
	ModuleName *string `json:"moduleName"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `json:"deps"`
	// List of dev dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDevDependency()`.
	// Experimental.
	DevDeps *[]*string `json:"devDeps"`
	// Use pip with a requirements.txt file to track project dependencies.
	// Experimental.
	Pip *bool `json:"pip"`
	// Use poetry to manage your project dependencies, virtual environment, and (optional) packaging/publishing.
	// Experimental.
	Poetry *bool `json:"poetry"`
	// Use projenrc in python.
	//
	// This will install `projen` as a python dependency and will add a `synth`
	// task which will run `.projenrc.py`.
	// Experimental.
	ProjenrcPython *bool `json:"projenrcPython"`
	// Options related to projenrc in python.
	// Experimental.
	ProjenrcPythonOptions *ProjenrcOptions `json:"projenrcPythonOptions"`
	// Include pytest tests.
	// Experimental.
	Pytest *bool `json:"pytest"`
	// pytest options.
	// Experimental.
	PytestOptions *PytestOptions `json:"pytestOptions"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample *bool `json:"sample"`
	// Use setuptools with a setup.py script for packaging and publishing.
	// Experimental.
	Setuptools *bool `json:"setuptools"`
	// Use venv to manage a virtual environment for installing dependencies inside.
	// Experimental.
	Venv *bool `json:"venv"`
	// Venv options.
	// Experimental.
	VenvOptions *VenvOptions `json:"venvOptions"`
}

// Python code sample.
// Experimental.
type PythonSample interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_PythonSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_PythonSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() projen.Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddPackages(packages ...*string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
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

func RequirementsFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.python.RequirementsFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Adds the specified packages provided in semver format.
//
// Comment lines (start with `#`) are ignored.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (r *jsiiProxy_RequirementsFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (r *jsiiProxy_RequirementsFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (r *jsiiProxy_RequirementsFile) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
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
	PackageProvider IPackageProvider `json:"packageProvider"`
}

// Python packaging script where package metadata can be placed.
// Experimental.
type SetupPy interface {
	projen.FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() projen.Project
	Readonly() *bool
	SetReadonly(val *bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
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

func SetupPy_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.python.SetupPy",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_SetupPy) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_SetupPy) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (s *jsiiProxy_SetupPy) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
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
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `json:"classifiers"`
	// A short project description.
	// Experimental.
	Description *string `json:"description"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage"`
	// The project license.
	// Experimental.
	License *string `json:"license"`
	// Name of the package.
	// Experimental.
	Name *string `json:"name"`
	// List of submodules to be packaged.
	// Experimental.
	Packages *[]*string `json:"packages"`
	// Manually specify package version.
	// Experimental.
	Version *string `json:"version"`
}

// Manages packaging through setuptools with a setup.py script.
// Experimental.
type Setuptools interface {
	projen.Component
	IPythonPackaging
	Project() projen.Project
	PublishTask() projen.Task
	PublishTestTask() projen.Task
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_Setuptools) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_Setuptools) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	SetupEnvironment()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (v *jsiiProxy_Venv) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (v *jsiiProxy_Venv) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
// Experimental.
func (v *jsiiProxy_Venv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		v,
		"setupEnvironment",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Envdir *string `json:"envdir"`
}

