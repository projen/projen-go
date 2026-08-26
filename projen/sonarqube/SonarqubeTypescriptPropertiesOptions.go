package sonarqube


// Options for `SonarqubeTypescriptProperties`.
//
// Extends base options with TypeScript-specific defaults.
// Experimental.
type SonarqubeTypescriptPropertiesOptions struct {
	// The project's unique key.
	//
	// Can include up to 400 characters. Allowed characters:
	// letters, digits, dash, underscore, periods, and colons.
	//
	// Maps to `sonar.projectKey`. This parameter is mandatory.
	// Experimental.
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Coverage-related options (`sonar.coverage.*`).
	// Default: - no coverage configuration.
	//
	// Experimental.
	Coverage *SonarqubeCoverageOptions `field:"optional" json:"coverage" yaml:"coverage"`
	// Duplication detection options (`sonar.cpd.*`).
	// Default: - no CPD configuration.
	//
	// Experimental.
	Cpd *SonarqubeCpdOptions `field:"optional" json:"cpd" yaml:"cpd"`
	// Comma-separated file path patterns to exclude from the analysis scope.
	//
	// Maps to `sonar.exclusions`.
	// Default: - no exclusions.
	//
	// Experimental.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Additional arbitrary properties to include in the configuration.
	//
	// Use this for properties not covered by the typed options.
	// Keys use dot-notation (e.g., `sonar.java.binaries`).
	//
	// These are applied as overrides after the typed options above, so a key
	// that is a prefix of a typed option (e.g. `"sonar.coverage"`) replaces
	// that entire subtree rather than merging with it.
	// Default: - no additional properties.
	//
	// Experimental.
	ExtraProperties *map[string]*string `field:"optional" json:"extraProperties" yaml:"extraProperties"`
	// Options for the generated properties file.
	// Default: - default file options.
	//
	// Experimental.
	FileOptions *SonarqubeFileOptions `field:"optional" json:"fileOptions" yaml:"fileOptions"`
	// JavaScript-specific options (`sonar.javascript.*`).
	// Default: - no JavaScript configuration.
	//
	// Experimental.
	Javascript *SonarqubeJavascriptOptions `field:"optional" json:"javascript" yaml:"javascript"`
	// The language for analysis.
	//
	// Maps to `sonar.language`.
	// Default: - auto-detected.
	//
	// Experimental.
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Logging options (`sonar.log.*`).
	// Default: - INFO level.
	//
	// Experimental.
	Log *SonarqubeLogOptions `field:"optional" json:"log" yaml:"log"`
	// The key of the organization to which the project belongs.
	//
	// Maps to `sonar.organization`. Mandatory for SonarQube Cloud.
	// Default: - no organization.
	//
	// Experimental.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The quality profile name.
	//
	// Maps to `sonar.profile`.
	// Default: - uses the default profile configured on the server.
	//
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The project's base directory when the analysis needs to take place in a directory other than the one from which it was started.
	//
	// Maps to `sonar.projectBaseDir`.
	// Default: - the directory from which the analysis was started.
	//
	// Experimental.
	ProjectBaseDir *string `field:"optional" json:"projectBaseDir" yaml:"projectBaseDir"`
	// Name of the project displayed on the web interface.
	//
	// Maps to `sonar.projectName`.
	// Default: - not set.
	//
	// Experimental.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// The project version.
	//
	// Maps to `sonar.projectVersion`.
	// Default: - not set.
	//
	// Experimental.
	ProjectVersion *string `field:"optional" json:"projectVersion" yaml:"projectVersion"`
	// Quality gate options (`sonar.qualitygate.*`).
	// Default: - quality gate not awaited.
	//
	// Experimental.
	Qualitygate *SonarqubeQualityGateOptions `field:"optional" json:"qualitygate" yaml:"qualitygate"`
	// The SonarQube Cloud instance's region.
	//
	// Maps to `sonar.region`.
	// Default: SonarqubeRegion.EU
	//
	// Experimental.
	Region SonarqubeRegion `field:"optional" json:"region" yaml:"region"`
	// Rust-specific options (`sonar.rust.*`).
	// Default: - no Rust configuration.
	//
	// Experimental.
	Rust *SonarqubeRustOptions `field:"optional" json:"rust" yaml:"rust"`
	// SCM-related options (`sonar.scm.*`).
	// Default: - no SCM configuration.
	//
	// Experimental.
	Scm *SonarqubeScmOptions `field:"optional" json:"scm" yaml:"scm"`
	// Encoding of the source files.
	//
	// Maps to `sonar.sourceEncoding`.
	// Default: - system encoding.
	//
	// Experimental.
	SourceEncoding *string `field:"optional" json:"sourceEncoding" yaml:"sourceEncoding"`
	// Comma-separated paths to directories containing main source code (non-test code).
	//
	// Maps to `sonar.sources`.
	// Default: - the project base directory.
	//
	// Experimental.
	Sources *string `field:"optional" json:"sources" yaml:"sources"`
	// Comma-separated paths to directories containing test code.
	//
	// Maps to `sonar.tests`.
	// Default: - no test code analyzed.
	//
	// Experimental.
	Tests *string `field:"optional" json:"tests" yaml:"tests"`
	// TypeScript-specific options (`sonar.typescript.*`).
	// Default: - no TypeScript configuration.
	//
	// Experimental.
	Typescript *SonarqubeTypescriptOptions `field:"optional" json:"typescript" yaml:"typescript"`
}

