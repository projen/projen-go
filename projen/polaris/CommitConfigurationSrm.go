package polaris


// Software Risk Manager configuration to use when storing defects in Software Risk Manager.
// Experimental.
type CommitConfigurationSrm struct {
	// The URL of the Software Risk Manager to use for the analysis (if doing a remote analysis) and the analysis results.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The name of the branch to associate the analysis results with in Software Risk Manager.
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The name of the parent branch of the actual branch.
	// Experimental.
	ParentBranch *string `field:"optional" json:"parentBranch" yaml:"parentBranch"`
	// The ID of the project to associate the analysis results with in Software Risk Manager.
	// Experimental.
	ProjectId *float64 `field:"optional" json:"projectId" yaml:"projectId"`
	// The name of the project to associate the analysis results with in Software Risk Manager.
	// Experimental.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// The name of the file to read the Software Risk Manager API key from.
	//
	// By default, the file located at $HOME/.bridge/srm-token.txt is used.
	// Experimental.
	TokenFile *string `field:"optional" json:"tokenFile" yaml:"tokenFile"`
}

