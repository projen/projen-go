package javascript


// The node package manager to use.
// Experimental.
type NodePackageManager string

const (
	// Use `yarn` as the package manager.
	// Experimental.
	NodePackageManager_YARN NodePackageManager = "YARN"
	// Use `npm` as the package manager.
	// Experimental.
	NodePackageManager_NPM NodePackageManager = "NPM"
	// Use `pnpm` as the package manager.
	// Experimental.
	NodePackageManager_PNPM NodePackageManager = "PNPM"
)

