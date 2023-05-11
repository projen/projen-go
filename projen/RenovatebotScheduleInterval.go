package projen


// How often to check for new versions and raise pull requests for version updates.
// See: https://docs.renovatebot.com/presets-schedule/
//
// Experimental.
type RenovatebotScheduleInterval string

const (
	// Run at any time.
	// Experimental.
	RenovatebotScheduleInterval_ANY_TIME RenovatebotScheduleInterval = "ANY_TIME"
	// Weekly schedule on early monday mornings.
	// Experimental.
	RenovatebotScheduleInterval_EARLY_MONDAYS RenovatebotScheduleInterval = "EARLY_MONDAYS"
	// Schedule daily.
	// Experimental.
	RenovatebotScheduleInterval_DAILY RenovatebotScheduleInterval = "DAILY"
	// Schedule monthly.
	// Experimental.
	RenovatebotScheduleInterval_MONTHLY RenovatebotScheduleInterval = "MONTHLY"
	// Schedule quarterly.
	// Experimental.
	RenovatebotScheduleInterval_QUARTERLY RenovatebotScheduleInterval = "QUARTERLY"
	// Schedule for weekends.
	// Experimental.
	RenovatebotScheduleInterval_WEEKENDS RenovatebotScheduleInterval = "WEEKENDS"
	// Schedule for weekdays.
	// Experimental.
	RenovatebotScheduleInterval_WEEKDAYS RenovatebotScheduleInterval = "WEEKDAYS"
)

