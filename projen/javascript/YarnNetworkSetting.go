package javascript


// https://yarnpkg.com/configuration/yarnrc#networkSettings.
// Experimental.
type YarnNetworkSetting struct {
	// Deprecated: - use httpsCaFilePath in Yarn v4 and newer.
	CaFilePath *string `field:"optional" json:"caFilePath" yaml:"caFilePath"`
	// Experimental.
	EnableNetwork *bool `field:"optional" json:"enableNetwork" yaml:"enableNetwork"`
	// Experimental.
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// Experimental.
	HttpsCaFilePath *string `field:"optional" json:"httpsCaFilePath" yaml:"httpsCaFilePath"`
	// Experimental.
	HttpsCertFilePath *string `field:"optional" json:"httpsCertFilePath" yaml:"httpsCertFilePath"`
	// Experimental.
	HttpsKeyFilePath *string `field:"optional" json:"httpsKeyFilePath" yaml:"httpsKeyFilePath"`
	// Experimental.
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
}

