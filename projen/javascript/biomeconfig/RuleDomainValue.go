package biomeconfig


// Experimental.
type RuleDomainValue string

const (
	// Enables all the rules that belong to this domain (all).
	// Experimental.
	RuleDomainValue_ALL RuleDomainValue = "ALL"
	// Disables all the rules that belong to this domain (none).
	// Experimental.
	RuleDomainValue_NONE RuleDomainValue = "NONE"
	// Enables only the recommended rules for this domain (recommended).
	// Experimental.
	RuleDomainValue_RECOMMENDED RuleDomainValue = "RECOMMENDED"
)

