package javascript


// Versioning settings for pnpm's native workspace release management, used by `pnpm change` and recursive `pnpm version`.
// Experimental.
type PnpmWorkspaceYamlSchemaVersioning struct {
	// Controls where release changelog content is stored.
	// Experimental.
	Changelog *PnpmWorkspaceYamlSchemaVersioningChangelog `field:"optional" json:"changelog" yaml:"changelog"`
	// Ties member projects to a lead project and constrains their major versions to the lead's major-version band.
	// Experimental.
	Epics *[]*PnpmWorkspaceYamlSchemaVersioningEpics `field:"optional" json:"epics" yaml:"epics"`
	// Groups of workspace projects that always release together at one shared version.
	//
	// The shared version is the highest current version in the group, bumped by the largest bump any member needs.
	// Experimental.
	Fixed *[]*[]*string `field:"optional" json:"fixed" yaml:"fixed"`
	// Workspace projects permanently excluded from versioning and dependent propagation.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// Maps a workspace project to a release lane.
	//
	// Unlisted projects are on the reserved `main` lane and release stable versions.
	// Experimental.
	Lanes *map[string]*string `field:"optional" json:"lanes" yaml:"lanes"`
	// Caps the bump that a release from the current checkout may apply, after dependent propagation and fixed-group resolution.
	// Experimental.
	MaxBump PnpmWorkspaceYamlSchemaVersioningMaxBump `field:"optional" json:"maxBump" yaml:"maxBump"`
}

