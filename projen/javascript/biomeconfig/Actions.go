package biomeconfig


// Experimental.
type Actions struct {
	// Selects the baseline set of assist actions.
	//
	// `recommended` enables Biome's
	// recommended actions, `all` enables all actions, and `none` starts with no actions
	// enabled. Group presets and explicit action settings override this preset.
	// Defaults to `recommended`.
	// Default: recommended`.
	//
	// Experimental.
	Preset PresetConfig `field:"optional" json:"preset" yaml:"preset"`
	// Enables or disables Biome's recommended assist actions across all action groups.
	//
	// Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Configures source-level actions such as organizing imports and sorting declarations.
	//
	// These actions are exposed through editor source actions and through CLI checking and application.
	// Experimental.
	Source *Source `field:"optional" json:"source" yaml:"source"`
}

