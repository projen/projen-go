package vscode

import (
	"github.com/projen/projen-go/projen"
)

// Constructor options for the DevContainer component.
//
// The default docker image used for GitHub Codespaces is defined here:.
// See: https://github.com/microsoft/vscode-dev-containers/tree/master/containers/codespaces-linux
//
// Experimental.
type DevContainerOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage projen.DevEnvironmentDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks *[]projen.Task `field:"optional" json:"tasks" yaml:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions *[]*string `field:"optional" json:"vscodeExtensions" yaml:"vscodeExtensions"`
}

