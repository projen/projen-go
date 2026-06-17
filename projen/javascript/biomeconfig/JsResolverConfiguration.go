package biomeconfig


// Resolver options specific to JavaScript files.
// Experimental.
type JsResolverConfiguration struct {
	// Enables pnpm workspace catalog resolution for JavaScript package manifests.
	//
	// Opt-in:
	// - Set `javascript.resolver.experimentalPnpmCatalogs` to `true`.
	//
	// Scope:
	// - Resolves `catalog:` and `catalog:<name>` dependency versions from
	// `package.json`.
	// - Applies to `dependencies`, `devDependencies`, and `peerDependencies`.
	//
	// Fail-safe behavior:
	// - If `pnpm-workspace.yaml` is missing, unreadable, or cannot be parsed,
	// Biome silently falls back to the default behavior (as if this option
	// were disabled).
	// - Unknown keys and unsupported value shapes in `pnpm-workspace.yaml` are
	// ignored.
	//
	// Limitations:
	// - Only `pnpm-workspace.yaml` is read.
	// - Biome only reads top-level `catalog` / `catalogs` mappings and scalar
	// string entries.
	//
	// Default: `false`.
	// Experimental.
	ExperimentalPnpmCatalogs *bool `field:"optional" json:"experimentalPnpmCatalogs" yaml:"experimentalPnpmCatalogs"`
}

