package polaris


// Experimental.
type AnalyzeConnectConfiguration struct {
	// Absolute URL of where to perform Coverity Connect analysis.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The authentication key file to use when authenticating to Coverity Connect to perform analysis.
	//
	// By default, the file located at $HOME/.coverity/ak-<hostname>-<port> is used.
	// Experimental.
	AuthKeyFile *string `field:"optional" json:"authKeyFile" yaml:"authKeyFile"`
	// File containing additional certificates to trust in addition to the ones in the system certificate store and the Coverity TFT store.
	//
	// By default system CA certificates are used.
	// Experimental.
	CaCertsFile *string `field:"optional" json:"caCertsFile" yaml:"caCertsFile"`
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
	// Artifacts to upload following analysis when the analysis location is Connect.
	// Experimental.
	UploadArtifacts AnalyzeConnectConfigurationUploadArtifacts `field:"optional" json:"uploadArtifacts" yaml:"uploadArtifacts"`
}

