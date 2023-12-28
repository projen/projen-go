//go:build !no_runtime_type_checking

package typescript

import (
	"fmt"
)

func validateTsJestBabelConfig_CustomParameters(config *map[string]interface{}) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func validateTsJestBabelConfig_FromFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

