package awscdk


// Options for `CdkJson`.
// Experimental.
type CdkConfigOptions struct {
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Default: - no build command.
	//
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Default: "cdk.out"
	//
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Default: - no additional context.
	//
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Default: true.
	//
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Default: ApprovalLevel.BROADENING
	//
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Default: [].
	//
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Default: [].
	//
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
	// The command line to execute in order to synthesize the CDK application (language specific).
	// Experimental.
	App *string `field:"required" json:"app" yaml:"app"`
}

