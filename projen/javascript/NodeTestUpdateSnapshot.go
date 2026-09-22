package javascript


// Whether to update snapshots in task "test" (which is executed in task "build" and build workflows), or create a separate task "test:update" for updating snapshots.
// Experimental.
type NodeTestUpdateSnapshot string

const (
	// Always update snapshots in "test" task.
	// Experimental.
	NodeTestUpdateSnapshot_ALWAYS NodeTestUpdateSnapshot = "ALWAYS"
	// Never update snapshots in "test" task and create a separate "test:update" task.
	// Experimental.
	NodeTestUpdateSnapshot_NEVER NodeTestUpdateSnapshot = "NEVER"
)

