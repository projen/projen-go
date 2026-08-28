package polaris


// Coverity Connect configuration to use when committing defects to Coverity Connect.
// Experimental.
type CommitConfigurationConnect struct {
	// The name of the stream to commit the results to.
	// Experimental.
	Stream *string `field:"required" json:"stream" yaml:"stream"`
	// Absolute URL of where to commit the Coverity Connect results.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The authentication key file to use when authenticating to Coverity Connect to commit defects.
	//
	// By default, the file located at $HOME/.coverity/ak-<hostname>-<port> is used.
	// Experimental.
	AuthKeyFile *string `field:"optional" json:"authKeyFile" yaml:"authKeyFile"`
	// File containing additional certificates to trust in addition to the ones in the system certificate store and the Coverity TFT store.
	//
	// By default system CA certificates are used.
	// Experimental.
	CaCertsFile *string `field:"optional" json:"caCertsFile" yaml:"caCertsFile"`
	// If true, analysis results will not be committed to Coverity Connect.
	//
	// Instead, results compared to a reference snapshot may be saved locally as specified by the "commit.local" settings.
	// Experimental.
	ComparisonOnly *bool `field:"optional" json:"comparisonOnly" yaml:"comparisonOnly"`
	// Output file to which analysis results should be written instead of being committed to Coverity Connect.
	//
	// The output includes a comparison against the latest snapshot for the specified stream.
	// Experimental.
	ComparisonReport *string `field:"optional" json:"comparisonReport" yaml:"comparisonReport"`
	// Additional arguments to pass to "cov-commit-defects" during the commit phase.
	// Experimental.
	CovCommitDefectsArgs *[]*string `field:"optional" json:"covCommitDefectsArgs" yaml:"covCommitDefectsArgs"`
	// A description for the committed snapshot.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether to trust self-signed certificates presented by Coverity Connect that are not currently trusted.
	// Experimental.
	OnNewCert CommitConfigurationConnectOnNewCert `field:"optional" json:"onNewCert" yaml:"onNewCert"`
	// The name of the project to use when creating a new stream.
	//
	// Ignored when stream creation is not needed. By default the stream name is used.
	// Experimental.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// File containing the client certificate in PEM format, that should be presented to the proxy when making a request.
	// Experimental.
	ProxyClientCertFile *string `field:"optional" json:"proxyClientCertFile" yaml:"proxyClientCertFile"`
	// File containing the client certificate private key in PEM format, for the proxy-client-cert-file.
	// Experimental.
	ProxyClientKeyFile *string `field:"optional" json:"proxyClientKeyFile" yaml:"proxyClientKeyFile"`
	// URL for a forward proxy to use when communicating with Coverity Connect.
	//
	// Must be an https URL.
	// Experimental.
	ProxyUrl *string `field:"optional" json:"proxyUrl" yaml:"proxyUrl"`
	// The name of the source control management system.
	// Experimental.
	Scm CommitConfigurationConnectScm `field:"optional" json:"scm" yaml:"scm"`
	// Specifies how to select a reference snapshot to use for a comparison report.
	// Experimental.
	Snapshot *SnapshotConfiguration `field:"optional" json:"snapshot" yaml:"snapshot"`
	// Specifies how new defects should be handled.
	// Experimental.
	Triage *CommitConfigurationConnectTriage `field:"optional" json:"triage" yaml:"triage"`
	// Artifacts to upload following analysis when the analysis location is Connect.
	// Experimental.
	UploadArtifacts CommitConfigurationConnectUploadArtifacts `field:"optional" json:"uploadArtifacts" yaml:"uploadArtifacts"`
	// A project version for the committed snapshot.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

