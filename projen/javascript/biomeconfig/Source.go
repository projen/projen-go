package biomeconfig


// A list of rules that belong to this group.
// Experimental.
type Source struct {
	// Remove duplicate CSS classes.
	//
	// See https://biomejs.dev/assist/actions/no-duplicate-classes
	// Experimental.
	NoDuplicateClasses interface{} `field:"optional" json:"noDuplicateClasses" yaml:"noDuplicateClasses"`
	// Sorts imports and exports in your JavaScript and TypeScript files.
	//
	// See https://biomejs.dev/assist/actions/organize-imports
	// Experimental.
	OrganizeImports interface{} `field:"optional" json:"organizeImports" yaml:"organizeImports"`
	// Enables the recommended rules for this group.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Enforce attribute sorting in JSX elements.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-attributes
	// Experimental.
	UseSortedAttributes interface{} `field:"optional" json:"useSortedAttributes" yaml:"useSortedAttributes"`
	// Sort interface members by key.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-interface-members
	// Experimental.
	UseSortedInterfaceMembers interface{} `field:"optional" json:"useSortedInterfaceMembers" yaml:"useSortedInterfaceMembers"`
	// Sort the keys of a JSON object in natural order.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-keys
	// Experimental.
	UseSortedKeys interface{} `field:"optional" json:"useSortedKeys" yaml:"useSortedKeys"`
	// Enforce ordering of CSS properties and nested rules.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-properties
	// Experimental.
	UseSortedProperties interface{} `field:"optional" json:"useSortedProperties" yaml:"useSortedProperties"`
}

