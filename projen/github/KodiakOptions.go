package github


// Configure Kodiak (`.kodiak.toml`).
// See: https://kodiakhq.com/docs/config-reference
//
// Experimental.
type KodiakOptions struct {
	// Auto-approval settings (`[approve]`).
	// Default: - no auto-approval rules configured.
	//
	// Experimental.
	Approve *KodiakApproveOptions `field:"optional" json:"approve" yaml:"approve"`
	// Merge settings (`[merge]`).
	// Default: - default Kodiak merge behavior.
	//
	// Experimental.
	Merge *KodiakMergeOptions `field:"optional" json:"merge" yaml:"merge"`
	// Branch update settings (`[update]`).
	// Default: - default Kodiak update behavior.
	//
	// Experimental.
	Update *KodiakUpdateOptions `field:"optional" json:"update" yaml:"update"`
	// Kodiak configuration format version.
	// Default: 1.
	//
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

