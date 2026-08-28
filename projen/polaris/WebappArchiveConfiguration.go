package polaris


// Experimental.
type WebappArchiveConfiguration struct {
	// Specifies the path to the web application archive file or path to the directory containing the exploded web application.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Indicates whether the web-app should be checked to see if it is valid during capture.
	//
	// The validation check checks that there is a "/WEB-INF/web.xml" file and that > 20% of classes for the web application were captured.
	// Experimental.
	ValidateWebapp *bool `field:"optional" json:"validateWebapp" yaml:"validateWebapp"`
}

