package javascript


// https://yarnpkg.com/configuration/yarnrc#npmRegistries.
// Experimental.
type YarnNpmRegistry struct {
	// Experimental.
	NpmAlwaysAuth *bool `field:"optional" json:"npmAlwaysAuth" yaml:"npmAlwaysAuth"`
	// Experimental.
	NpmAuthIdent *string `field:"optional" json:"npmAuthIdent" yaml:"npmAuthIdent"`
	// Experimental.
	NpmAuthToken *string `field:"optional" json:"npmAuthToken" yaml:"npmAuthToken"`
}

