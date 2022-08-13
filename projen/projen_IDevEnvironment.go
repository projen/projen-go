// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract interface for container-based development environments, such as Gitpod and GitHub Codespaces.
// Experimental.
type IDevEnvironment interface {
	// Add a custom Docker image or Dockerfile for the container.
	// Experimental.
	AddDockerImage(image DevEnvironmentDockerImage)
	// Adds ports that should be exposed (forwarded) from the container.
	// Experimental.
	AddPorts(ports ...*string)
	// Adds tasks to run when the container starts.
	// Experimental.
	AddTasks(tasks ...Task)
	// Adds a list of VSCode extensions that should be automatically installed in the container.
	// Experimental.
	AddVscodeExtensions(extensions ...*string)
}

// The jsii proxy for IDevEnvironment
type jsiiProxy_IDevEnvironment struct {
	_ byte // padding
}

func (i *jsiiProxy_IDevEnvironment) AddDockerImage(image DevEnvironmentDockerImage) {
	_jsii_.InvokeVoid(
		i,
		"addDockerImage",
		[]interface{}{image},
	)
}

func (i *jsiiProxy_IDevEnvironment) AddPorts(ports ...*string) {
	args := []interface{}{}
	for _, a := range ports {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addPorts",
		args,
	)
}

func (i *jsiiProxy_IDevEnvironment) AddTasks(tasks ...Task) {
	args := []interface{}{}
	for _, a := range tasks {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTasks",
		args,
	)
}

func (i *jsiiProxy_IDevEnvironment) AddVscodeExtensions(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addVscodeExtensions",
		args,
	)
}

