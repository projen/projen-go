package uvconfig


// Experimental.
type IndexFormat string

const (
	// A PyPI-style index implementing the Simple Repository API.
	//
	// (simple).
	// Experimental.
	IndexFormat_SIMPLE IndexFormat = "SIMPLE"
	// A `--find-links`-style index containing a flat list of wheels and source distributions.
	//
	// (flat).
	// Experimental.
	IndexFormat_FLAT IndexFormat = "FLAT"
)

