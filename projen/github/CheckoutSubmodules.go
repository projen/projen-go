package github


// Whether to checkout Git submodules in CI workflows.
// Experimental.
type CheckoutSubmodules string

const (
	// Don't checkout submodules.
	// Experimental.
	CheckoutSubmodules_DISABLED CheckoutSubmodules = "DISABLED"
	// Checkout only top-level submodules.
	// Experimental.
	CheckoutSubmodules_ENABLED CheckoutSubmodules = "ENABLED"
	// Checkout submodules recursively.
	// Experimental.
	CheckoutSubmodules_RECURSIVE CheckoutSubmodules = "RECURSIVE"
)

