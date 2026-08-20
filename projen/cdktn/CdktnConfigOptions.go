package cdktn


// Options for `CdktnConfig`.
// Experimental.
type CdktnConfigOptions struct {
	// CDKTN output directory.
	// Default: "cdktf.out"
	//
	// Experimental.
	CdktnOut *string `field:"optional" json:"cdktnOut" yaml:"cdktnOut"`
	// Additional context to include in `cdktf.json`.
	// Default: - no additional context.
	//
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// CDK project identifier.
	// Default: - Automatically generated.
	//
	// Experimental.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Whether report crashing to a remote server.
	// Default: true.
	//
	// Experimental.
	SendCrashReports *bool `field:"optional" json:"sendCrashReports" yaml:"sendCrashReports"`
	// Terraform modules to build.
	// Default: [].
	//
	// Experimental.
	TerraformModules *[]*string `field:"optional" json:"terraformModules" yaml:"terraformModules"`
	// Terraform providers to build.
	// Default: [].
	//
	// Experimental.
	TerraformProviders *[]*string `field:"optional" json:"terraformProviders" yaml:"terraformProviders"`
	// The command line to execute in order to synthesize the CDKTN application (language specific).
	// Experimental.
	App *string `field:"required" json:"app" yaml:"app"`
	// The programming language used in the CDKTN app.
	// Default: "typescript".
	//
	// Experimental.
	Language *string `field:"optional" json:"language" yaml:"language"`
}

