package uvconfig


// The oldest required release of each Linux libc.
// Experimental.
type MinimumLibcVersion struct {
	// Experimental.
	Glibc *string `field:"optional" json:"glibc" yaml:"glibc"`
	// Experimental.
	Musl *string `field:"optional" json:"musl" yaml:"musl"`
}

