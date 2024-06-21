//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func (i *jsiiProxy_ICompareString) validateCompareParameters(a *string, b *string) error {
	if a == nil {
		return fmt.Errorf("parameter a is required, but nil was provided")
	}

	if b == nil {
		return fmt.Errorf("parameter b is required, but nil was provided")
	}

	return nil
}

