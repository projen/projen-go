package sonarqube


// File options for the generated `sonar-project.properties` file.
// Experimental.
type SonarqubeFileOptions struct {
	// A comment to include at the top of the file.
	// Default: - no additional comment.
	//
	// Experimental.
	Comment *[]*string `field:"optional" json:"comment" yaml:"comment"`
	// Whether the generated file should be committed to git.
	// Default: true.
	//
	// Experimental.
	Committed *bool `field:"optional" json:"committed" yaml:"committed"`
	// Adds the projen marker to the file.
	// Default: - marker will be included as long as the project is not ejected.
	//
	// Experimental.
	Marker *bool `field:"optional" json:"marker" yaml:"marker"`
	// Whether the generated file should be readonly.
	// Default: true.
	//
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

