package javascript


// https://yarnpkg.com/configuration/yarnrc#nodeLinker.
// Experimental.
type YarnNodeLinker string

const (
	// Experimental.
	YarnNodeLinker_PNP YarnNodeLinker = "PNP"
	// Experimental.
	YarnNodeLinker_PNPM YarnNodeLinker = "PNPM"
	// Experimental.
	YarnNodeLinker_NODE_MODULES YarnNodeLinker = "NODE_MODULES"
)

