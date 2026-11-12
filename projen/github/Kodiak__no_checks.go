//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_Kodiak) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (k *jsiiProxy_Kodiak) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateKodiak_IsComponentParameters(x interface{}) error {
	return nil
}

func validateKodiak_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKodiakParameters(github GitHub, options *KodiakOptions) error {
	return nil
}

