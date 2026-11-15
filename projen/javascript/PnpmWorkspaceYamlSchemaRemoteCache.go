package javascript


// The server that shares task results and dependency builds between machines, and the keys every shared artifact is signed and verified with.
//
// Read by both sideEffectsCache.remote and the task cache of pnpm pipeline. Only org may be set in pnpm-workspace.yaml; the fields naming the server and its trust material (url, team, trustedKeys, publish, privateKey, keyId, builderId, imageDigest, architectureBaseline, buildEnv) are read from the global configuration file or the PNPM_REMOTE_CACHE_* environment variables instead, and setting them here fails with ERR_PNPM_WORKSPACE_REMOTE_CACHE_TRUST. Added in pnpm v12.12.0.
// Experimental.
type PnpmWorkspaceYamlSchemaRemoteCache struct {
	// The organization that owns the artifacts.
	//
	// On a pnpr server, one of the organizations it declares under artifacts.orgs. It also names the organization pnpm pipeline --report records a run under.
	// Experimental.
	Org *string `field:"optional" json:"org" yaml:"org"`
}

