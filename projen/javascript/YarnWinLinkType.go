package javascript


// https://yarnpkg.com/configuration/yarnrc#winLinkType.
// Experimental.
type YarnWinLinkType string

const (
	// Experimental.
	YarnWinLinkType_JUNCTIONS YarnWinLinkType = "JUNCTIONS"
	// Experimental.
	YarnWinLinkType_SYMLINKS YarnWinLinkType = "SYMLINKS"
)

