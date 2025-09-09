package biomeconfig


// A list of rules that belong to this group.
// Experimental.
type Source struct {
	// Provides a code action to sort the imports and exports in the file using a built-in or custom order.
	// Experimental.
	OrganizeImports interface{} `field:"optional" json:"organizeImports" yaml:"organizeImports"`
	// Enables the recommended rules for this group.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Enforce attribute sorting in JSX elements.
	// Experimental.
	UseSortedAttributes interface{} `field:"optional" json:"useSortedAttributes" yaml:"useSortedAttributes"`
	// Sort the keys of a JSON object in natural order.
	// Experimental.
	UseSortedKeys interface{} `field:"optional" json:"useSortedKeys" yaml:"useSortedKeys"`
	// Enforce ordering of CSS properties and nested rules.
	// Experimental.
	UseSortedProperties interface{} `field:"optional" json:"useSortedProperties" yaml:"useSortedProperties"`
}

