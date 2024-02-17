package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Represents the npm `package.json` file.
// Experimental.
type NodePackage interface {
	projen.Component
	// Allow project to take library dependencies.
	// Experimental.
	AllowLibraryDependencies() *bool
	// Options for npm packages using AWS CodeArtifact.
	//
	// This is required if publishing packages to, or installing scoped packages from AWS CodeArtifact.
	// Default: - undefined.
	//
	// Experimental.
	CodeArtifactOptions() *CodeArtifactOptions
	// The module's entrypoint (e.g. `lib/index.js`).
	// Experimental.
	Entrypoint() *string
	// The package.json file.
	// Experimental.
	File() projen.JsonFile
	// Renders `yarn install` or `npm install` with lockfile update (not frozen).
	// Experimental.
	InstallAndUpdateLockfileCommand() *string
	// The task for installing project dependencies (frozen).
	// Experimental.
	InstallCiTask() projen.Task
	// Returns the command to execute in order to install all dependencies (always frozen).
	// Experimental.
	InstallCommand() *string
	// The task for installing project dependencies (non-frozen).
	// Experimental.
	InstallTask() projen.Task
	// The SPDX license of this module.
	//
	// `undefined` if this package is not licensed.
	// Experimental.
	License() *string
	// The name of the lock file.
	// Experimental.
	LockFile() *string
	// Deprecated: use `addField(x, y)`.
	Manifest() interface{}
	// Maximum node version required by this package.
	// Default: - no maximum.
	//
	// Experimental.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Default: - no minimum.
	//
	// Experimental.
	MinNodeVersion() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// npm package access level.
	// Experimental.
	NpmAccess() NpmAccess
	// Should provenance statements be generated when package is published.
	// Experimental.
	NpmProvenance() *bool
	// The npm registry host (e.g. `registry.npmjs.org`).
	// Experimental.
	NpmRegistry() *string
	// npm registry (e.g. `https://registry.npmjs.org`). Use `npmRegistryHost` to get just the host name.
	// Experimental.
	NpmRegistryUrl() *string
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret() *string
	// The package manager to use.
	// Experimental.
	PackageManager() NodePackageManager
	// The name of the npm package.
	// Experimental.
	PackageName() *string
	// The version of PNPM to use if using PNPM as a package manager.
	// Default: "7".
	//
	// Experimental.
	PnpmVersion() *string
	// Experimental.
	Project() projen.Project
	// The command which executes "projen".
	// Deprecated: use `project.projenCommand` instead.
	ProjenCommand() *string
	// Options for privately hosted scoped packages.
	// Default: undefined.
	//
	// Experimental.
	ScopedPackagesOptions() *[]*ScopedPackagesOptions
	// Experimental.
	AddBin(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Experimental.
	AddBundledDeps(deps ...*string)
	// Defines normal dependencies.
	// Experimental.
	AddDeps(deps ...*string)
	// Defines development/test dependencies.
	// Experimental.
	AddDevDeps(deps ...*string)
	// Adds an `engines` requirement to your package.
	// Experimental.
	AddEngine(engine *string, version *string)
	// Directly set fields in `package.json`.
	// Experimental.
	AddField(name *string, value interface{})
	// Adds keywords to package.json (deduplicated).
	// Experimental.
	AddKeywords(keywords ...*string)
	// Defines resolutions for dependencies to change the normally resolved version of a dependency to something else.
	// Experimental.
	AddPackageResolutions(resolutions ...*string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Experimental.
	AddPeerDeps(deps ...*string)
	// Sets the package version.
	// Experimental.
	AddVersion(version *string)
	// Indicates if a script by the given name is defined.
	// Deprecated: Use `project.tasks.tryFind(name)`
	HasScript(name *string) *bool
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Removes an npm script (always successful).
	// Experimental.
	RemoveScript(name *string)
	// Add a npm package.json script.
	// Experimental.
	SetScript(name *string, command *string)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Attempt to resolve the currently installed version for a given dependency.
	// Experimental.
	TryResolveDependencyVersion(dependencyName *string) *string
}

// The jsii proxy struct for NodePackage
type jsiiProxy_NodePackage struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NodePackage) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) CodeArtifactOptions() *CodeArtifactOptions {
	var returns *CodeArtifactOptions
	_jsii_.Get(
		j,
		"codeArtifactOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) InstallAndUpdateLockfileCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installAndUpdateLockfileCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) InstallCiTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installCiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) InstallCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) License() *string {
	var returns *string
	_jsii_.Get(
		j,
		"license",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) LockFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		j,
		"npmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmProvenance() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"npmProvenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"npmRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmRegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"npmRegistryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmTokenSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"npmTokenSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) PackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) PackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) PnpmVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pnpmVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) ScopedPackagesOptions() *[]*ScopedPackagesOptions {
	var returns *[]*ScopedPackagesOptions
	_jsii_.Get(
		j,
		"scopedPackagesOptions",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodePackage(project projen.Project, options *NodePackageOptions) NodePackage {
	_init_.Initialize()

	if err := validateNewNodePackageParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodePackage{}

	_jsii_.Create(
		"projen.javascript.NodePackage",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNodePackage_Override(n NodePackage, project projen.Project, options *NodePackageOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NodePackage",
		[]interface{}{project, options},
		n,
	)
}

// Test whether the given construct is a component.
// Experimental.
func NodePackage_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodePackage_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NodePackage",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func NodePackage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodePackage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NodePackage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the `NodePackage` instance associated with a project or `undefined` if there is no NodePackage.
//
// Returns: A NodePackage, or undefined.
// Experimental.
func NodePackage_Of(project projen.Project) NodePackage {
	_init_.Initialize()

	if err := validateNodePackage_OfParameters(project); err != nil {
		panic(err)
	}
	var returns NodePackage

	_jsii_.StaticInvoke(
		"projen.javascript.NodePackage",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodePackage) AddBin(bins *map[string]*string) {
	if err := n.validateAddBinParameters(bins); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addBin",
		[]interface{}{bins},
	)
}

func (n *jsiiProxy_NodePackage) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addBundledDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDevDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddEngine(engine *string, version *string) {
	if err := n.validateAddEngineParameters(engine, version); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addEngine",
		[]interface{}{engine, version},
	)
}

func (n *jsiiProxy_NodePackage) AddField(name *string, value interface{}) {
	if err := n.validateAddFieldParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addField",
		[]interface{}{name, value},
	)
}

func (n *jsiiProxy_NodePackage) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addKeywords",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddPackageResolutions(resolutions ...*string) {
	args := []interface{}{}
	for _, a := range resolutions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addPackageResolutions",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addPeerDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddVersion(version *string) {
	if err := n.validateAddVersionParameters(version); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addVersion",
		[]interface{}{version},
	)
}

func (n *jsiiProxy_NodePackage) HasScript(name *string) *bool {
	if err := n.validateHasScriptParameters(name); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodePackage) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodePackage) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodePackage) RemoveScript(name *string) {
	if err := n.validateRemoveScriptParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

func (n *jsiiProxy_NodePackage) SetScript(name *string, command *string) {
	if err := n.validateSetScriptParameters(name, command); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"setScript",
		[]interface{}{name, command},
	)
}

func (n *jsiiProxy_NodePackage) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodePackage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodePackage) TryResolveDependencyVersion(dependencyName *string) *string {
	if err := n.validateTryResolveDependencyVersionParameters(dependencyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"tryResolveDependencyVersion",
		[]interface{}{dependencyName},
		&returns,
	)

	return returns
}

