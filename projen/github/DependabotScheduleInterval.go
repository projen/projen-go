package github


// How often to check for new versions and raise pull requests for version updates.
// Experimental.
type DependabotScheduleInterval string

const (
	// Runs on every weekday, Monday to Friday.
	// Experimental.
	DependabotScheduleInterval_DAILY DependabotScheduleInterval = "DAILY"
	// Runs once each week.
	//
	// By default, this is on Monday.
	// Experimental.
	DependabotScheduleInterval_WEEKLY DependabotScheduleInterval = "WEEKLY"
	// Runs once each month.
	//
	// This is on the first day of the month.
	// Experimental.
	DependabotScheduleInterval_MONTHLY DependabotScheduleInterval = "MONTHLY"
)

