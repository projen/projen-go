package polaris


// Experimental.
type SnapshotConfiguration struct {
	// Date and time of snapshot to use for comparison report.
	//
	// The value should be of the form "YYYY-MM-DDThh:mm:ss" where date and time are separated by a "T", optionally followed by a time zone specification consisting of either "Z" denoting UTC or a "+" or "-" character followed by colon-separated hours and minutes east of UTC. Example: "2023-12-27T13:21:05-08:00". If no time zone is specified, the local time zone is assumed. This key is mutually exclusive with the "id" and "reference" keys.
	// Experimental.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// ID of snapshot to use for comparison report.
	//
	// This key is mutually exclusive with the "date" and "reference" keys.
	// Experimental.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// One of "idir", "latest", or "scm".
	//
	// "idir" will use the snapshot created closest to, but not after, the creation date of the intermediate directory. "latest" will use the snapshot with the latest code-version date in the specified stream. "scm" will query the SCM to determine the version that was most recently checked out or updated, and then use the closest snapshot. This key is mutually exclusive with the "date" and "id" keys.
	// Experimental.
	Reference interface{} `field:"optional" json:"reference" yaml:"reference"`
}

