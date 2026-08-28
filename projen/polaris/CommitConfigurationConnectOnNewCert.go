package polaris


// Indicates whether to trust self-signed certificates presented by Coverity Connect that are not currently trusted.
// Experimental.
type CommitConfigurationConnectOnNewCert string

const (
	// trust.
	// Experimental.
	CommitConfigurationConnectOnNewCert_TRUST CommitConfigurationConnectOnNewCert = "TRUST"
	// distrust.
	// Experimental.
	CommitConfigurationConnectOnNewCert_DISTRUST CommitConfigurationConnectOnNewCert = "DISTRUST"
)

