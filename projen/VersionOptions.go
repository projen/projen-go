// CDK for software projects
package projen


// Options for `Version`.
// Experimental.
type VersionOptions struct {
	// The name of the directory into which `changelog.md` and `version.txt` files are emitted.
	// Experimental.
	ArtifactsDirectory *string `field:"required" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// A name of a .json file to set the `version` field in after a bump.
	//
	// Example:
	//   "package.json"
	//
	// Experimental.
	VersionInputFile *string `field:"required" json:"versionInputFile" yaml:"versionInputFile"`
	// The tag prefix corresponding to this version.
	// Experimental.
	TagPrefix *string `field:"optional" json:"tagPrefix" yaml:"tagPrefix"`
	// Custom configuration for versionrc file used by standard-release.
	// Experimental.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
}

