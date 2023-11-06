package javascript


// https://yarnpkg.com/configuration/yarnrc#nmHoistingLimits.
// Experimental.
type YarnNmHoistingLimit string

const (
	// Experimental.
	YarnNmHoistingLimit_DEPENDENCIES YarnNmHoistingLimit = "DEPENDENCIES"
	// Experimental.
	YarnNmHoistingLimit_NONE YarnNmHoistingLimit = "NONE"
	// Experimental.
	YarnNmHoistingLimit_WORKSPACES YarnNmHoistingLimit = "WORKSPACES"
)

