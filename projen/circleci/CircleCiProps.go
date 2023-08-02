package circleci


// Options for class {@link Circleci}.
// See: https://circleci.com/docs/2.0/configuration-reference/
//
// Experimental.
type CircleCiProps struct {
	// List of Jobs to create unique steps per pipeline, e.g. ```json jobs: [{  identifier: "compile",  docker: { image: "golang:alpine" }  steps: ["checkout", run: {command: "go build ."}] }] ```.
	// See: https://circleci.com/docs/2.0/configuration-reference/#jobs
	//
	// Experimental.
	Jobs *[]*Job `field:"optional" json:"jobs" yaml:"jobs"`
	// Contains a map of CirclCi Orbs ```json orbs: {  node: "circleci/node@5.0.1"  slack: "circleci/slack@4.8.3" } ```.
	// Experimental.
	Orbs *map[string]*string `field:"optional" json:"orbs" yaml:"orbs"`
	// The setup field enables you to conditionally trigger configurations from outside the primary .circleci parent directory, update pipeline parameters, or generate customized configurations.
	// See: https://circleci.com/docs/2.0/configuration-reference/#setup
	//
	// Experimental.
	Setup *bool `field:"optional" json:"setup" yaml:"setup"`
	// pipeline version.
	// See: https://circleci.com/docs/2.0/configuration-reference/#version
	//
	// Default: 2.1
	//
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
	// List of Workflows of pipeline, e.g. ```json workflows: {   {     identifier: "build",       jobs: [{          identifier: "node/install",          context: ["npm"],       }]   } } ```.
	// See: https://circleci.com/docs/2.0/configuration-reference/#workflows
	//
	// Experimental.
	Workflows *[]*Workflow `field:"optional" json:"workflows" yaml:"workflows"`
}

