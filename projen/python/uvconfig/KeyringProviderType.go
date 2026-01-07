package uvconfig


// Keyring provider type to use for credential lookup.
// Experimental.
type KeyringProviderType string

const (
	// Do not use keyring for credential lookup.
	//
	// (disabled).
	// Experimental.
	KeyringProviderType_DISABLED KeyringProviderType = "DISABLED"
	// Use the `keyring` command for credential lookup.
	//
	// (subprocess).
	// Experimental.
	KeyringProviderType_SUBPROCESS KeyringProviderType = "SUBPROCESS"
)

