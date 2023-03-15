// CDK for software projects
package projen


// Experimental.
type IgnoreFileOptions struct {
	// Filter out comment lines?
	// Experimental.
	FilterCommentLines *bool `field:"optional" json:"filterCommentLines" yaml:"filterCommentLines"`
	// Filter out blank/empty lines?
	// Experimental.
	FilterEmptyLines *bool `field:"optional" json:"filterEmptyLines" yaml:"filterEmptyLines"`
}

