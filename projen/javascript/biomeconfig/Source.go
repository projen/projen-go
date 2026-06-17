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
	// Enables a particular rule preset.
	// Experimental.
	Preset PresetConfig `field:"optional" json:"preset" yaml:"preset"`
	// Enables the recommended rules for this group.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Enforce attribute sorting in HTML elements.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-attributes
	// Experimental.
	UseSortedAttributes interface{} `field:"optional" json:"useSortedAttributes" yaml:"useSortedAttributes"`
	// Sort the members of an enum in natural order.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-enum-members
	// Experimental.
	UseSortedEnumMembers interface{} `field:"optional" json:"useSortedEnumMembers" yaml:"useSortedEnumMembers"`
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
	// Organize package.json fields according to established conventions. See https://biomejs.dev/assist/actions/use-sorted-package-json.
	// Experimental.
	UseSortedPackageJson interface{} `field:"optional" json:"useSortedPackageJson" yaml:"useSortedPackageJson"`
	// Enforce ordering of CSS properties and nested rules.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-properties
	// Experimental.
	UseSortedProperties interface{} `field:"optional" json:"useSortedProperties" yaml:"useSortedProperties"`
	// Sort GraphQL selection sets.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-selection-set
	// Experimental.
	UseSortedSelectionSet interface{} `field:"optional" json:"useSortedSelectionSet" yaml:"useSortedSelectionSet"`
	// Sort fields in GraphQL type definitions alphabetically.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-type-fields
	// Experimental.
	UseSortedTypeFields interface{} `field:"optional" json:"useSortedTypeFields" yaml:"useSortedTypeFields"`
}

