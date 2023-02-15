package javascript


// Experimental.
type UpdateSnapshot string

const (
	// Always update snapshots in "test" task.
	// Experimental.
	UpdateSnapshot_ALWAYS UpdateSnapshot = "ALWAYS"
	// Never update snapshots in "test" task and create a separate "test:update" task.
	// Experimental.
	UpdateSnapshot_NEVER UpdateSnapshot = "NEVER"
)

