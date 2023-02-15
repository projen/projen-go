package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Experimental.
type Mergify interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// Experimental.
	AddQueue(queue *MergifyQueue)
	// Experimental.
	AddRule(rule *MergifyRule)
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

// The jsii proxy struct for Mergify
type jsiiProxy_Mergify struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Mergify) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMergify(github GitHub, options *MergifyOptions) Mergify {
	_init_.Initialize()

	if err := validateNewMergifyParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mergify{}

	_jsii_.Create(
		"projen.github.Mergify",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMergify_Override(m Mergify, github GitHub, options *MergifyOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Mergify",
		[]interface{}{github, options},
		m,
	)
}

func (m *jsiiProxy_Mergify) AddQueue(queue *MergifyQueue) {
	if err := m.validateAddQueueParameters(queue); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addQueue",
		[]interface{}{queue},
	)
}

func (m *jsiiProxy_Mergify) AddRule(rule *MergifyRule) {
	if err := m.validateAddRuleParameters(rule); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addRule",
		[]interface{}{rule},
	)
}

func (m *jsiiProxy_Mergify) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mergify) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mergify) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

