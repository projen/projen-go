package uvconfig


// The method to use when linking.
//
// Defaults to [`LinkMode::Clone`] on macOS and Linux (which support copy-on-write on
// APFS and btrfs/xfs/bcachefs respectively), and [`LinkMode::Hardlink`] on other
// platforms.
// Default: LinkMode::Clone`] on macOS and Linux (which support copy-on-write on.
//
// Experimental.
type LinkMode string

const (
	// Clone (i.e., copy-on-write) packages from the source into the destination. (clone).
	// Experimental.
	LinkMode_CLONE LinkMode = "CLONE"
	// Copy packages from the source into the destination.
	//
	// (copy).
	// Experimental.
	LinkMode_COPY LinkMode = "COPY"
	// Hard link packages from the source into the destination.
	//
	// (hardlink).
	// Experimental.
	LinkMode_HARDLINK LinkMode = "HARDLINK"
	// Symbolically link packages from the source into the destination.
	//
	// (symlink).
	// Experimental.
	LinkMode_SYMLINK LinkMode = "SYMLINK"
)

