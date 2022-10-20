//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"
)

func validateUpgradeDependenciesSchedule_ExpressionsParameters(cron *[]*string) error {
	if cron == nil {
		return fmt.Errorf("parameter cron is required, but nil was provided")
	}

	return nil
}

