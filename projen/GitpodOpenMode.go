// CDK for software projects
package projen


// Configure how the terminal should be opened relative to the previous task.
// Experimental.
type GitpodOpenMode string

const (
	// Opens in the same tab group right after the previous tab.
	// Experimental.
	GitpodOpenMode_TAB_AFTER GitpodOpenMode = "TAB_AFTER"
	// Opens in the same tab group left before the previous tab.
	// Experimental.
	GitpodOpenMode_TAB_BEFORE GitpodOpenMode = "TAB_BEFORE"
	// Splits and adds the terminal to the right.
	// Experimental.
	GitpodOpenMode_SPLIT_RIGHT GitpodOpenMode = "SPLIT_RIGHT"
	// Splits and adds the terminal to the left.
	// Experimental.
	GitpodOpenMode_SPLIT_LEFT GitpodOpenMode = "SPLIT_LEFT"
	// Splits and adds the terminal to the top.
	// Experimental.
	GitpodOpenMode_SPLIT_TOP GitpodOpenMode = "SPLIT_TOP"
	// Splits and adds the terminal to the bottom.
	// Experimental.
	GitpodOpenMode_SPLIT_BOTTOM GitpodOpenMode = "SPLIT_BOTTOM"
)

