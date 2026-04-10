package projen


// Behaviour when a release timestamp is missing for `minimumReleaseAge`.
// See: https://docs.renovatebot.com/configuration-options/#minimumreleaseagebehaviour
//
// Experimental.
type RenovatebotMinimumReleaseAgeBehaviour string

const (
	// A release without a timestamp is not treated as stable.
	// Experimental.
	RenovatebotMinimumReleaseAgeBehaviour_TIMESTAMP_REQUIRED RenovatebotMinimumReleaseAgeBehaviour = "TIMESTAMP_REQUIRED"
	// A release without a timestamp is treated as stable.
	// Experimental.
	RenovatebotMinimumReleaseAgeBehaviour_TIMESTAMP_OPTIONAL RenovatebotMinimumReleaseAgeBehaviour = "TIMESTAMP_OPTIONAL"
)

