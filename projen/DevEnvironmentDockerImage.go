package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Options for specifying the Docker image of the container.
// Experimental.
type DevEnvironmentDockerImage interface {
	// The relative path of a Dockerfile that defines the container contents.
	// Experimental.
	DockerFile() *string
	// A publicly available Docker image.
	// Experimental.
	Image() *string
}

// The jsii proxy struct for DevEnvironmentDockerImage
type jsiiProxy_DevEnvironmentDockerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_DevEnvironmentDockerImage) DockerFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevEnvironmentDockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// The relative path of a Dockerfile that defines the container contents.
//
// Example:
//   '.gitpod.Docker'
//
// Experimental.
func DevEnvironmentDockerImage_FromFile(dockerFile *string) DevEnvironmentDockerImage {
	_init_.Initialize()

	if err := validateDevEnvironmentDockerImage_FromFileParameters(dockerFile); err != nil {
		panic(err)
	}
	var returns DevEnvironmentDockerImage

	_jsii_.StaticInvoke(
		"projen.DevEnvironmentDockerImage",
		"fromFile",
		[]interface{}{dockerFile},
		&returns,
	)

	return returns
}

// A publicly available Docker image.
//
// Example:
//   'ubuntu:latest'
//
// Experimental.
func DevEnvironmentDockerImage_FromImage(image *string) DevEnvironmentDockerImage {
	_init_.Initialize()

	if err := validateDevEnvironmentDockerImage_FromImageParameters(image); err != nil {
		panic(err)
	}
	var returns DevEnvironmentDockerImage

	_jsii_.StaticInvoke(
		"projen.DevEnvironmentDockerImage",
		"fromImage",
		[]interface{}{image},
		&returns,
	)

	return returns
}

