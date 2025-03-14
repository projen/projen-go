package workflows


// A job matrix.
// Experimental.
type JobMatrix struct {
	// Each option you define in the matrix has a key and value.
	//
	// The keys you
	// define become properties in the matrix context and you can reference the
	// property in other areas of your workflow file. For example, if you define
	// the key os that contains an array of operating systems, you can use the
	// matrix.os property as the value of the runs-on keyword to create a job
	// for each operating system.
	// Experimental.
	Domain *map[string]interface{} `field:"optional" json:"domain" yaml:"domain"`
	// You can remove a specific configurations defined in the build matrix using the exclude option.
	//
	// Using exclude removes a job defined by the
	// build matrix.
	// Experimental.
	Exclude *[]*map[string]interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// You can add additional configuration options to a build matrix job that already exists.
	//
	// For example, if you want to use a specific version of npm
	// when the job that uses windows-latest and version 8 of node runs, you can
	// use include to specify that additional option.
	// Experimental.
	Include *[]*map[string]interface{} `field:"optional" json:"include" yaml:"include"`
}

