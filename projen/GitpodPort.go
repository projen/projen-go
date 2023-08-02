package projen


// Options for an exposed port on Gitpod.
// Experimental.
type GitpodPort struct {
	// What to do when a service on a port is detected.
	// Default: GitpodOnOpen.NOTIFY
	//
	// Experimental.
	OnOpen GitpodOnOpen `field:"optional" json:"onOpen" yaml:"onOpen"`
	// A port that should be exposed (forwarded) from the container.
	//
	// Example:
	//   "8080"
	//
	// Experimental.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Whether the port visibility should be private or public.
	// Default: GitpodPortVisibility.PUBLIC
	//
	// Experimental.
	Visibility GitpodPortVisibility `field:"optional" json:"visibility" yaml:"visibility"`
}

