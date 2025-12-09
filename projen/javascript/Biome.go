package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/biomeconfig"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Biome component.
// Experimental.
type Biome interface {
	projen.Component
	// Biome configuration file content.
	// Experimental.
	File() projen.JsonFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Biome task.
	// Experimental.
	Task() projen.Task
	// Add a file pattern to biome.
	//
	// Use ! or !! to ignore a file pattern.
	// See: https://biomejs.dev/guides/configure-biome/#control-files-via-configuration
	//
	// Experimental.
	AddFilePattern(pattern *string)
	// Add a biome override to set rules for a specific file pattern.
	// See: https://biomejs.dev/reference/configuration/#overrides
	//
	// Experimental.
	AddOverride(override *biomeconfig.OverridePattern)
	// Expand the linting rules applied.
	//
	// Use `undefined` to remove the rule or group.
	//
	// Example:
	//   biome.expandLintingRules({
	//     style: undefined,
	//     suspicious: {
	//       noExplicitAny: undefined,
	//       noDuplicateCase: "info",
	//     }
	//   })
	//
	// See: https://biomejs.dev/reference/configuration/#linterrulesgroup
	//
	// Experimental.
	ExpandLinterRules(rules *biomeconfig.Rules)
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Biome
type jsiiProxy_Biome struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Biome) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Biome) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Biome) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Biome) Task() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}


// Experimental.
func NewBiome(project NodeProject, options *BiomeOptions) Biome {
	_init_.Initialize()

	if err := validateNewBiomeParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Biome{}

	_jsii_.Create(
		"projen.javascript.Biome",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewBiome_Override(b Biome, project NodeProject, options *BiomeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Biome",
		[]interface{}{project, options},
		b,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Biome_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBiome_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Biome",
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
func Biome_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBiome_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Biome",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Biome_Of(project projen.Project) Biome {
	_init_.Initialize()

	if err := validateBiome_OfParameters(project); err != nil {
		panic(err)
	}
	var returns Biome

	_jsii_.StaticInvoke(
		"projen.javascript.Biome",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Biome) AddFilePattern(pattern *string) {
	if err := b.validateAddFilePatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addFilePattern",
		[]interface{}{pattern},
	)
}

func (b *jsiiProxy_Biome) AddOverride(override *biomeconfig.OverridePattern) {
	if err := b.validateAddOverrideParameters(override); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{override},
	)
}

func (b *jsiiProxy_Biome) ExpandLinterRules(rules *biomeconfig.Rules) {
	if err := b.validateExpandLinterRulesParameters(rules); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"expandLinterRules",
		[]interface{}{rules},
	)
}

func (b *jsiiProxy_Biome) PostSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"postSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Biome) PreSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"preSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Biome) Synthesize() {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Biome) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

