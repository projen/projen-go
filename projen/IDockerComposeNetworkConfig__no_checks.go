//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDockerComposeNetworkConfig) validateAddNetworkConfigurationParameters(networkName *string, configuration *DockerComposeNetworkConfig) error {
	return nil
}

