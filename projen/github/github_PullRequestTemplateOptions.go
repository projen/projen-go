package github


// Options for `PullRequestTemplate`.
// Experimental.
type PullRequestTemplateOptions struct {
	// The contents of the template.
	//
	// You can use `addLine()` to add additional lines.
	// Experimental.
	Lines *[]*string `field:"optional" json:"lines" yaml:"lines"`
}

