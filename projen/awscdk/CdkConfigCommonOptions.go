package awscdk


// Common options for `cdk.json`.
// Experimental.
type CdkConfigCommonOptions struct {
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
	// Feature flags that should be enabled in `cdk.json`.
	//
	// Make sure to double-check any changes to feature flags in `cdk.json` before deploying.
	// Unexpected changes may cause breaking changes in your CDK app.
	// You can overwrite any feature flag by passing it into the context field.
	// Default: - no feature flags are enabled by default.
	//
	// Experimental.
	FeatureFlags ICdkFeatureFlags `field:"optional" json:"featureFlags" yaml:"featureFlags"`
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
}

