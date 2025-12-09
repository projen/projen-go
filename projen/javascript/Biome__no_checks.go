//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Biome) validateAddFilePatternParameters(pattern *string) error {
	return nil
}

func (b *jsiiProxy_Biome) validateAddOverrideParameters(override *biomeconfig.OverridePattern) error {
	return nil
}

func (b *jsiiProxy_Biome) validateExpandLinterRulesParameters(rules *biomeconfig.Rules) error {
	return nil
}

func validateBiome_IsComponentParameters(x interface{}) error {
	return nil
}

func validateBiome_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBiome_OfParameters(project projen.Project) error {
	return nil
}

func validateNewBiomeParameters(project NodeProject, options *BiomeOptions) error {
	return nil
}

