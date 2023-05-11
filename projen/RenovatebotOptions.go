package projen


// Options for Renovatebot.
// Experimental.
type RenovatebotOptions struct {
	// You can use the `ignore` option to customize which dependencies are updated.
	//
	// The ignore option supports just package name.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// Ignores updates to `projen`.
	//
	// This is required since projen updates may cause changes in committed files
	// and anti-tamper checks will fail.
	//
	// Projen upgrades are covered through the `ProjenUpgrade` class.
	// Experimental.
	IgnoreProjen *bool `field:"optional" json:"ignoreProjen" yaml:"ignoreProjen"`
	// List of labels to apply to the created PR's.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Experimental.
	Marker *bool `field:"optional" json:"marker" yaml:"marker"`
	// Experimental.
	OverrideConfig interface{} `field:"optional" json:"overrideConfig" yaml:"overrideConfig"`
	// How often to check for new versions and raise pull requests.
	//
	// Can be given in CRON or LATER format, and use multiple schedules
	// (e.g. different for weekdays and weekends). Multiple rules are
	// handles as OR.
	//
	// Some normal scheduling values defined in enum `RenovatebotScheduleInterval`.
	// See: https://docs.renovatebot.com/configuration-options/#schedule
	//
	// Experimental.
	ScheduleInterval *[]*string `field:"optional" json:"scheduleInterval" yaml:"scheduleInterval"`
}

