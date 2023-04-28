//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"
)

func validateTypescriptConfigExtends_FromPathsParameters(paths *[]*string) error {
	if paths == nil {
		return fmt.Errorf("parameter paths is required, but nil was provided")
	}

	return nil
}

func validateTypescriptConfigExtends_FromTypescriptConfigsParameters(configs *[]TypescriptConfig) error {
	if configs == nil {
		return fmt.Errorf("parameter configs is required, but nil was provided")
	}

	return nil
}

