package gitlab


// The content kind of what users can download via url.
// Experimental.
type LinkType string

const (
	// Experimental.
	LinkType_IMAGE LinkType = "IMAGE"
	// Experimental.
	LinkType_OTHER LinkType = "OTHER"
	// Experimental.
	LinkType_PACKAGE LinkType = "PACKAGE"
	// Experimental.
	LinkType_RUNBOOK LinkType = "RUNBOOK"
)

