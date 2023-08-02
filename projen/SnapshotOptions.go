package projen


// Options for the Snapshot synthesis.
// Experimental.
type SnapshotOptions struct {
	// Parse .json files as a JS object for improved inspection. This will fail if the contents are invalid JSON.
	// Default: true parse .json files into an object
	//
	// Experimental.
	ParseJson *bool `field:"optional" json:"parseJson" yaml:"parseJson"`
}

