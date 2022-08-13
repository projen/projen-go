// CDK for software projects
package projen


// Choices for how to display commented out options in projenrc files.
//
// Does not apply to projenrc.json files.
// Experimental.
type InitProjectOptionHints string

const (
	// Display all possible options (grouped by which interface they belong to).
	// Experimental.
	InitProjectOptionHints_ALL InitProjectOptionHints = "ALL"
	// Display only featured options, in alphabetical order.
	// Experimental.
	InitProjectOptionHints_FEATURED InitProjectOptionHints = "FEATURED"
	// Display no extra options.
	// Experimental.
	InitProjectOptionHints_NONE InitProjectOptionHints = "NONE"
)

