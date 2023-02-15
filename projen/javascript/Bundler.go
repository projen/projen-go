package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Adds support for bundling JavaScript applications and dependencies into a single file.
//
// In the future, this will also supports bundling websites.
// Experimental.
type Bundler interface {
	projen.Component
	// Root bundle directory.
	// Experimental.
	Bundledir() *string
	// Gets or creates the singleton "bundle" task of the project.
	//
	// If the project doesn't have a "bundle" task, it will be created and spawned
	// during the pre-compile phase.
	// Experimental.
	BundleTask() projen.Task
	// The semantic version requirement for `esbuild` (if defined).
	// Experimental.
	EsbuildVersion() *string
	// Experimental.
	Project() projen.Project
	// Adds a task to the project which bundles a specific entrypoint and all of its dependencies into a single javascript output file.
	// Experimental.
	AddBundle(entrypoint *string, options *AddBundleOptions) *Bundle
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

// The jsii proxy struct for Bundler
type jsiiProxy_Bundler struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Bundler) Bundledir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundledir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bundler) BundleTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"bundleTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bundler) EsbuildVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esbuildVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bundler) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Creates a `Bundler`.
// Experimental.
func NewBundler(project projen.Project, options *BundlerOptions) Bundler {
	_init_.Initialize()

	if err := validateNewBundlerParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Bundler{}

	_jsii_.Create(
		"projen.javascript.Bundler",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Creates a `Bundler`.
// Experimental.
func NewBundler_Override(b Bundler, project projen.Project, options *BundlerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Bundler",
		[]interface{}{project, options},
		b,
	)
}

// Returns the `Bundler` instance associated with a project or `undefined` if there is no Bundler.
//
// Returns: A bundler.
// Experimental.
func Bundler_Of(project projen.Project) Bundler {
	_init_.Initialize()

	if err := validateBundler_OfParameters(project); err != nil {
		panic(err)
	}
	var returns Bundler

	_jsii_.StaticInvoke(
		"projen.javascript.Bundler",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bundler) AddBundle(entrypoint *string, options *AddBundleOptions) *Bundle {
	if err := b.validateAddBundleParameters(entrypoint, options); err != nil {
		panic(err)
	}
	var returns *Bundle

	_jsii_.Invoke(
		b,
		"addBundle",
		[]interface{}{entrypoint, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bundler) PostSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"postSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Bundler) PreSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"preSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Bundler) Synthesize() {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		nil, // no parameters
	)
}

