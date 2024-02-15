package github


// Experimental.
type DownloadArtifactWith struct {
	// When multiple artifacts are matched, this changes the behavior of the destination directories If true, the downloaded artifacts will be in the same directory specified by path If false, the downloaded artifacts will be extracted into individual named directories within the specified path.
	// Default: false.
	//
	// Experimental.
	MergeMultiple *bool `field:"optional" json:"mergeMultiple" yaml:"mergeMultiple"`
	// Name of the artifact to download.
	// Default: - If unspecified, all artifacts for the run are downloaded.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A file, directory or wildcard pattern that describes what to download.
	//
	// Supports basic tilde expansion.
	// Default: - $GITHUB_WORKSPACE.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// A glob pattern to the artifacts that should be downloaded This is ignored if name is specified.
	// Experimental.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The repository owner and the repository name joined together by "/" If github-token is specified, this is the repository that artifacts will be downloaded from.
	// Default: - ${{ github.repository }}
	//
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The id of the workflow run where the desired download artifact was uploaded from If github-token is specified, this is the run that artifacts will be downloaded from.
	// Default: - ${{ github.run_id }}
	//
	// Experimental.
	RunId *string `field:"optional" json:"runId" yaml:"runId"`
	// The GitHub token used to authenticate with the GitHub API to download artifacts from a different repository or from a different workflow run.
	// Default: - If unspecified, the action will download artifacts from the current repo and the current workflow run.
	//
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

