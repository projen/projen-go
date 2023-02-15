package release


// Experimental.
type ScheduledReleaseOptions struct {
	// Cron schedule for releases.
	//
	// Only defined if this is a scheduled release.
	//
	// Example:
	//   '0 17 * * *' - every day at 5 pm
	//
	// Experimental.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
}

