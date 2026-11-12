package github


// Options for Kodiak `[merge]`.
// Experimental.
type KodiakMergeOptions struct {
	// Configure automatic merging for dependency update bots such as Dependabot or Renovate.
	// Default: - dependency auto-merge is not configured.
	//
	// Experimental.
	AutomergeDependencies *KodiakAutomergeDependencies `field:"optional" json:"automergeDependencies" yaml:"automergeDependencies"`
	// Label required on a pull request to trigger auto-merge.
	// Default: "automerge".
	//
	// Experimental.
	AutomergeLabel *string `field:"optional" json:"automergeLabel" yaml:"automergeLabel"`
	// Automatically delete the head branch after merging.
	// Default: false.
	//
	// Experimental.
	DeleteBranchOnMerge *bool `field:"optional" json:"deleteBranchOnMerge" yaml:"deleteBranchOnMerge"`
	// Prevent Kodiak from merging pull requests while still allowing updates.
	// Default: false.
	//
	// Experimental.
	DoNotMerge *bool `field:"optional" json:"doNotMerge" yaml:"doNotMerge"`
	// Merge strategy used when merging pull requests.
	// Default: - the first merge method enabled on the repository (merge, then squash, then rebase).
	//
	// Experimental.
	Method MergeMethod `field:"optional" json:"method" yaml:"method"`
	// Prioritize pull requests that are already up-to-date and passing checks.
	// Default: false.
	//
	// Experimental.
	PrioritizeReadyToMerge *bool `field:"optional" json:"prioritizeReadyToMerge" yaml:"prioritizeReadyToMerge"`
	// Whether pull requests must have `automergeLabel` to be merged by Kodiak.
	// Default: true.
	//
	// Experimental.
	RequireAutomergeLabel *bool `field:"optional" json:"requireAutomergeLabel" yaml:"requireAutomergeLabel"`
}

