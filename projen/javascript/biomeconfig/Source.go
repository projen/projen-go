package biomeconfig


// Configures one group of assist actions.
// Experimental.
type Source struct {
	// Remove repeated CSS class names from HTML class attributes.
	//
	// See https://biomejs.dev/assist/actions/no-duplicate-classes
	// Experimental.
	NoDuplicateClasses interface{} `field:"optional" json:"noDuplicateClasses" yaml:"noDuplicateClasses"`
	// Organize imports and exports in JavaScript and TypeScript files.
	//
	// See https://biomejs.dev/assist/actions/organize-imports
	// Experimental.
	OrganizeImports interface{} `field:"optional" json:"organizeImports" yaml:"organizeImports"`
	// Selects the baseline set of assist actions for this group.
	//
	// `recommended` enables
	// the group's recommended actions, `all` enables all actions in the group, and
	// `none` starts with no actions enabled in the group. Explicit action settings
	// override this group preset, which overrides `assist.actions.preset` for the group.
	// Experimental.
	Preset PresetConfig `field:"optional" json:"preset" yaml:"preset"`
	// Enables or disables the recommended actions in this group.
	//
	// The group's `preset`
	// setting takes precedence when both options are set.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Sort the arguments of GraphQL fields and directives in natural order.
	//
	// See https://biomejs.dev/assist/actions/use-graphql-sorted-arguments
	// Experimental.
	UseGraphqlSortedArguments interface{} `field:"optional" json:"useGraphqlSortedArguments" yaml:"useGraphqlSortedArguments"`
	// Sort the named definitions in a GraphQL document.
	//
	// See https://biomejs.dev/assist/actions/use-graphql-sorted-definitions
	// Experimental.
	UseGraphqlSortedDefinitions interface{} `field:"optional" json:"useGraphqlSortedDefinitions" yaml:"useGraphqlSortedDefinitions"`
	// Sort the variable definitions of GraphQL operations in natural order.
	//
	// See https://biomejs.dev/assist/actions/use-graphql-sorted-variables
	// Experimental.
	UseGraphqlSortedVariables interface{} `field:"optional" json:"useGraphqlSortedVariables" yaml:"useGraphqlSortedVariables"`
	// Sort HTML attributes and framework directives into a consistent order.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-attributes
	// Experimental.
	UseSortedAttributes interface{} `field:"optional" json:"useSortedAttributes" yaml:"useSortedAttributes"`
	// Sort GraphQL enum values in natural order.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-enum-members
	// Experimental.
	UseSortedEnumMembers interface{} `field:"optional" json:"useSortedEnumMembers" yaml:"useSortedEnumMembers"`
	// Sort named members of a TypeScript interface.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-interface-members
	// Experimental.
	UseSortedInterfaceMembers interface{} `field:"optional" json:"useSortedInterfaceMembers" yaml:"useSortedInterfaceMembers"`
	// Sort keys in a JSON object.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-keys
	// Experimental.
	UseSortedKeys interface{} `field:"optional" json:"useSortedKeys" yaml:"useSortedKeys"`
	// Organize package.json fields according to established conventions. See https://biomejs.dev/assist/actions/use-sorted-package-json.
	// Experimental.
	UseSortedPackageJson interface{} `field:"optional" json:"useSortedPackageJson" yaml:"useSortedPackageJson"`
	// Sort CSS properties and nested rules into a consistent order.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-properties
	// Experimental.
	UseSortedProperties interface{} `field:"optional" json:"useSortedProperties" yaml:"useSortedProperties"`
	// Sort selections inside GraphQL operations and fragments.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-selection-set
	// Experimental.
	UseSortedSelectionSet interface{} `field:"optional" json:"useSortedSelectionSet" yaml:"useSortedSelectionSet"`
	// Sort fields in GraphQL type definitions by name.
	//
	// See https://biomejs.dev/assist/actions/use-sorted-type-fields
	// Experimental.
	UseSortedTypeFields interface{} `field:"optional" json:"useSortedTypeFields" yaml:"useSortedTypeFields"`
}

