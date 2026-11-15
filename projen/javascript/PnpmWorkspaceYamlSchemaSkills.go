package javascript


// Agent skills that approved dependencies ship.
//
// Added in pnpm v12.11.0.
// Experimental.
type PnpmWorkspaceYamlSchemaSkills struct {
	// The directories, relative to the workspace root, that pnpm links approved agent skills into.
	//
	// When set, it replaces the detected directories, and pnpm creates the ones that are missing. An empty list turns agent skills off.
	// Experimental.
	Dirs *[]*string `field:"optional" json:"dirs" yaml:"dirs"`
}

