//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"
)

func validateNewWatchPluginParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

