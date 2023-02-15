package workflows


// CRON schedule options.
// Experimental.
type CronScheduleOptions struct {
	// See: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07
	//
	// Experimental.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

