package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines project tasks.
//
// Tasks extend the projen CLI by adding subcommands to it. Task definitions are
// synthesized into `.projen/tasks.json`.
// Experimental.
type Tasks interface {
	Component
	// All tasks.
	// Experimental.
	All() *[]Task
	// Returns a copy of the currently global environment for this project.
	// Experimental.
	Env() *map[string]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// Adds global environment.
	// Experimental.
	AddEnvironment(name *string, value *string)
	// Adds a task to a project.
	// Experimental.
	AddTask(name *string, options *TaskOptions) Task
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Experimental.
	RemoveTask(name *string) Task
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Finds a task by name.
	//
	// Returns `undefined` if the task cannot be found.
	// Experimental.
	TryFind(name *string) Task
}

// The jsii proxy struct for Tasks
type jsiiProxy_Tasks struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Tasks) All() *[]Task {
	var returns *[]Task
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tasks) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tasks) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tasks) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewTasks(project Project) Tasks {
	_init_.Initialize()

	if err := validateNewTasksParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_Tasks{}

	_jsii_.Create(
		"projen.Tasks",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewTasks_Override(t Tasks, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Tasks",
		[]interface{}{project},
		t,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Tasks_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTasks_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Tasks",
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
func Tasks_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTasks_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Tasks",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tasks) AddEnvironment(name *string, value *string) {
	if err := t.validateAddEnvironmentParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

func (t *jsiiProxy_Tasks) AddTask(name *string, options *TaskOptions) Task {
	if err := t.validateAddTaskParameters(name, options); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		t,
		"addTask",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tasks) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Tasks) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Tasks) RemoveTask(name *string) Task {
	if err := t.validateRemoveTaskParameters(name); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		t,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tasks) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Tasks) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tasks) TryFind(name *string) Task {
	if err := t.validateTryFindParameters(name); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		t,
		"tryFind",
		[]interface{}{name},
		&returns,
	)

	return returns
}

