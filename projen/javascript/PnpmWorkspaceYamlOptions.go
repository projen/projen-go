package javascript


// Options for `PnpmWorkspaceYaml`.
// See: https://pnpm.io/pnpm-workspace_yaml
//
// Experimental.
type PnpmWorkspaceYamlOptions struct {
	// A map of package matchers to explicitly allow (`true`) or disallow (`false`) script execution.
	//
	// This field replaces `onlyBuiltDependencies` and `ignoredBuiltDependencies` (which are also deprecated by this new setting), providing a single source of truth.
	// Experimental.
	AllowBuilds interface{} `field:"optional" json:"allowBuilds" yaml:"allowBuilds"`
	// A list of deprecated versions that the warnings are suppressed.
	// Experimental.
	AllowedDeprecatedVersions *map[string]*string `field:"optional" json:"allowedDeprecatedVersions" yaml:"allowedDeprecatedVersions"`
	// When true, installation won't fail if some of the patches from the "patchedDependencies" field were not applied.
	// Experimental.
	AllowNonAppliedPatches *bool `field:"optional" json:"allowNonAppliedPatches" yaml:"allowNonAppliedPatches"`
	// When true, installation won't fail if some of the patches from the "patchedDependencies" field were not applied.
	//
	// (Previously named "allowNonAppliedPatches").
	// Experimental.
	AllowUnusedPatches *bool `field:"optional" json:"allowUnusedPatches" yaml:"allowUnusedPatches"`
	// Experimental.
	AuditConfig *PnpmWorkspaceYamlSchemaAuditConfig `field:"optional" json:"auditConfig" yaml:"auditConfig"`
	// Controls the level of issues reported by `pnpm audit`.
	//
	// When set to 'low', all vulnerabilities are reported. When set to 'moderate', 'high', or 'critical', only vulnerabilities with that severity or higher are reported.
	// Experimental.
	AuditLevel PnpmWorkspaceYamlSchemaAuditLevel `field:"optional" json:"auditLevel" yaml:"auditLevel"`
	// When true, any missing non-optional peer dependencies are automatically installed.
	// Experimental.
	AutoInstallPeers *bool `field:"optional" json:"autoInstallPeers" yaml:"autoInstallPeers"`
	// When set to true, it prevents the resolution of exotic protocols (like git+ssh: or direct https: tarballs) in transitive dependencies.
	//
	// Only direct dependencies are allowed to use exotic sources.
	// Experimental.
	BlockExoticSubdeps *bool `field:"optional" json:"blockExoticSubdeps" yaml:"blockExoticSubdeps"`
	// The Certificate Authority signing certificate that is trusted for SSL connections to the registry.
	// Experimental.
	Ca *string `field:"optional" json:"ca" yaml:"ca"`
	// The location of the cache (package metadata and dlx).
	// Experimental.
	CacheDir *string `field:"optional" json:"cacheDir" yaml:"cacheDir"`
	// A path to a file containing one or multiple Certificate Authority signing certificates.
	// Experimental.
	Cafile *string `field:"optional" json:"cafile" yaml:"cafile"`
	// Define dependency version ranges as reusable constants, for later reference in package.json files. This (singular) field creates a catalog named default.
	// Experimental.
	Catalog *map[string]*string `field:"optional" json:"catalog" yaml:"catalog"`
	// Controlling if and how dependencies are added to the default catalog.
	// Experimental.
	CatalogMode PnpmWorkspaceYamlSchemaCatalogMode `field:"optional" json:"catalogMode" yaml:"catalogMode"`
	// Define arbitrarily named catalogs.
	// Experimental.
	Catalogs *map[string]*map[string]*string `field:"optional" json:"catalogs" yaml:"catalogs"`
	// A client certificate to pass when accessing the registry.
	// Experimental.
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// The maximum number of child processes to allocate simultaneously to build node_modules.
	// Experimental.
	ChildConcurrency *float64 `field:"optional" json:"childConcurrency" yaml:"childConcurrency"`
	// When set to `true`, pnpm will remove unused catalog entries during installation.
	// Experimental.
	CleanupUnusedCatalogs *bool `field:"optional" json:"cleanupUnusedCatalogs" yaml:"cleanupUnusedCatalogs"`
	// Controls colors in the output.
	// Experimental.
	Color PnpmWorkspaceYamlSchemaColor `field:"optional" json:"color" yaml:"color"`
	// Config dependencies allow you to share and centralize configuration files, settings, and hooks across multiple projects.
	//
	// They are installed before all regular dependencies ('dependencies', 'devDependencies', 'optionalDependencies'), making them ideal for setting up custom hooks, patches, and catalog entries.
	// Experimental.
	ConfigDependencies interface{} `field:"optional" json:"configDependencies" yaml:"configDependencies"`
	// If set to true, all build scripts (e.g. preinstall, install, postinstall) from dependencies will run automatically, without requiring approval.
	// Experimental.
	DangerouslyAllowAllBuilds *bool `field:"optional" json:"dangerouslyAllowAllBuilds" yaml:"dangerouslyAllowAllBuilds"`
	// When set to true, dependencies that are already symlinked to the root node_modules directory of the workspace will not be symlinked to subproject node_modules directories.
	// Experimental.
	DedupeDirectDeps *bool `field:"optional" json:"dedupeDirectDeps" yaml:"dedupeDirectDeps"`
	// When this setting is enabled, dependencies that are injected will be symlinked from the workspace whenever possible.
	// Experimental.
	DedupeInjectedDeps *bool `field:"optional" json:"dedupeInjectedDeps" yaml:"dedupeInjectedDeps"`
	// When this setting is set to true, packages with peer dependencies will be deduplicated after peers resolution.
	// Experimental.
	DedupePeerDependents *bool `field:"optional" json:"dedupePeerDependents" yaml:"dedupePeerDependents"`
	// When enabled, peer dependency suffixes use version-only identifiers (`name@version`) instead of full dep paths, eliminating nested suffixes like `(foo@1.0.0(bar@2.0.0))`. This dramatically reduces the number of package instances in projects with many recursive peer dependencies.
	// Experimental.
	DedupePeers *bool `field:"optional" json:"dedupePeers" yaml:"dedupePeers"`
	// When deploying a package or installing a local package, all files of the package are copied.
	// Experimental.
	DeployAllFiles *bool `field:"optional" json:"deployAllFiles" yaml:"deployAllFiles"`
	// When set to true, installation will fail if the workspace has cycles.
	// Experimental.
	DisallowWorkspaceCycles *bool `field:"optional" json:"disallowWorkspaceCycles" yaml:"disallowWorkspaceCycles"`
	// The time in minutes after which dlx cache expires.
	// Experimental.
	DlxCacheMaxAge *float64 `field:"optional" json:"dlxCacheMaxAge" yaml:"dlxCacheMaxAge"`
	// UNDOCUMENTED.
	//
	// When `true`, `pnpm publish` writes the README file's content into the published package.json (the `readme` field), so registries such as npmjs.com render the package's README. Added in pnpm 6.28.0; pnpm does not embed the README unless this is enabled. It also won't override a `readme` field already set in the package.json
	// Experimental.
	EmbedReadme *bool `field:"optional" json:"embedReadme" yaml:"embedReadme"`
	// When enabled, node_modules contains only symlinks to a central virtual store, rather than to node_modules/.pnpm.
	// Experimental.
	EnableGlobalVirtualStore *bool `field:"optional" json:"enableGlobalVirtualStore" yaml:"enableGlobalVirtualStore"`
	// When false, pnpm will not write any files to the modules directory (node_modules).
	// Experimental.
	EnableModulesDir *bool `field:"optional" json:"enableModulesDir" yaml:"enableModulesDir"`
	// When true, pnpm will run any pre/post scripts automatically.
	// Experimental.
	EnablePrePostScripts *bool `field:"optional" json:"enablePrePostScripts" yaml:"enablePrePostScripts"`
	// If this is enabled, pnpm will not install any package that claims to not be compatible with the current Node version.
	// Experimental.
	EngineStrict *bool `field:"optional" json:"engineStrict" yaml:"engineStrict"`
	// Experimental.
	ExecutionEnv *PnpmWorkspaceYamlSchemaExecutionEnv `field:"optional" json:"executionEnv" yaml:"executionEnv"`
	// When false, the NODE_PATH environment variable is not set in the command shims.
	// Experimental.
	ExtendNodePath *bool `field:"optional" json:"extendNodePath" yaml:"extendNodePath"`
	// If true, pnpm will fail if no packages match the filter.
	// Experimental.
	FailIfNoMatch *bool `field:"optional" json:"failIfNoMatch" yaml:"failIfNoMatch"`
	// How many times to retry if pnpm fails to fetch from the registry.
	// Experimental.
	FetchRetries *float64 `field:"optional" json:"fetchRetries" yaml:"fetchRetries"`
	// The exponential factor for retry backoff.
	// Experimental.
	FetchRetryFactor *float64 `field:"optional" json:"fetchRetryFactor" yaml:"fetchRetryFactor"`
	// The maximum fallback timeout to ensure the retry factor does not make requests too long.
	// Experimental.
	FetchRetryMaxtimeout *float64 `field:"optional" json:"fetchRetryMaxtimeout" yaml:"fetchRetryMaxtimeout"`
	// The minimum (base) timeout for retrying requests.
	// Experimental.
	FetchRetryMintimeout *float64 `field:"optional" json:"fetchRetryMintimeout" yaml:"fetchRetryMintimeout"`
	// The maximum amount of time to wait for HTTP requests to complete.
	// Experimental.
	FetchTimeout *float64 `field:"optional" json:"fetchTimeout" yaml:"fetchTimeout"`
	// By default, pnpm deploy will try creating a dedicated lockfile from a shared lockfile for deployment.
	//
	// If this setting is set to true, the legacy deploy behavior will be used.
	// Experimental.
	ForceLegacyDeploy *bool `field:"optional" json:"forceLegacyDeploy" yaml:"forceLegacyDeploy"`
	// When set to true, the generated lockfile name after installation will be named based on the current branch name to completely avoid merge conflicts.
	// Experimental.
	GitBranchLockfile *bool `field:"optional" json:"gitBranchLockfile" yaml:"gitBranchLockfile"`
	// Check if current branch is your publish branch, clean, and up-to-date with remote.
	// Experimental.
	GitChecks *bool `field:"optional" json:"gitChecks" yaml:"gitChecks"`
	// When fetching dependencies that are Git repositories, if the host is listed in this setting, pnpm will use shallow cloning to fetch only the needed commit, not all the history.
	// Experimental.
	GitShallowHosts *[]*string `field:"optional" json:"gitShallowHosts" yaml:"gitShallowHosts"`
	// Allows to set the target directory for the bin files of globally installed packages.
	// Experimental.
	GlobalBinDir *string `field:"optional" json:"globalBinDir" yaml:"globalBinDir"`
	// Specify a custom directory to store global packages.
	// Experimental.
	GlobalDir *string `field:"optional" json:"globalDir" yaml:"globalDir"`
	// The location of a global pnpmfile.
	//
	// A global pnpmfile is used by all projects during installation.
	// Experimental.
	GlobalPnpmfile *string `field:"optional" json:"globalPnpmfile" yaml:"globalPnpmfile"`
	// When true, all dependencies are hoisted to node_modules/.pnpm/node_modules.
	// Experimental.
	Hoist *bool `field:"optional" json:"hoist" yaml:"hoist"`
	// Added a new hoistingLimits setting for `nodeLinker: hoisted` installs, mirroring yarn's `nmHoistingLimits`.
	//
	// It accepts `none` (the default — hoist as far as possible), workspaces (hoist only as far as each workspace package), or dependencies (hoist only up to each workspace package's direct dependencies).
	// Experimental.
	HoistingLimits PnpmWorkspaceYamlSchemaHoistingLimits `field:"optional" json:"hoistingLimits" yaml:"hoistingLimits"`
	// Tells pnpm which packages should be hoisted to node_modules/.pnpm/node_modules.
	// Experimental.
	HoistPattern *[]*string `field:"optional" json:"hoistPattern" yaml:"hoistPattern"`
	// When true, packages from the workspaces are symlinked to either <workspace_root>/node_modules/.pnpm/node_modules or to <workspace_root>/node_modules depending on other hoisting settings (hoistPattern and publicHoistPattern).
	// Experimental.
	HoistWorkspacePackages *bool `field:"optional" json:"hoistWorkspacePackages" yaml:"hoistWorkspacePackages"`
	// A proxy to use for outgoing HTTPS requests.
	//
	// If the HTTPS_PROXY, https_proxy, HTTP_PROXY or http_proxy environment variables are set, their values will be used instead.
	// Experimental.
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// During installation the dependencies of some packages are automatically patched.
	//
	// If you want to disable this, set this config to false.
	// Experimental.
	IgnoreCompatibilityDb *bool `field:"optional" json:"ignoreCompatibilityDb" yaml:"ignoreCompatibilityDb"`
	// A list of package names that should not be built during installation.
	// Experimental.
	IgnoredBuiltDependencies *[]*string `field:"optional" json:"ignoredBuiltDependencies" yaml:"ignoredBuiltDependencies"`
	// Do not execute any scripts of the installed packages.
	//
	// Scripts of the projects are executed.
	// Experimental.
	IgnoreDepScripts *bool `field:"optional" json:"ignoreDepScripts" yaml:"ignoreDepScripts"`
	// A list of optional dependencies that the install should be skipped.
	// Experimental.
	IgnoredOptionalDependencies *[]*string `field:"optional" json:"ignoredOptionalDependencies" yaml:"ignoredOptionalDependencies"`
	// Default is undefined.
	//
	// Errors out when a patch with an exact version or version range fails. Ignores failures from name-only patches. When true, prints a warning instead of failing when any patch cannot be applied. When false, errors out for any patch failure.
	// Default: undefined. Errors out when a patch with an exact version or version range fails. Ignores failures from name-only patches. When true, prints a warning instead of failing when any patch cannot be applied. When false, errors out for any patch failure.
	//
	// Experimental.
	IgnorePatchFailures *bool `field:"optional" json:"ignorePatchFailures" yaml:"ignorePatchFailures"`
	// .pnpmfile.cjs will be ignored. Useful together with --ignore-scripts when you want to make sure that no script gets executed during install.
	// Experimental.
	IgnorePnpmfile *bool `field:"optional" json:"ignorePnpmfile" yaml:"ignorePnpmfile"`
	// Do not execute any scripts defined in the project package.json and its dependencies.
	// Experimental.
	IgnoreScripts *bool `field:"optional" json:"ignoreScripts" yaml:"ignoreScripts"`
	// When set to true, no workspace cycle warnings will be printed.
	// Experimental.
	IgnoreWorkspaceCycles *bool `field:"optional" json:"ignoreWorkspaceCycles" yaml:"ignoreWorkspaceCycles"`
	// Adding a new dependency to the root workspace package fails, unless the --ignore-workspace-root-check or -w flag is used.
	// Experimental.
	IgnoreWorkspaceRootCheck *bool `field:"optional" json:"ignoreWorkspaceRootCheck" yaml:"ignoreWorkspaceRootCheck"`
	// When executing commands recursively in a workspace, execute them on the root workspace project as well.
	// Experimental.
	IncludeWorkspaceRoot *bool `field:"optional" json:"includeWorkspaceRoot" yaml:"includeWorkspaceRoot"`
	// Enables hard-linking of all local workspace dependencies instead of symlinking them.
	// Experimental.
	InjectWorkspacePackages *bool `field:"optional" json:"injectWorkspacePackages" yaml:"injectWorkspacePackages"`
	// A client key to pass when accessing the registry.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// If this is enabled, locally available packages are linked to node_modules instead of being downloaded from the registry.
	// Experimental.
	LinkWorkspacePackages PnpmWorkspaceYamlSchemaLinkWorkspacePackages `field:"optional" json:"linkWorkspacePackages" yaml:"linkWorkspacePackages"`
	// The IP address of the local interface to use when making connections to the npm registry.
	// Experimental.
	LocalAddress *string `field:"optional" json:"localAddress" yaml:"localAddress"`
	// When set to false, pnpm won't read or generate a pnpm-lock.yaml file.
	// Experimental.
	Lockfile *bool `field:"optional" json:"lockfile" yaml:"lockfile"`
	// Add the full URL to the package's tarball to every entry in pnpm-lock.yaml.
	// Experimental.
	LockfileIncludeTarballUrl *bool `field:"optional" json:"lockfileIncludeTarballUrl" yaml:"lockfileIncludeTarballUrl"`
	// Any logs at or higher than the given level will be shown.
	// Experimental.
	Loglevel PnpmWorkspaceYamlSchemaLoglevel `field:"optional" json:"loglevel" yaml:"loglevel"`
	// When enabled, pnpm will automatically download and run the version of pnpm specified in the packageManager field of package.json.
	// Experimental.
	ManagePackageManagerVersions *bool `field:"optional" json:"managePackageManagerVersions" yaml:"managePackageManagerVersions"`
	// The maximum number of connections to use per origin (protocol/host/port combination).
	// Experimental.
	Maxsockets *float64 `field:"optional" json:"maxsockets" yaml:"maxsockets"`
	// This configuration matches the current branch name to determine whether to merge all git branch lockfile files.
	// Experimental.
	MergeGitBranchLockfilesBranchPattern *[]interface{} `field:"optional" json:"mergeGitBranchLockfilesBranchPattern" yaml:"mergeGitBranchLockfilesBranchPattern"`
	// minimumReleaseAge defines the minimum number of minutes that must pass after a version is published before pnpm will install it.
	//
	// This applies to all dependencies, including transitive ones.
	// Experimental.
	MinimumReleaseAge *float64 `field:"optional" json:"minimumReleaseAge" yaml:"minimumReleaseAge"`
	// If you set `minimumReleaseAge` but need certain dependencies to always install the newest version immediately, you can list them under `minimumReleaseAgeExclude`.
	//
	// The exclusion works by `package name` and applies to all versions of that package.
	// Experimental.
	MinimumReleaseAgeExclude *[]*string `field:"optional" json:"minimumReleaseAgeExclude" yaml:"minimumReleaseAgeExclude"`
	// When `true`, pnpm skips the `minimumReleaseAge` check for a package whose registry metadata does not include the time field (some private registries and mirrors omit it).
	//
	// Set to `false` to fail resolution in that case instead of installing the package.
	// Experimental.
	MinimumReleaseAgeIgnoreMissingTime *bool `field:"optional" json:"minimumReleaseAgeIgnoreMissingTime" yaml:"minimumReleaseAgeIgnoreMissingTime"`
	// Controls how pnpm behaves when no version of a dependency satisfies the minimumReleaseAge constraint within the requested range.
	//
	// https://pnpm.io/settings#minimumreleaseagestrict
	// Experimental.
	MinimumReleaseAgeStrict *bool `field:"optional" json:"minimumReleaseAgeStrict" yaml:"minimumReleaseAgeStrict"`
	// The time in minutes after which orphan packages from the modules directory should be removed.
	// Experimental.
	ModulesCacheMaxAge *float64 `field:"optional" json:"modulesCacheMaxAge" yaml:"modulesCacheMaxAge"`
	// The directory in which dependencies will be installed (instead of node_modules).
	// Experimental.
	ModulesDir *string `field:"optional" json:"modulesDir" yaml:"modulesDir"`
	// Controls the maximum number of HTTP(S) requests to process simultaneously.
	// Experimental.
	NetworkConcurrency *float64 `field:"optional" json:"networkConcurrency" yaml:"networkConcurrency"`
	// A list of dependencies to run builds for.
	// Experimental.
	NeverBuiltDependencies *[]*string `field:"optional" json:"neverBuiltDependencies" yaml:"neverBuiltDependencies"`
	// Configure custom Node.js download mirrors in `pnpm-workspace.yaml`. The keys are release channels (`release`, `rc`, `nightly`, `v8-canary`, etc.) and the values are base URLs.
	// Experimental.
	NodeDownloadMirrors *map[string]*string `field:"optional" json:"nodeDownloadMirrors" yaml:"nodeDownloadMirrors"`
	// Defines what linker should be used for installing Node packages.
	// Experimental.
	NodeLinker PnpmWorkspaceYamlSchemaNodeLinker `field:"optional" json:"nodeLinker" yaml:"nodeLinker"`
	// Options to pass through to Node.js via the NODE_OPTIONS environment variable.
	// Experimental.
	NodeOptions *string `field:"optional" json:"nodeOptions" yaml:"nodeOptions"`
	// The Node.js version to use when checking a package's engines setting.
	// Experimental.
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// A comma-separated string of domain extensions that a proxy should not be used for.
	// Experimental.
	Noproxy *string `field:"optional" json:"noproxy" yaml:"noproxy"`
	// The location of the npm binary that pnpm uses for some actions, like publishing.
	// Experimental.
	NpmPath *string `field:"optional" json:"npmPath" yaml:"npmPath"`
	// The path to a file containing registry authentication tokens.
	//
	// By default, pnpm reads auth tokens from ~/.npmrc as a fallback for registry authentication. Use this setting to point to a different file instead.
	// Experimental.
	NpmrcAuthFile *string `field:"optional" json:"npmrcAuthFile" yaml:"npmrcAuthFile"`
	// A list of package names that are allowed to be executed during installation.
	// Experimental.
	OnlyBuiltDependencies *[]*string `field:"optional" json:"onlyBuiltDependencies" yaml:"onlyBuiltDependencies"`
	// Specifies a JSON file that lists the only packages permitted to run installation scripts during the pnpm install process.
	// Experimental.
	OnlyBuiltDependenciesFile *string `field:"optional" json:"onlyBuiltDependenciesFile" yaml:"onlyBuiltDependenciesFile"`
	// When enabled, a fast check will be performed before proceeding to installation.
	//
	// This way a repeat install or an install on a project with everything up-to-date becomes a lot faster.
	// Experimental.
	OptimisticRepeatInstall *bool `field:"optional" json:"optimisticRepeatInstall" yaml:"optimisticRepeatInstall"`
	// Used to override any dependency in the dependency graph.
	// Experimental.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Used to extend the existing package definitions with additional information.
	// Experimental.
	PackageExtensions interface{} `field:"optional" json:"packageExtensions" yaml:"packageExtensions"`
	// Controls the way packages are imported from the store (if you want to disable symlinks inside node_modules, then you need to change the nodeLinker setting, not this one).
	// Experimental.
	PackageImportMethod PnpmWorkspaceYamlSchemaPackageImportMethod `field:"optional" json:"packageImportMethod" yaml:"packageImportMethod"`
	// If this setting is disabled, pnpm will not fail if a different package manager is specified in the packageManager field of package.json. When enabled, only the package name is checked (since pnpm v9.2.0), so you can still run any version of pnpm regardless of the version specified in the packageManager field.
	// Experimental.
	PackageManagerStrict *bool `field:"optional" json:"packageManagerStrict" yaml:"packageManagerStrict"`
	// When enabled, pnpm will fail if its version doesn't exactly match the version specified in the packageManager field of package.json.
	// Experimental.
	PackageManagerStrictVersion *bool `field:"optional" json:"packageManagerStrictVersion" yaml:"packageManagerStrictVersion"`
	// Workspace package paths.
	//
	// Glob patterns are supported.
	// Experimental.
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// A list of dependencies that are patched.
	// Experimental.
	PatchedDependencies *map[string]*string `field:"optional" json:"patchedDependencies" yaml:"patchedDependencies"`
	// The generated patch file will be saved to this directory.
	// Experimental.
	PatchesDir *string `field:"optional" json:"patchesDir" yaml:"patchesDir"`
	// Experimental.
	PeerDependencyRules *PnpmWorkspaceYamlSchemaPeerDependencyRules `field:"optional" json:"peerDependencyRules" yaml:"peerDependencyRules"`
	// Max length of the peer IDs suffix added to dependency keys in the lockfile.
	//
	// If the suffix is longer, it is replaced with a hash.
	// Experimental.
	PeersSuffixMaxLength *float64 `field:"optional" json:"peersSuffixMaxLength" yaml:"peersSuffixMaxLength"`
	// Overrides the `onFail` behavior of both the `packageManager` field and `devEngines.packageManager` when the running pnpm version does not match the declared one.
	// Experimental.
	PmOnFail PnpmWorkspaceYamlSchemaPmOnFail `field:"optional" json:"pmOnFail" yaml:"pmOnFail"`
	// The location of the local pnpmfile.
	// Experimental.
	Pnpmfile *string `field:"optional" json:"pnpmfile" yaml:"pnpmfile"`
	// When set to true and the available pnpm-lock.yaml satisfies the package.json dependencies directive, a headless installation is performed.
	// Experimental.
	PreferFrozenLockfile *bool `field:"optional" json:"preferFrozenLockfile" yaml:"preferFrozenLockfile"`
	// Bypass staleness checks for cached data.
	//
	// Missing data will still be requested from the server.
	// Experimental.
	PreferOffline *bool `field:"optional" json:"preferOffline" yaml:"preferOffline"`
	// Create symlinks to executables in node_modules/.bin instead of command shims. This setting is ignored on Windows, where only command shims work.
	// Experimental.
	PreferSymlinkedExecutables *bool `field:"optional" json:"preferSymlinkedExecutables" yaml:"preferSymlinkedExecutables"`
	// If this is enabled, local packages from the workspace are preferred over packages from the registry, even if there is a newer version of the package in the registry.
	// Experimental.
	PreferWorkspacePackages *bool `field:"optional" json:"preferWorkspacePackages" yaml:"preferWorkspacePackages"`
	// When publishing from a supported cloud CI/CD system, the package will be publicly linked to where it was built and published from.
	// Experimental.
	Provenance *bool `field:"optional" json:"provenance" yaml:"provenance"`
	// A proxy to use for outgoing http requests.
	//
	// If the HTTP_PROXY or http_proxy environment variables are set, proxy settings will be honored by the underlying request library.
	// Experimental.
	Proxy *string `field:"optional" json:"proxy" yaml:"proxy"`
	// Unlike hoistPattern, which hoists dependencies to a hidden modules directory inside the virtual store, publicHoistPattern hoists dependencies matching the pattern to the root modules directory.
	// Experimental.
	PublicHoistPattern *[]*string `field:"optional" json:"publicHoistPattern" yaml:"publicHoistPattern"`
	// The primary branch of the repository which is used for publishing the latest changes.
	// Experimental.
	PublishBranch *string `field:"optional" json:"publishBranch" yaml:"publishBranch"`
	// If this is enabled, the primary behaviour of pnpm install becomes that of pnpm install -r, meaning the install is performed on all workspace or subdirectory packages.
	// Experimental.
	RecursiveInstall *bool `field:"optional" json:"recursiveInstall" yaml:"recursiveInstall"`
	// Configure registries for scoped packages in `pnpm-workspace.yaml`. The `default` key sets the main registry (equivalent to the `registry` `.npmrc` setting). Scoped keys configure registries for specific package scopes.
	// Experimental.
	Registries *map[string]*string `field:"optional" json:"registries" yaml:"registries"`
	// The base URL of the npm package registry (trailing slash included).
	// Experimental.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// Set this to true if the registry that you are using returns the "time" field in the abbreviated metadata.
	// Experimental.
	RegistrySupportsTimeField *bool `field:"optional" json:"registrySupportsTimeField" yaml:"registrySupportsTimeField"`
	// Allows you to customize the output style of the logs.
	//
	// https://pnpm.io/cli/install#--reportername
	// Experimental.
	Reporter PnpmWorkspaceYamlSchemaReporter `field:"optional" json:"reporter" yaml:"reporter"`
	// A list of scripts that must exist in each project.
	// Experimental.
	RequiredScripts *[]*string `field:"optional" json:"requiredScripts" yaml:"requiredScripts"`
	// Determines how pnpm resolves dependencies, See https://pnpm.io/settings#resolutionmode.
	// Experimental.
	ResolutionMode PnpmWorkspaceYamlSchemaResolutionMode `field:"optional" json:"resolutionMode" yaml:"resolutionMode"`
	// When enabled, dependencies of the root workspace project are used to resolve peer dependencies of any projects in the workspace.
	// Experimental.
	ResolvePeersFromWorkspaceRoot *bool `field:"optional" json:"resolvePeersFromWorkspaceRoot" yaml:"resolvePeersFromWorkspaceRoot"`
	// Overrides the `onFail` field of `devEngines.runtime` (and `engines.runtime`) in the root project's `package.json`. This is useful when you want a different local behavior than what is written in the manifest — for instance, forcing pnpm to download the declared runtime even when the manifest sets `onFail: "warn"`.
	// Experimental.
	RuntimeOnFail PnpmWorkspaceYamlSchemaRuntimeOnFail `field:"optional" json:"runtimeOnFail" yaml:"runtimeOnFail"`
	// Saved dependencies will be configured with an exact version rather than using pnpm's default semver range operator.
	// Experimental.
	SaveExact *bool `field:"optional" json:"saveExact" yaml:"saveExact"`
	// Configure how versions of packages installed to a package.json file get prefixed.
	// Experimental.
	SavePrefix PnpmWorkspaceYamlSchemaSavePrefix `field:"optional" json:"savePrefix" yaml:"savePrefix"`
	// This setting controls how dependencies that are linked from the workspace are added to package.json.
	// Experimental.
	SaveWorkspaceProtocol PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol `field:"optional" json:"saveWorkspaceProtocol" yaml:"saveWorkspaceProtocol"`
	// The shell to use for scripts run with the pnpm run command.
	// Experimental.
	ScriptShell *string `field:"optional" json:"scriptShell" yaml:"scriptShell"`
	// By default, pnpm creates a semistrict node_modules, meaning dependencies have access to undeclared dependencies but modules outside of node_modules do not.
	// Experimental.
	ShamefullyHoist *bool `field:"optional" json:"shamefullyHoist" yaml:"shamefullyHoist"`
	// If this is enabled, pnpm creates a single pnpm-lock.yaml file in the root of the workspace.
	// Experimental.
	SharedWorkspaceLockfile *bool `field:"optional" json:"sharedWorkspaceLockfile" yaml:"sharedWorkspaceLockfile"`
	// When true, pnpm will use a JavaScript implementation of a bash-like shell to execute scripts.
	// Experimental.
	ShellEmulator *bool `field:"optional" json:"shellEmulator" yaml:"shellEmulator"`
	// Use and cache the results of (pre/post)install hooks.
	// Experimental.
	SideEffectsCache *bool `field:"optional" json:"sideEffectsCache" yaml:"sideEffectsCache"`
	// Only use the side effects cache if present, do not create it for new packages.
	// Experimental.
	SideEffectsCacheReadonly *bool `field:"optional" json:"sideEffectsCacheReadonly" yaml:"sideEffectsCacheReadonly"`
	// The location where all the packages are saved on the disk.
	// Experimental.
	StateDir *string `field:"optional" json:"stateDir" yaml:"stateDir"`
	// The location where all the packages are saved on the disk.
	// Experimental.
	StoreDir *string `field:"optional" json:"storeDir" yaml:"storeDir"`
	// When strictDepBuilds is enabled, the installation will exit with a non-zero exit code if any dependencies have unreviewed build scripts (aka postinstall scripts).
	// Experimental.
	StrictDepBuilds *bool `field:"optional" json:"strictDepBuilds" yaml:"strictDepBuilds"`
	// If this is enabled, commands will fail if there is a missing or invalid peer dependency in the tree.
	// Experimental.
	StrictPeerDependencies *bool `field:"optional" json:"strictPeerDependencies" yaml:"strictPeerDependencies"`
	// Whether or not to do SSL key validation when making requests to the registry via HTTPS.
	// Experimental.
	StrictSsl *bool `field:"optional" json:"strictSsl" yaml:"strictSsl"`
	// Some registries allow the exact same content to be published under different package names and/or versions.
	// Experimental.
	StrictStorePkgContentCheck *bool `field:"optional" json:"strictStorePkgContentCheck" yaml:"strictStorePkgContentCheck"`
	// Specifies architectures for which you'd like to install optional dependencies, even if they don't match the architecture of the system running the install.
	// Experimental.
	SupportedArchitectures *PnpmWorkspaceYamlSchemaSupportedArchitectures `field:"optional" json:"supportedArchitectures" yaml:"supportedArchitectures"`
	// When symlink is set to false, pnpm creates a virtual store directory without any symlinks.
	//
	// It is a useful setting together with nodeLinker=pnp.
	// Experimental.
	Symlink *bool `field:"optional" json:"symlink" yaml:"symlink"`
	// Injected workspace dependencies are collections of hardlinks, which don't add or remove the files when their sources change.
	// Experimental.
	SyncInjectedDepsAfterScripts *[]*string `field:"optional" json:"syncInjectedDepsAfterScripts" yaml:"syncInjectedDepsAfterScripts"`
	// If you pnpm add a package and you don't provide a specific version, then it will install the package at the version registered under the tag from this setting.
	// Experimental.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// A new trustLockfile setting controls whether pnpm install re-applies the `minimumReleaseAge` / `trustPolicy: 'no-downgrade'` checks to every entry in the loaded lockfile.
	//
	// When true, the install treats the lockfile as already-trusted and skips the verification pass — useful for closed-source projects where every commit comes from a trusted author. The default is false, so verification stays on by default.
	// Experimental.
	TrustLockfile *bool `field:"optional" json:"trustLockfile" yaml:"trustLockfile"`
	// When set to no-downgrade, pnpm will fail if a package's trust level has decreased compared to previous releases.
	//
	// For example, if a package was previously published by a trusted publisher but now only has provenance or no trust evidence, installation will fail. This helps prevent installing potentially compromised versions.
	// Experimental.
	TrustPolicy PnpmWorkspaceYamlSchemaTrustPolicy `field:"optional" json:"trustPolicy" yaml:"trustPolicy"`
	// You can now list one or more specific packages or versions that pnpm should allow to install, even if those packages don't satisfy the trust policy requirement.
	// Experimental.
	TrustPolicyExclude *[]*string `field:"optional" json:"trustPolicyExclude" yaml:"trustPolicyExclude"`
	// Allows ignoring the trust policy check for packages published more than the specified number of minutes ago.
	//
	// This is useful when enabling strict trust policies, as it allows older versions of packages (which may lack a process for publishing with signatures or provenance) to be installed without manual exclusion, assuming they are safe due to their age.
	// Experimental.
	TrustPolicyIgnoreAfter *float64 `field:"optional" json:"trustPolicyIgnoreAfter" yaml:"trustPolicyIgnoreAfter"`
	// Set to true to enable UID/GID switching when running package scripts.
	//
	// If set explicitly to false, then installing as a non-root user will fail.
	// Experimental.
	UnsafePerm *bool `field:"optional" json:"unsafePerm" yaml:"unsafePerm"`
	// Experimental.
	UpdateConfig *PnpmWorkspaceYamlSchemaUpdateConfig `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// When true, pnpm will check for updates to the installed packages and notify the user.
	// Experimental.
	UpdateNotifier *bool `field:"optional" json:"updateNotifier" yaml:"updateNotifier"`
	// Experimental option that enables beta features of the CLI.
	// Experimental.
	UseBetaCli *bool `field:"optional" json:"useBetaCli" yaml:"useBetaCli"`
	// Specifies which exact Node.js version should be used for the project's runtime.
	// Experimental.
	UseNodeVersion *string `field:"optional" json:"useNodeVersion" yaml:"useNodeVersion"`
	// When true, all the output is written to stderr.
	// Experimental.
	UseStderr *bool `field:"optional" json:"useStderr" yaml:"useStderr"`
	// This setting allows the checking of the state of dependencies before running scripts.
	// Experimental.
	VerifyDepsBeforeRun interface{} `field:"optional" json:"verifyDepsBeforeRun" yaml:"verifyDepsBeforeRun"`
	// By default, if a file in the store has been modified, the content of this file is checked before linking it to a project's node_modules.
	// Experimental.
	VerifyStoreIntegrity *bool `field:"optional" json:"verifyStoreIntegrity" yaml:"verifyStoreIntegrity"`
	// The directory with links to the store.
	// Experimental.
	VirtualStoreDir *string `field:"optional" json:"virtualStoreDir" yaml:"virtualStoreDir"`
	// Sets the maximum allowed length of directory names inside the virtual store directory (node_modules/.pnpm).
	// Experimental.
	VirtualStoreDirMaxLength *float64 `field:"optional" json:"virtualStoreDirMaxLength" yaml:"virtualStoreDirMaxLength"`
	// When set to true, pnpm populates the virtual store without creating importer symlinks, hoisting, bin links, or running lifecycle scripts.
	//
	// This is useful for pre-populating a store (e.g., in Nix builds) without creating unnecessary project-level artifacts. pnpm fetch uses this mode internally.
	// Experimental.
	VirtualStoreOnly *bool `field:"optional" json:"virtualStoreOnly" yaml:"virtualStoreOnly"`
	// Set the maximum number of tasks to run simultaneously.
	//
	// For unlimited concurrency use Infinity. You can set the value to <= 0 and it will use amount of CPU cores of the host minus the absolute value of the provided number as: max(1, (number of cores) - abs(workspaceConcurrency)).
	// Experimental.
	WorkspaceConcurrency *float64 `field:"optional" json:"workspaceConcurrency" yaml:"workspaceConcurrency"`
}

