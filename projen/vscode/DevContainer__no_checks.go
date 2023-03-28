//go:build no_runtime_type_checking

package vscode

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DevContainer) validateAddDockerImageParameters(image projen.DevEnvironmentDockerImage) error {
	return nil
}

func (d *jsiiProxy_DevContainer) validateAddFeaturesParameters(features *[]*DevContainerFeature) error {
	return nil
}

func validateNewDevContainerParameters(project projen.Project, options *DevContainerOptions) error {
	return nil
}

