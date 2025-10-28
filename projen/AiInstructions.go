package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Generates instruction files for AI coding assistants with projen-specific guidance.
//
// This component creates configuration files that help AI tools like GitHub Copilot,
// Cursor IDE, Claude Code, and Amazon Q understand that the project is managed by projen
// and should follow projen conventions.
//
// Example:
//   const project = new TypeScriptProject({
//     name: "my-project",
//     defaultReleaseBranch: "main",
//   });
//
//   // Basic usage - generates files for all supported AI agents
//   new AiInstructions(project);
//
//   // Custom usage - specify which agents and add custom instructions
//   new AiInstructions(project, {
//     agents: [AiAgent.GITHUB_COPILOT, AiAgent.CURSOR],
//     agentSpecificInstructions: {
//       [AiAgent.GITHUB_COPILOT]: ["Always use descriptive commit messages."],
//     },
//   });
//
//   // Add more instructions after instantiation
//   const ai = new AiInstructions(project);
//   ai.addInstructions("Use functional programming patterns.");
//   ai.addInstructions("Always write comprehensive tests.");
//
// Experimental.
type AiInstructions interface {
	Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// Add instructions for a specific AI agent.
	//
	// This can also be used to add instructions for an AI agent that was previously not enabled.
	//
	// Example:
	//   aiInstructions.addAgentSpecificInstructions(AiAgent.GITHUB_COPILOT, "Use descriptive commit messages.");
	//
	// Experimental.
	AddAgentSpecificInstructions(agent AiAgent, instructions ...*string)
	// Adds instructions that will be included for all selected AI agents.
	//
	// Example:
	//   aiInstructions.addInstructions("Always use TypeScript strict mode.");
	//   aiInstructions.addInstructions("Prefer functional programming.", "Avoid mutations.");
	//
	// Experimental.
	AddInstructions(instructions ...*string)
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

// The jsii proxy struct for AiInstructions
type jsiiProxy_AiInstructions struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_AiInstructions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructions) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAiInstructions(project Project, options *AiInstructionsOptions) AiInstructions {
	_init_.Initialize()

	if err := validateNewAiInstructionsParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiInstructions{}

	_jsii_.Create(
		"projen.AiInstructions",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAiInstructions_Override(a AiInstructions, project Project, options *AiInstructionsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.AiInstructions",
		[]interface{}{project, options},
		a,
	)
}

// Returns development best practices instructions for AI agents.
// Experimental.
func AiInstructions_BestPractices(project Project) *string {
	_init_.Initialize()

	if err := validateAiInstructions_BestPracticesParameters(project); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"projen.AiInstructions",
		"bestPractices",
		[]interface{}{project},
		&returns,
	)

	return returns
}

// Test whether the given construct is a component.
// Experimental.
func AiInstructions_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiInstructions_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.AiInstructions",
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
func AiInstructions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiInstructions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.AiInstructions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns projen-specific instructions for AI agents.
// Experimental.
func AiInstructions_Projen(project Project) *string {
	_init_.Initialize()

	if err := validateAiInstructions_ProjenParameters(project); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"projen.AiInstructions",
		"projen",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiInstructions) AddAgentSpecificInstructions(agent AiAgent, instructions ...*string) {
	if err := a.validateAddAgentSpecificInstructionsParameters(agent); err != nil {
		panic(err)
	}
	args := []interface{}{agent}
	for _, a := range instructions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addAgentSpecificInstructions",
		args,
	)
}

func (a *jsiiProxy_AiInstructions) AddInstructions(instructions ...*string) {
	args := []interface{}{}
	for _, a := range instructions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addInstructions",
		args,
	)
}

func (a *jsiiProxy_AiInstructions) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiInstructions) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiInstructions) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiInstructions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

