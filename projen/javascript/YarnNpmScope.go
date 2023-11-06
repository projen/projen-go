package javascript


// https://yarnpkg.com/configuration/yarnrc#npmScopes.
// Experimental.
type YarnNpmScope struct {
	// Experimental.
	NpmAlwaysAuth *bool `field:"optional" json:"npmAlwaysAuth" yaml:"npmAlwaysAuth"`
	// Experimental.
	NpmAuthIdent *string `field:"optional" json:"npmAuthIdent" yaml:"npmAuthIdent"`
	// Experimental.
	NpmAuthToken *string `field:"optional" json:"npmAuthToken" yaml:"npmAuthToken"`
	// Experimental.
	NpmPublishRegistry *string `field:"optional" json:"npmPublishRegistry" yaml:"npmPublishRegistry"`
	// Experimental.
	NpmRegistryServer *string `field:"optional" json:"npmRegistryServer" yaml:"npmRegistryServer"`
}

