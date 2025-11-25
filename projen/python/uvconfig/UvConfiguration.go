package uvconfig


// Metadata and configuration for uv.
// Experimental.
type UvConfiguration struct {
	// The default version specifier when adding a dependency.
	//
	// When adding a dependency to the project, if no constraint or URL is provided, a constraint
	// is added based on the latest compatible version of the package. By default, a lower bound
	// constraint is used, e.g., `>=1.2.3`.
	//
	// When `--frozen` is provided, no resolution is performed, and dependencies are always added
	// without constraints.
	//
	// This option is in preview and may change in any future release.
	// Experimental.
	AddBounds *string `field:"optional" json:"addBounds" yaml:"addBounds"`
	// Allow insecure connections to host.
	//
	// Expects to receive either a hostname (e.g., `localhost`), a host-port pair (e.g.,
	// `localhost:8080`), or a URL (e.g., `https://localhost`).
	//
	// WARNING: Hosts included in this list will not be verified against the system's certificate
	// store. Only use `--allow-insecure-host` in a secure network with verified sources, as it
	// bypasses SSL verification and could expose you to MITM attacks.
	// Experimental.
	AllowInsecureHost *[]*string `field:"optional" json:"allowInsecureHost" yaml:"allowInsecureHost"`
	// Configuration for the uv build backend.
	//
	// Note that those settings only apply when using the `uv_build` backend, other build backends
	// (such as hatchling) have their own configuration.
	// Experimental.
	BuildBackend *BuildBackendSettings `field:"optional" json:"buildBackend" yaml:"buildBackend"`
	// PEP 508-style requirements, e.g., `ruff==0.5.0`, or `ruff @ https://...`.
	// Experimental.
	BuildConstraintDependencies *[]*string `field:"optional" json:"buildConstraintDependencies" yaml:"buildConstraintDependencies"`
	// Path to the cache directory.
	// Experimental.
	CacheDir *string `field:"optional" json:"cacheDir" yaml:"cacheDir"`
	// The keys to consider when caching builds for the project.
	//
	// Cache keys enable you to specify the files or directories that should trigger a rebuild when
	// modified. By default, uv will rebuild a project whenever the `pyproject.toml`, `setup.py`,
	// or `setup.cfg` files in the project directory are modified, or if a `src` directory is
	// added or removed, i.e.:
	//
	// ```toml
	// cache-keys = [{ file = "pyproject.toml" }, { file = "setup.py" }, { file = "setup.cfg" }, { dir = "src" }]
	// ```
	//
	// As an example: if a project uses dynamic metadata to read its dependencies from a
	// `requirements.txt` file, you can specify `cache-keys = [{ file = "requirements.txt" }, { file = "pyproject.toml" }]`
	// to ensure that the project is rebuilt whenever the `requirements.txt` file is modified (in
	// addition to watching the `pyproject.toml`).
	//
	// Globs are supported, following the syntax of the [`glob`](https://docs.rs/glob/0.3.1/glob/struct.Pattern.html)
	// crate. For example, to invalidate the cache whenever a `.toml` file in the project directory
	// or any of its subdirectories is modified, you can specify `cache-keys = [{ file = "*_/*.toml" }]`.
	// Note that the use of globs can be expensive, as uv may need to walk the filesystem to
	// determine whether any files have changed.
	//
	// Cache keys can also include version control information. For example, if a project uses
	// `setuptools_scm` to read its version from a Git commit, you can specify `cache-keys = [{ git = { commit = true }, { file = "pyproject.toml" }]`
	// to include the current Git commit hash in the cache key (in addition to the
	// `pyproject.toml`). Git tags are also supported via `cache-keys = [{ git = { commit = true, tags = true } }]`.
	//
	// Cache keys can also include environment variables. For example, if a project relies on
	// `MACOSX_DEPLOYMENT_TARGET` or other environment variables to determine its behavior, you can
	// specify `cache-keys = [{ env = "MACOSX_DEPLOYMENT_TARGET" }]` to invalidate the cache
	// whenever the environment variable changes.
	//
	// Cache keys only affect the project defined by the `pyproject.toml` in which they're
	// specified (as opposed to, e.g., affecting all members in a workspace), and all paths and
	// globs are interpreted as relative to the project directory.
	// Experimental.
	CacheKeys *[]interface{} `field:"optional" json:"cacheKeys" yaml:"cacheKeys"`
	// Check an index URL for existing files to skip duplicate uploads.
	//
	// This option allows retrying publishing that failed after only some, but not all files have
	// been uploaded, and handles error due to parallel uploads of the same file.
	//
	// Before uploading, the index is checked. If the exact same file already exists in the index,
	// the file will not be uploaded. If an error occurred during the upload, the index is checked
	// again, to handle cases where the identical file was uploaded twice in parallel.
	//
	// The exact behavior will vary based on the index. When uploading to PyPI, uploading the same
	// file succeeds even without `--check-url`, while most other indexes error.
	//
	// The index must provide one of the supported hashes (SHA-256, SHA-384, or SHA-512).
	// Experimental.
	CheckUrl *string `field:"optional" json:"checkUrl" yaml:"checkUrl"`
	// Compile Python files to bytecode after installation.
	//
	// By default, uv does not compile Python (`.py`) files to bytecode (`__pycache__/*.pyc`);
	// instead, compilation is performed lazily the first time a module is imported. For use-cases
	// in which start time is critical, such as CLI applications and Docker containers, this option
	// can be enabled to trade longer installation times for faster start times.
	//
	// When enabled, uv will process the entire site-packages directory (including packages that
	// are not being modified by the current operation) for consistency. Like pip, it will also
	// ignore errors.
	// Experimental.
	CompileBytecode *bool `field:"optional" json:"compileBytecode" yaml:"compileBytecode"`
	// The maximum number of source distributions that uv will build concurrently at any given time.
	//
	// Defaults to the number of available CPU cores.
	// Default: the number of available CPU cores.
	//
	// Experimental.
	ConcurrentBuilds *float64 `field:"optional" json:"concurrentBuilds" yaml:"concurrentBuilds"`
	// The maximum number of in-flight concurrent downloads that uv will perform at any given time.
	// Experimental.
	ConcurrentDownloads *float64 `field:"optional" json:"concurrentDownloads" yaml:"concurrentDownloads"`
	// The number of threads used when installing and unzipping packages.
	//
	// Defaults to the number of available CPU cores.
	// Default: the number of available CPU cores.
	//
	// Experimental.
	ConcurrentInstalls *float64 `field:"optional" json:"concurrentInstalls" yaml:"concurrentInstalls"`
	// Settings to pass to the [PEP 517](https://peps.python.org/pep-0517/) build backend, specified as `KEY=VALUE` pairs.
	// Experimental.
	ConfigSettings *map[string]interface{} `field:"optional" json:"configSettings" yaml:"configSettings"`
	// Settings to pass to the [PEP 517](https://peps.python.org/pep-0517/) build backend for specific packages, specified as `KEY=VALUE` pairs.
	//
	// Accepts a map from package names to string key-value pairs.
	// Experimental.
	ConfigSettingsPackage *map[string]*map[string]interface{} `field:"optional" json:"configSettingsPackage" yaml:"configSettingsPackage"`
	// A list of sets of conflicting groups or extras.
	// Experimental.
	Conflicts *[]*[]*SchemaConflictItem `field:"optional" json:"conflicts" yaml:"conflicts"`
	// PEP 508-style requirements, e.g., `ruff==0.5.0`, or `ruff @ https://...`.
	// Experimental.
	ConstraintDependencies *[]*string `field:"optional" json:"constraintDependencies" yaml:"constraintDependencies"`
	// The list of `dependency-groups` to install by default.
	//
	// Can also be the literal `"all"` to default enable all groups.
	// Experimental.
	DefaultGroups interface{} `field:"optional" json:"defaultGroups" yaml:"defaultGroups"`
	// Additional settings for `dependency-groups`.
	//
	// Currently this can only be used to add `requires-python` constraints
	// to dependency groups (typically to inform uv that your dev tooling
	// has a higher python requirement than your actual project).
	//
	// This cannot be used to define dependency groups, use the top-level
	// `[dependency-groups]` table for that.
	// Experimental.
	DependencyGroups *map[string]*DependencyGroupSettings `field:"optional" json:"dependencyGroups" yaml:"dependencyGroups"`
	// Pre-defined static metadata for dependencies of the project (direct or transitive).
	//
	// When
	// provided, enables the resolver to use the specified metadata instead of querying the
	// registry or building the relevant package from source.
	//
	// Metadata should be provided in adherence with the [Metadata 2.3](https://packaging.python.org/en/latest/specifications/core-metadata/)
	// standard, though only the following fields are respected:
	//
	// - `name`: The name of the package.
	// - (Optional) `version`: The version of the package. If omitted, the metadata will be applied
	// to all versions of the package.
	// - (Optional) `requires-dist`: The dependencies of the package (e.g., `werkzeug>=0.14`).
	// - (Optional) `requires-python`: The Python version required by the package (e.g., `>=3.10`).
	// - (Optional) `provides-extra`: The extras provided by the package.
	// Experimental.
	DependencyMetadata *[]*StaticMetadata `field:"optional" json:"dependencyMetadata" yaml:"dependencyMetadata"`
	// PEP 508-style requirements, e.g., `ruff==0.5.0`, or `ruff @ https://...`.
	// Experimental.
	DevDependencies *[]*string `field:"optional" json:"devDependencies" yaml:"devDependencies"`
	// A list of environment markers, e.g., `python_version >= '3.6'`.
	// Experimental.
	Environments *[]*string `field:"optional" json:"environments" yaml:"environments"`
	// Package names to exclude, e.g., `werkzeug`, `numpy`.
	// Experimental.
	ExcludeDependencies *[]*string `field:"optional" json:"excludeDependencies" yaml:"excludeDependencies"`
	// Limit candidate packages to those that were uploaded prior to a given point in time.
	//
	// Accepts a superset of [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339.html) (e.g.,
	// `2006-12-02T02:07:43Z`). A full timestamp is required to ensure that the resolver will
	// behave consistently across timezones.
	// Experimental.
	ExcludeNewer *string `field:"optional" json:"excludeNewer" yaml:"excludeNewer"`
	// Limit candidate packages for specific packages to those that were uploaded prior to the given date.
	//
	// Accepts package-date pairs in a dictionary format.
	// Experimental.
	ExcludeNewerPackage *map[string]*string `field:"optional" json:"excludeNewerPackage" yaml:"excludeNewerPackage"`
	// Additional build dependencies for packages.
	//
	// This allows extending the PEP 517 build environment for the project's dependencies with
	// additional packages. This is useful for packages that assume the presence of packages like
	// `pip`, and do not declare them as build dependencies.
	// Experimental.
	ExtraBuildDependencies *map[string]*[]interface{} `field:"optional" json:"extraBuildDependencies" yaml:"extraBuildDependencies"`
	// Extra environment variables to set when building certain packages.
	//
	// Environment variables will be added to the environment when building the
	// specified packages.
	// Experimental.
	ExtraBuildVariables *map[string]*map[string]*string `field:"optional" json:"extraBuildVariables" yaml:"extraBuildVariables"`
	// Extra URLs of package indexes to use, in addition to `--index-url`.
	//
	// Accepts either a repository compliant with [PEP 503](https://peps.python.org/pep-0503/)
	// (the simple repository API), or a local directory laid out in the same format.
	//
	// All indexes provided via this flag take priority over the index specified by
	// [`index_url`](#index-url) or [`index`](#index) with `default = true`. When multiple indexes
	// are provided, earlier values take priority.
	//
	// To control uv's resolution strategy when multiple indexes are present, see
	// [`index_strategy`](#index-strategy).
	//
	// (Deprecated: use `index` instead.)
	// Experimental.
	ExtraIndexUrl *[]*string `field:"optional" json:"extraIndexUrl" yaml:"extraIndexUrl"`
	// Locations to search for candidate distributions, in addition to those found in the registry indexes.
	//
	// If a path, the target must be a directory that contains packages as wheel files (`.whl`) or
	// source distributions (e.g., `.tar.gz` or `.zip`) at the top level.
	//
	// If a URL, the page must contain a flat list of links to package files adhering to the
	// formats described above.
	// Experimental.
	FindLinks *[]*string `field:"optional" json:"findLinks" yaml:"findLinks"`
	// The strategy to use when selecting multiple versions of a given package across Python versions and platforms.
	//
	// By default, uv will optimize for selecting the latest version of each package for each
	// supported Python version (`requires-python`), while minimizing the number of selected
	// versions across platforms.
	//
	// Under `fewest`, uv will minimize the number of selected versions for each package,
	// preferring older versions that are compatible with a wider range of supported Python
	// versions or platforms.
	// Experimental.
	ForkStrategy *string `field:"optional" json:"forkStrategy" yaml:"forkStrategy"`
	// The indexes to use when resolving dependencies.
	//
	// Accepts either a repository compliant with [PEP 503](https://peps.python.org/pep-0503/)
	// (the simple repository API), or a local directory laid out in the same format.
	//
	// Indexes are considered in the order in which they're defined, such that the first-defined
	// index has the highest priority. Further, the indexes provided by this setting are given
	// higher priority than any indexes specified via [`index_url`](#index-url) or
	// [`extra_index_url`](#extra-index-url). uv will only consider the first index that contains
	// a given package, unless an alternative [index strategy](#index-strategy) is specified.
	//
	// If an index is marked as `explicit = true`, it will be used exclusively for the
	// dependencies that select it explicitly via `[tool.uv.sources]`, as in:
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "pytorch"
	// url = "https://download.pytorch.org/whl/cu121"
	// explicit = true
	//
	// [tool.uv.sources]
	// torch = { index = "pytorch" }
	// ```
	//
	// If an index is marked as `default = true`, it will be moved to the end of the prioritized list, such that it is
	// given the lowest priority when resolving packages. Additionally, marking an index as default will disable the
	// PyPI default index.
	// Experimental.
	Index *[]*Index `field:"optional" json:"index" yaml:"index"`
	// The strategy to use when resolving against multiple index URLs.
	//
	// By default, uv will stop at the first index on which a given package is available, and
	// limit resolutions to those present on that first index (`first-index`). This prevents
	// "dependency confusion" attacks, whereby an attacker can upload a malicious package under the
	// same name to an alternate index.
	// Experimental.
	IndexStrategy *string `field:"optional" json:"indexStrategy" yaml:"indexStrategy"`
	// The URL of the Python package index (by default: <https://pypi.org/simple>).
	//
	// Accepts either a repository compliant with [PEP 503](https://peps.python.org/pep-0503/)
	// (the simple repository API), or a local directory laid out in the same format.
	//
	// The index provided by this setting is given lower priority than any indexes specified via
	// [`extra_index_url`](#extra-index-url) or [`index`](#index).
	//
	// (Deprecated: use `index` instead.)
	// Experimental.
	IndexUrl *string `field:"optional" json:"indexUrl" yaml:"indexUrl"`
	// Attempt to use `keyring` for authentication for index URLs.
	//
	// At present, only `--keyring-provider subprocess` is supported, which configures uv to
	// use the `keyring` CLI to handle authentication.
	// Experimental.
	KeyringProvider *string `field:"optional" json:"keyringProvider" yaml:"keyringProvider"`
	// The method to use when installing packages from the global cache.
	//
	// Defaults to `clone` (also known as Copy-on-Write) on macOS, and `hardlink` on Linux and
	// Windows.
	//
	// WARNING: The use of symlink link mode is discouraged, as they create tight coupling between
	// the cache and the target environment. For example, clearing the cache (`uv cache clean`)
	// will break all installed packages by way of removing the underlying source files. Use
	// symlinks with caution.
	// Default: clone` (also known as Copy-on-Write) on macOS, and `hardlink` on Linux and.
	//
	// Experimental.
	LinkMode *string `field:"optional" json:"linkMode" yaml:"linkMode"`
	// Whether the project is managed by uv.
	//
	// If `false`, uv will ignore the project when
	// `uv run` is invoked.
	// Experimental.
	Managed *bool `field:"optional" json:"managed" yaml:"managed"`
	// Whether to load TLS certificates from the platform's native certificate store.
	//
	// By default, uv loads certificates from the bundled `webpki-roots` crate. The
	// `webpki-roots` are a reliable set of trust roots from Mozilla, and including them in uv
	// improves portability and performance (especially on macOS).
	//
	// However, in some cases, you may want to use the platform's native certificate store,
	// especially if you're relying on a corporate trust root (e.g., for a mandatory proxy) that's
	// included in your system's certificate store.
	// Experimental.
	NativeTls *bool `field:"optional" json:"nativeTls" yaml:"nativeTls"`
	// Don't install pre-built wheels.
	//
	// The given packages will be built and installed from source. The resolver will still use
	// pre-built wheels to extract package metadata, if available.
	// Experimental.
	NoBinary *bool `field:"optional" json:"noBinary" yaml:"noBinary"`
	// Don't install pre-built wheels for a specific package.
	// Experimental.
	NoBinaryPackage *[]*string `field:"optional" json:"noBinaryPackage" yaml:"noBinaryPackage"`
	// Don't build source distributions.
	//
	// When enabled, resolving will not run arbitrary Python code. The cached wheels of
	// already-built source distributions will be reused, but operations that require building
	// distributions will exit with an error.
	// Experimental.
	NoBuild *bool `field:"optional" json:"noBuild" yaml:"noBuild"`
	// Disable isolation when building source distributions.
	//
	// Assumes that build dependencies specified by [PEP 518](https://peps.python.org/pep-0518/)
	// are already installed.
	// Experimental.
	NoBuildIsolation *bool `field:"optional" json:"noBuildIsolation" yaml:"noBuildIsolation"`
	// Disable isolation when building source distributions for a specific package.
	//
	// Assumes that the packages' build dependencies specified by [PEP 518](https://peps.python.org/pep-0518/)
	// are already installed.
	// Experimental.
	NoBuildIsolationPackage *[]*string `field:"optional" json:"noBuildIsolationPackage" yaml:"noBuildIsolationPackage"`
	// Don't build source distributions for a specific package.
	// Experimental.
	NoBuildPackage *[]*string `field:"optional" json:"noBuildPackage" yaml:"noBuildPackage"`
	// Avoid reading from or writing to the cache, instead using a temporary directory for the duration of the operation.
	// Experimental.
	NoCache *bool `field:"optional" json:"noCache" yaml:"noCache"`
	// Ignore all registry indexes (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`.
	// Experimental.
	NoIndex *bool `field:"optional" json:"noIndex" yaml:"noIndex"`
	// Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any local or Git sources.
	// Experimental.
	NoSources *bool `field:"optional" json:"noSources" yaml:"noSources"`
	// Disable network access, relying only on locally cached data and locally available files.
	// Experimental.
	Offline *bool `field:"optional" json:"offline" yaml:"offline"`
	// PEP 508-style requirements, e.g., `ruff==0.5.0`, or `ruff @ https://...`.
	// Experimental.
	OverrideDependencies *[]*string `field:"optional" json:"overrideDependencies" yaml:"overrideDependencies"`
	// Whether the project should be considered a Python package, or a non-package ("virtual") project.
	//
	// Packages are built and installed into the virtual environment in editable mode and thus
	// require a build backend, while virtual projects are _not_ built or installed; instead, only
	// their dependencies are included in the virtual environment.
	//
	// Creating a package requires that a `build-system` is present in the `pyproject.toml`, and
	// that the project adheres to a structure that adheres to the build backend's expectations
	// (e.g., a `src` layout).
	// Experimental.
	Package *bool `field:"optional" json:"package" yaml:"package"`
	// Experimental.
	Pip *PipOptions `field:"optional" json:"pip" yaml:"pip"`
	// The strategy to use when considering pre-release versions.
	//
	// By default, uv will accept pre-releases for packages that _only_ publish pre-releases,
	// along with first-party requirements that contain an explicit pre-release marker in the
	// declared specifiers (`if-necessary-or-explicit`).
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Whether to enable experimental, preview features.
	// Experimental.
	Preview *bool `field:"optional" json:"preview" yaml:"preview"`
	// The URL for publishing packages to the Python package index (by default: <https://upload.pypi.org/legacy/>).
	// Experimental.
	PublishUrl *string `field:"optional" json:"publishUrl" yaml:"publishUrl"`
	// Mirror URL to use for downloading managed PyPy installations.
	//
	// By default, managed PyPy installations are downloaded from [downloads.python.org](https://downloads.python.org/).
	// This variable can be set to a mirror URL to use a different source for PyPy installations.
	// The provided URL will replace `https://downloads.python.org/pypy` in, e.g., `https://downloads.python.org/pypy/pypy3.8-v7.3.7-osx64.tar.bz2`.
	//
	// Distributions can be read from a
	// local directory by using the `file://` URL scheme.
	// Experimental.
	PypyInstallMirror *string `field:"optional" json:"pypyInstallMirror" yaml:"pypyInstallMirror"`
	// Whether to allow Python downloads.
	// Experimental.
	PythonDownloads *string `field:"optional" json:"pythonDownloads" yaml:"pythonDownloads"`
	// URL pointing to JSON of custom Python installations.
	//
	// Note that currently, only local paths are supported.
	// Experimental.
	PythonDownloadsJsonUrl *string `field:"optional" json:"pythonDownloadsJsonUrl" yaml:"pythonDownloadsJsonUrl"`
	// Mirror URL for downloading managed Python installations.
	//
	// By default, managed Python installations are downloaded from [`python-build-standalone`](https://github.com/astral-sh/python-build-standalone).
	// This variable can be set to a mirror URL to use a different source for Python installations.
	// The provided URL will replace `https://github.com/astral-sh/python-build-standalone/releases/download` in, e.g., `https://github.com/astral-sh/python-build-standalone/releases/download/20240713/cpython-3.12.4%2B20240713-aarch64-apple-darwin-install_only.tar.gz`.
	//
	// Distributions can be read from a local directory by using the `file://` URL scheme.
	// Experimental.
	PythonInstallMirror *string `field:"optional" json:"pythonInstallMirror" yaml:"pythonInstallMirror"`
	// Whether to prefer using Python installations that are already present on the system, or those that are downloaded and installed by uv.
	// Experimental.
	PythonPreference *string `field:"optional" json:"pythonPreference" yaml:"pythonPreference"`
	// Reinstall all packages, regardless of whether they're already installed.
	//
	// Implies `refresh`.
	// Experimental.
	Reinstall *bool `field:"optional" json:"reinstall" yaml:"reinstall"`
	// Reinstall a specific package, regardless of whether it's already installed.
	//
	// Implies
	// `refresh-package`.
	// Experimental.
	ReinstallPackage *[]*string `field:"optional" json:"reinstallPackage" yaml:"reinstallPackage"`
	// A list of environment markers, e.g., `sys_platform == 'darwin'.
	// Experimental.
	RequiredEnvironments *[]*string `field:"optional" json:"requiredEnvironments" yaml:"requiredEnvironments"`
	// Enforce a requirement on the version of uv.
	//
	// If the version of uv does not meet the requirement at runtime, uv will exit
	// with an error.
	//
	// Accepts a [PEP 440](https://peps.python.org/pep-0440/) specifier, like `==0.5.0` or `>=0.5.0`.
	// Experimental.
	RequiredVersion *string `field:"optional" json:"requiredVersion" yaml:"requiredVersion"`
	// The strategy to use when selecting between the different compatible versions for a given package requirement.
	//
	// By default, uv will use the latest compatible version of each package (`highest`).
	// Experimental.
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
	// The sources to use when resolving dependencies.
	//
	// `tool.uv.sources` enriches the dependency metadata with additional sources, incorporated
	// during development. A dependency source can be a Git repository, a URL, a local path, or an
	// alternative registry.
	//
	// See [Dependencies](https://docs.astral.sh/uv/concepts/projects/dependencies/) for more.
	// Experimental.
	Sources *map[string]*[]interface{} `field:"optional" json:"sources" yaml:"sources"`
	// Configure trusted publishing.
	//
	// By default, uv checks for trusted publishing when running in a supported environment, but
	// ignores it if it isn't configured.
	//
	// uv's supported environments for trusted publishing include GitHub Actions and GitLab CI/CD.
	// Experimental.
	TrustedPublishing TrustedPublishing `field:"optional" json:"trustedPublishing" yaml:"trustedPublishing"`
	// Allow package upgrades, ignoring pinned versions in any existing output file.
	// Experimental.
	Upgrade *bool `field:"optional" json:"upgrade" yaml:"upgrade"`
	// Allow upgrades for a specific package, ignoring pinned versions in any existing output file.
	//
	// Accepts both standalone package names (`ruff`) and version specifiers (`ruff<0.5.0`).
	// Experimental.
	UpgradePackage *[]*string `field:"optional" json:"upgradePackage" yaml:"upgradePackage"`
	// The workspace definition for the project, if any.
	// Experimental.
	Workspace *ToolUvWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

