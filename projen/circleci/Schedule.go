package circleci


// A workflow may have a schedule indicating it runs at a certain time.
// See: https://circleci.com/docs/2.0/configuration-reference/#schedule
//
// Experimental.
type Schedule struct {
	// Experimental.
	Filters *Filter `field:"required" json:"filters" yaml:"filters"`
	// The cron key is defined using POSIX crontab syntax.
	// Experimental.
	Cron *string `field:"optional" json:"cron" yaml:"cron"`
}

