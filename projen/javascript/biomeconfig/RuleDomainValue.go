package biomeconfig


// Experimental.
type RuleDomainValue string

const (
	// Enables all rules in the domain.
	//
	// (all).
	// Experimental.
	RuleDomainValue_ALL RuleDomainValue = "ALL"
	// Disables the domain.
	//
	// (none).
	// Experimental.
	RuleDomainValue_NONE RuleDomainValue = "NONE"
	// Enables the domain's recommended non-nursery rules.
	//
	// (recommended).
	// Experimental.
	RuleDomainValue_RECOMMENDED RuleDomainValue = "RECOMMENDED"
)

