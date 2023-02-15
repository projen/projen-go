package gitlab


// Used to run a job multiple times in parallel in a single pipeline.
// Experimental.
type Parallel struct {
	// Defines different variables for jobs that are running in parallel.
	// Experimental.
	Matrix *[]*map[string]*[]interface{} `field:"required" json:"matrix" yaml:"matrix"`
}

