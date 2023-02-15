// CDK for software projects
package projen


// Options for the Snapshot synthesis.
// Experimental.
type SnapshotOptions struct {
	// Parse .json files as a JS object for improved inspection. This will fail if the contents are invalid JSON.
	// Experimental.
	ParseJson *bool `field:"optional" json:"parseJson" yaml:"parseJson"`
}

