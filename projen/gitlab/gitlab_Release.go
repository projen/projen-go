package gitlab


// Indicates that the job creates a Release.
// Experimental.
type Release struct {
	// Specifies the longer description of the Release.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The tag_name must be specified.
	//
	// It can refer to an existing Git tag or can be specified by the user.
	// Experimental.
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
	// Experimental.
	Assets *Assets `field:"optional" json:"assets" yaml:"assets"`
	// The title of each milestone the release is associated with.
	// Experimental.
	Milestones *[]*string `field:"optional" json:"milestones" yaml:"milestones"`
	// The Release name.
	//
	// If omitted, it is populated with the value of release: tag_name.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If the release: tag_name doesn’t exist yet, the release is created from ref.
	//
	// ref can be a commit SHA, another tag name, or a branch name.
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// The date and time when the release is ready.
	//
	// Defaults to the current date and time if not defined. Should be enclosed in quotes and expressed in ISO 8601 format.
	// Experimental.
	ReleasedAt *string `field:"optional" json:"releasedAt" yaml:"releasedAt"`
}

