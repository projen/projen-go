//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"
)

func validateDestination_FileParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

