package sonarqube


// Options for `sonar.typescript.*` properties.
// Experimental.
type SonarqubeTypescriptOptions struct {
	// Path to the TypeScript configuration file.
	//
	// Maps to `sonar.typescript.tsconfigPath`.
	// Default: - not set.
	//
	// Experimental.
	TsconfigPath *string `field:"optional" json:"tsconfigPath" yaml:"tsconfigPath"`
}

