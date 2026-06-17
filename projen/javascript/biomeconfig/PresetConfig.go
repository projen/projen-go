package biomeconfig


// A preset configuration for enabling a set of rules.
// Experimental.
type PresetConfig string

const (
	// recommended.
	// Experimental.
	PresetConfig_RECOMMENDED PresetConfig = "RECOMMENDED"
	// all.
	// Experimental.
	PresetConfig_ALL PresetConfig = "ALL"
	// none.
	// Experimental.
	PresetConfig_NONE PresetConfig = "NONE"
)

