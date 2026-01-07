package uvconfig


// Experimental.
type LinkMode string

const (
	// Clone (i.e., copy-on-write) packages from the wheel into the `site-packages` directory. (clone).
	// Experimental.
	LinkMode_CLONE LinkMode = "CLONE"
	// Copy packages from the wheel into the `site-packages` directory.
	//
	// (copy).
	// Experimental.
	LinkMode_COPY LinkMode = "COPY"
	// Hard link packages from the wheel into the `site-packages` directory.
	//
	// (hardlink).
	// Experimental.
	LinkMode_HARDLINK LinkMode = "HARDLINK"
	// Symbolically link packages from the wheel into the `site-packages` directory.
	//
	// (symlink).
	// Experimental.
	LinkMode_SYMLINK LinkMode = "SYMLINK"
)

