//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func (i *jsiiProxy_IDockerComposeNetworkBinding) validateBindParameters(networkConfig IDockerComposeNetworkConfig) error {
	if networkConfig == nil {
		return fmt.Errorf("parameter networkConfig is required, but nil was provided")
	}

	return nil
}

