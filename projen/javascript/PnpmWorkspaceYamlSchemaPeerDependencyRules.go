package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaPeerDependencyRules struct {
	// Any peer dependency matching the pattern will be resolved from any version, regardless of the range specified in "peerDependencies".
	// Experimental.
	AllowAny *[]*string `field:"optional" json:"allowAny" yaml:"allowAny"`
	// Unmet peer dependency warnings will not be printed for peer dependencies of the specified range.
	// Experimental.
	AllowedVersions interface{} `field:"optional" json:"allowedVersions" yaml:"allowedVersions"`
	// pnpm will not print warnings about missing peer dependencies from this list.
	// Experimental.
	IgnoreMissing *[]*string `field:"optional" json:"ignoreMissing" yaml:"ignoreMissing"`
}

