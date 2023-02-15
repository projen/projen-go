// CDK for software projects
package projen


// Resolve options.
// Experimental.
type ResolveOptions struct {
	// Context arguments.
	// Experimental.
	Args *[]interface{} `field:"optional" json:"args" yaml:"args"`
	// Omits empty arrays and objects.
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
}

