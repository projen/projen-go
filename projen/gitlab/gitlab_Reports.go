package gitlab


// Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests.
// See: https://docs.gitlab.com/ee/ci/yaml/#artifactsreports
//
// Experimental.
type Reports struct {
	// Path for file(s) that should be parsed as Cobertura XML coverage report.
	// Experimental.
	Cobertura *[]*string `field:"optional" json:"cobertura" yaml:"cobertura"`
	// Path to file or list of files with code quality report(s) (such as Code Climate).
	// Experimental.
	Codequality *[]*string `field:"optional" json:"codequality" yaml:"codequality"`
	// Path to file or list of files with Container scanning vulnerabilities report(s).
	// Experimental.
	ContainerScanning *[]*string `field:"optional" json:"containerScanning" yaml:"containerScanning"`
	// Path to file or list of files with DAST vulnerabilities report(s).
	// Experimental.
	Dast *[]*string `field:"optional" json:"dast" yaml:"dast"`
	// Path to file or list of files with Dependency scanning vulnerabilities report(s).
	// Experimental.
	DependencyScanning *[]*string `field:"optional" json:"dependencyScanning" yaml:"dependencyScanning"`
	// Path to file or list of files containing runtime-created variables for this job.
	// Experimental.
	Dotenv *[]*string `field:"optional" json:"dotenv" yaml:"dotenv"`
	// Path for file(s) that should be parsed as JUnit XML result.
	// Experimental.
	Junit *[]*string `field:"optional" json:"junit" yaml:"junit"`
	// Deprecated in 12.8: Path to file or list of files with license report(s).
	// Experimental.
	LicenseManagement *[]*string `field:"optional" json:"licenseManagement" yaml:"licenseManagement"`
	// Path to file or list of files with license report(s).
	// Experimental.
	LicenseScanning *[]*string `field:"optional" json:"licenseScanning" yaml:"licenseScanning"`
	// Path to file or list of files containing code intelligence (Language Server Index Format).
	// Experimental.
	Lsif *[]*string `field:"optional" json:"lsif" yaml:"lsif"`
	// Path to file or list of files with custom metrics report(s).
	// Experimental.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Path to file or list of files with performance metrics report(s).
	// Experimental.
	Performance *[]*string `field:"optional" json:"performance" yaml:"performance"`
	// Path to file or list of files with requirements report(s).
	// Experimental.
	Requirements *[]*string `field:"optional" json:"requirements" yaml:"requirements"`
	// Path to file or list of files with SAST vulnerabilities report(s).
	// Experimental.
	Sast *[]*string `field:"optional" json:"sast" yaml:"sast"`
	// Path to file or list of files with secret detection report(s).
	// Experimental.
	SecretDetection *[]*string `field:"optional" json:"secretDetection" yaml:"secretDetection"`
	// Path to file or list of files with terraform plan(s).
	// Experimental.
	Terraform *[]*string `field:"optional" json:"terraform" yaml:"terraform"`
}

