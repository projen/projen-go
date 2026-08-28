package polaris


// Experimental.
type CheckerConfiguration struct {
	// Indicates whether all checkers should be enabled.
	// Experimental.
	All *bool `field:"optional" json:"all" yaml:"all"`
	// Indicates whether all security checkers should be enabled.
	//
	// This includes the Security, Android Security, and Web App Security categories, and other security checkers that require explicit enablement.
	// Experimental.
	AllSecurity *bool `field:"optional" json:"allSecurity" yaml:"allSecurity"`
	// If set to true, enables android security checkers.
	// Experimental.
	AndroidSecurity *bool `field:"optional" json:"androidSecurity" yaml:"androidSecurity"`
	// Enables audit checkers.
	// Experimental.
	Audit *bool `field:"optional" json:"audit" yaml:"audit"`
	// Indicates whether the brakeman checkers should be enabled or disabled.
	// Experimental.
	Brakeman *bool `field:"optional" json:"brakeman" yaml:"brakeman"`
	// Enables C, C++, Objective-C, Objective-C++ security-related checkers that are disabled by default.
	// Experimental.
	CFamilySecurity *bool `field:"optional" json:"cFamilySecurity" yaml:"cFamilySecurity"`
	// Map from checker name to configuration for the checker.
	//
	// The configuration indicates whether the checker should be enabled or not and allows users to set options used to configure the checker.
	// Experimental.
	CheckerConfig interface{} `field:"optional" json:"checkerConfig" yaml:"checkerConfig"`
	// Specifies CodeXM (.cxm) files to use in the analysis.
	// Experimental.
	Codexm *[]*string `field:"optional" json:"codexm" yaml:"codexm"`
	// Enables C, C++ concurrency checkers that are disabled by default.
	// Experimental.
	Concurrency *bool `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Specifies whether to enable the default set of checkers.
	//
	// If set to true, the default set of checkers is enabled. Set to false to get more control over which checkers are enabled.
	// Experimental.
	Default *bool `field:"optional" json:"default" yaml:"default"`
	// Enables or disables PMD for Apex analysis.
	// Experimental.
	Pmd *bool `field:"optional" json:"pmd" yaml:"pmd"`
	// Enables or disables recommended security checkers.
	// Experimental.
	RecommendedSecurityCheckers *bool `field:"optional" json:"recommendedSecurityCheckers" yaml:"recommendedSecurityCheckers"`
	// Enables C, C++ rule checkers.
	// Experimental.
	Rule *bool `field:"optional" json:"rule" yaml:"rule"`
	// Specifies how web application security analysis should be done.
	// Experimental.
	WebappSecurity *CheckerConfigurationWebappSecurity `field:"optional" json:"webappSecurity" yaml:"webappSecurity"`
}

