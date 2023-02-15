package awscdk


// Common options for `cdk.json`.
// Experimental.
type CdkConfigCommonOptions struct {
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
}

