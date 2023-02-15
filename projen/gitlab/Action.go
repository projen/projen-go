package gitlab


// Specifies what this job will do.
//
// 'start' (default) indicates the job will start the
// deployment. 'prepare' indicates this will not affect the deployment. 'stop' indicates
// this will stop the deployment.
// Experimental.
type Action string

const (
	// Experimental.
	Action_PREPARE Action = "PREPARE"
	// Experimental.
	Action_START Action = "START"
	// Experimental.
	Action_STOP Action = "STOP"
)

