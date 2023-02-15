package gitlab


// Link configuration for an asset.
// Experimental.
type Link struct {
	// The name of the link.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL to download a file.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The redirect link to the url.
	// Experimental.
	Filepath *string `field:"optional" json:"filepath" yaml:"filepath"`
	// The content kind of what users can download via url.
	// Experimental.
	LinkType LinkType `field:"optional" json:"linkType" yaml:"linkType"`
}

