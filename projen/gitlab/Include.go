package gitlab


// An included YAML file.
// See: https://docs.gitlab.com/ee/ci/yaml/#include
//
// Experimental.
type Include struct {
	// Files from another private project on the same GitLab instance.
	//
	// You can use `file` in combination with `project` only.
	// Experimental.
	File *[]*string `field:"optional" json:"file" yaml:"file"`
	// Relative path from local repository root (`/`) to the `yaml`/`yml` file template.
	//
	// The file must be on the same branch, and does not work across git submodules.
	// Experimental.
	Local *string `field:"optional" json:"local" yaml:"local"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project`.
	// Experimental.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Branch/Tag/Commit-hash for the target project.
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// URL to a `yaml`/`yml` template file using HTTP/HTTPS.
	// Experimental.
	Remote *string `field:"optional" json:"remote" yaml:"remote"`
	// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
	// Experimental.
	Rules *[]*IncludeRule `field:"optional" json:"rules" yaml:"rules"`
	// Use a `.gitlab-ci.yml` template as a base, e.g. `Nodejs.gitlab-ci.yml`.
	// Experimental.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

