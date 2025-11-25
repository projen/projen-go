package python


// There are two kinds of metadata: _static_ and _dynamic_.
//
// - Static metadata is listed in the `[project]` table directly and cannot be specified or changed by a tool.
// - Dynamic metadata key names are listed inside the `dynamic` key and represents metadata that a tool will later provide.
// Experimental.
type PyprojectTomlProject struct {
	// Valid name consists only of ASCII letters and numbers, period, underscore and hyphen.
	//
	// It must start and end with a letter or number.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// People or organizations considered as 'authors' of the project.
	//
	// Each author is a table with `name` key, `email` key, or both.
	// Experimental.
	Authors *[]*ProjectAuthor `field:"optional" json:"authors" yaml:"authors"`
	// List of [Trove classifiers](https://pypi.org/classifiers/) that describe the project. PyPI use the classifiers to categorize projects.
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// An array of [dependency specifier](https://packaging.python.org/en/latest/specifications/dependency-specifiers/) strings, each representing a mandatory dependent package of the project.
	// Experimental.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Summary description of the project in one line.
	//
	// Tools may not accept multiple lines.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies which keys are intentionally unspecified under `[project]` table so build backend can/will provide such metadata dynamically.
	//
	// Each key must be listed only once. It is an error to both list a key in `dynamic` and use the key directly in `[project]`.
	// One of the most common usage is `version`, which allows build backend to retrieve project version from source code or version control system instead of hardcoding it in `pyproject.toml`.
	// Experimental.
	Dynamic *[]PyprojectTomlProjectDynamic `field:"optional" json:"dynamic" yaml:"dynamic"`
	// Extra [entry point groups](https://packaging.python.org/en/latest/specifications/entry-points/) that allow applications to load plugins. For example, Pygments (a syntax highlighting tool) can use additional styles from separately installed packages through `[project.entry-points."pygments.styles"]`. Each key is the name of the entry-point group, and each value is a table of entry points.
	// Experimental.
	EntryPoints interface{} `field:"optional" json:"entryPoints" yaml:"entryPoints"`
	// Table of [entry points](https://packaging.python.org/en/latest/specifications/entry-points/) that allows package installers to create a GUI wrapper for. Each key is the name of the script to be created, and each value is the function or object to all, in form of either `importable.module` or `importable.module:object.attr`. Windows platform treats `gui_scripts` specially in that they are wrapped in a GUI executable, so they can be started without a console, but cannot use standard streams unless application code redirects them.
	// Experimental.
	GuiScripts *map[string]*string `field:"optional" json:"guiScripts" yaml:"guiScripts"`
	// An array of strings specifying the import names that the project exclusively provides when installed.
	// Experimental.
	ImportNames *[]*string `field:"optional" json:"importNames" yaml:"importNames"`
	// An array of strings specifying the import names that the project provides when installed, but not exclusively.
	// Experimental.
	ImportNamespaces *[]*string `field:"optional" json:"importNamespaces" yaml:"importNamespaces"`
	// List of keywords or tags that describe the project.
	//
	// They could be used by search engines to categorize the project.
	// Experimental.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// For now it is a table with either: - `file` key specifying a relative path to a license file, or - `text` key containing full license content.
	//
	// Newer tool may accept a single [SPDX license expression](https://spdx.github.io/spdx-spec/v2.2.2/SPDX-license-expressions/) string instead of a table.
	// Experimental.
	License interface{} `field:"optional" json:"license" yaml:"license"`
	// Relative paths or globs to paths of license files.
	//
	// Can be an empty list.
	// Experimental.
	LicenseFiles *[]*string `field:"optional" json:"licenseFiles" yaml:"licenseFiles"`
	// People or organizations considered as 'maintainers' of the project.
	//
	// Each maintainer is a table with `name` key, `email` key, or both.
	// Experimental.
	Maintainers *[]*ProjectAuthor `field:"optional" json:"maintainers" yaml:"maintainers"`
	// Each entry is a key/value pair, with the key specifying [extra feature name](https://packaging.python.org/en/latest/specifications/core-metadata/#provides-extra-multiple-use) (such as `socks` in `requests[socks]`), and value is an array of [dependency specifier](https://packaging.python.org/en/latest/specifications/dependency-specifiers/) strings.
	// Experimental.
	OptionalDependencies interface{} `field:"optional" json:"optionalDependencies" yaml:"optionalDependencies"`
	// Value can be a relative path to text / markdown (`.md` suffix) / reStructuredText (`.rst` suffix) readme file, or a table with either: - `file` key containing path of aforementioned readme file, or - `text` key containing the full readme text embedded inside `pyproject.toml`.
	// Experimental.
	Readme interface{} `field:"optional" json:"readme" yaml:"readme"`
	// Specifies the Python version(s) that the distribution is compatible with.
	//
	// Must be in the format specified in [Version specifiers](https://packaging.python.org/en/latest/specifications/version-specifiers/).
	// Experimental.
	RequiresPython *string `field:"optional" json:"requiresPython" yaml:"requiresPython"`
	// Table of [entry points](https://packaging.python.org/en/latest/specifications/entry-points/) that allows package installers to create a command-line wrapper for. Each key is the name of the script to be created, and each value is the function or object to all, in form of either `importable.module` or `importable.module:object.attr`. Windows platform treats `console_scripts` specially in that they are wrapped in a console executable, so they are attached to a console and can use `sys.stdin`, `sys.stdout` and `sys.stderr` for I/O.
	// Experimental.
	Scripts *map[string]*string `field:"optional" json:"scripts" yaml:"scripts"`
	// Table consisting one or multiple `label: URL` pairs.
	//
	// Common indexes like PyPI uses [well-known Project URLs](https://packaging.python.org/en/latest/specifications/well-known-project-urls/#well-known-labels) when presenting project pages.
	// Experimental.
	Urls *map[string]*string `field:"optional" json:"urls" yaml:"urls"`
	// Version of the project, as defined in the [Version specifier specification](https://packaging.python.org/en/latest/specifications/version-specifiers/), and preferably [already normalized](https://packaging.python.org/en/latest/specifications/version-specifiers/#normalization).
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

