//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICompareString) validateCompareParameters(a *string, b *string) error {
	return nil
}

