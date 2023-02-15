package javascript

import (
	"github.com/projen/projen-go/projen"
)

// Experimental.
type Bundle struct {
	// The task that produces this bundle.
	// Experimental.
	BundleTask projen.Task `field:"required" json:"bundleTask" yaml:"bundleTask"`
	// Base directory containing the output file (relative to project root).
	// Experimental.
	Outdir *string `field:"required" json:"outdir" yaml:"outdir"`
	// Location of the output file (relative to project root).
	// Experimental.
	Outfile *string `field:"required" json:"outfile" yaml:"outfile"`
	// The "watch" task for this bundle.
	// Experimental.
	WatchTask projen.Task `field:"optional" json:"watchTask" yaml:"watchTask"`
}

