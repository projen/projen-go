package javascript


// The node package manager to use.
// Experimental.
type NodePackageManager string

const (
	// Use `yarn` as the package manager.
	// Deprecated: For `yarn` 1.x use `YARN_CLASSIC` for `yarn` >= 2 use `YARN_BERRY`. Currently, `NodePackageManager.YARN` means `YARN_CLASSIC`. In the future, we might repurpose it to mean `YARN_BERRY`.
	NodePackageManager_YARN NodePackageManager = "YARN"
	// Use `yarn` versions >= 2 as the package manager.
	// Deprecated: use YARN_BERRY instead.
	NodePackageManager_YARN2 NodePackageManager = "YARN2"
	// Use `yarn` 1.x as the package manager.
	// Experimental.
	NodePackageManager_YARN_CLASSIC NodePackageManager = "YARN_CLASSIC"
	// Use `yarn` versions >= 2 as the package manager.
	// Experimental.
	NodePackageManager_YARN_BERRY NodePackageManager = "YARN_BERRY"
	// Use `npm` as the package manager.
	// Experimental.
	NodePackageManager_NPM NodePackageManager = "NPM"
	// Use `pnpm` as the package manager.
	// Experimental.
	NodePackageManager_PNPM NodePackageManager = "PNPM"
	// Use `bun` as the package manager.
	// Experimental.
	NodePackageManager_BUN NodePackageManager = "BUN"
)

