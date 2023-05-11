package projen


// Whether the port visibility should be private or public.
// Experimental.
type GitpodPortVisibility string

const (
	// Allows everyone with the port URL to access the port (default).
	// Experimental.
	GitpodPortVisibility_PUBLIC GitpodPortVisibility = "PUBLIC"
	// Only allows users with workspace access to access the port.
	// Experimental.
	GitpodPortVisibility_PRIVATE GitpodPortVisibility = "PRIVATE"
)

