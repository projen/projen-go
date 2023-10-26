//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Bundler) validateAddBundleParameters(entrypoint *string, options *AddBundleOptions) error {
	return nil
}

func validateBundler_IsComponentParameters(x interface{}) error {
	return nil
}

func validateBundler_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBundler_OfParameters(project projen.Project) error {
	return nil
}

func validateNewBundlerParameters(project projen.Project, options *BundlerOptions) error {
	return nil
}

