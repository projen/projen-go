package workflows


// CRON schedule options.
// Experimental.
type CronScheduleOptions struct {
	// See: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07
	//
	// Experimental.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
	// IANA time zone name used to evaluate the cron expression, e.g. `America/New_York`.
	// See: https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax
	//
	// Default: - UTC.
	//
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

