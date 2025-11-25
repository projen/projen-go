package uvconfig


// Settings for the uv build backend (`uv_build`).
//
// Note that those settings only apply when using the `uv_build` backend, other build backends
// (such as hatchling) have their own configuration.
//
// All options that accept globs use the portable glob patterns from
// [PEP 639](https://packaging.python.org/en/latest/specifications/glob-patterns/).
// Experimental.
type BuildBackendSettings struct {
	// Data includes for wheels.
	//
	// Each entry is a directory, whose contents are copied to the matching directory in the wheel
	// in `<name>-<version>.data/(purelib|platlib|headers|scripts|data)`. Upon installation, this
	// data is moved to its target location, as defined by
	// <https://docs.python.org/3.12/library/sysconfig.html#installation-paths>. Usually, small
	// data files are included by placing them in the Python module instead of using data includes.
	//
	// - `scripts`: Installed to the directory for executables, `<venv>/bin` on Unix or
	// `<venv>\Scripts` on Windows. This directory is added to `PATH` when the virtual
	// environment  is activated or when using `uv run`, so this data type can be used to install
	// additional binaries. Consider using `project.scripts` instead for Python entrypoints.
	// - `data`: Installed over the virtualenv environment root.
	//
	// Warning: This may override existing files!
	//
	// - `headers`: Installed to the include directory. Compilers building Python packages
	// with this package as build requirement use the include directory to find additional header
	// files.
	// - `purelib` and `platlib`: Installed to the `site-packages` directory. It is not recommended
	// to use these two options.
	// Experimental.
	Data *WheelDataIncludes `field:"optional" json:"data" yaml:"data"`
	// If set to `false`, the default excludes aren't applied.
	//
	// Default excludes: `__pycache__`, `*.pyc`, and `*.pyo`.
	// Experimental.
	DefaultExcludes *bool `field:"optional" json:"defaultExcludes" yaml:"defaultExcludes"`
	// The name of the module directory inside `module-root`.
	//
	// The default module name is the package name with dots and dashes replaced by underscores.
	//
	// Package names need to be valid Python identifiers, and the directory needs to contain a
	// `__init__.py`. An exception are stubs packages, whose name ends with `-stubs`, with the stem
	// being the module name, and which contain a `__init__.pyi` file.
	//
	// For namespace packages with a single module, the path can be dotted, e.g., `foo.bar` or
	// `foo-stubs.bar`.
	//
	// For namespace packages with multiple modules, the path can be a list, e.g.,
	// `["foo", "bar"]`. We recommend using a single module per package, splitting multiple
	// packages into a workspace.
	//
	// Note that using this option runs the risk of creating two packages with different names but
	// the same module names. Installing such packages together leads to unspecified behavior,
	// often with corrupted files or directory trees.
	// Experimental.
	ModuleName interface{} `field:"optional" json:"moduleName" yaml:"moduleName"`
	// The directory that contains the module directory.
	//
	// Common values are `src` (src layout, the default) or an empty path (flat layout).
	// Experimental.
	ModuleRoot *string `field:"optional" json:"moduleRoot" yaml:"moduleRoot"`
	// Build a namespace package.
	//
	// Build a PEP 420 implicit namespace package, allowing more than one root `__init__.py`.
	//
	// Use this option when the namespace package contains multiple root `__init__.py`, for
	// namespace packages with a single root `__init__.py` use a dotted `module-name` instead.
	//
	// To compare dotted `module-name` and `namespace = true`, the first example below can be
	// expressed with `module-name = "cloud.database"`: There is one root `__init__.py` `database`.
	// In the second example, we have three roots (`cloud.database`, `cloud.database_pro`,
	// `billing.modules.database_pro`), so `namespace = true` is required.
	//
	// ```text
	// src
	// └── cloud
	// └── database
	// ├── __init__.py
	// ├── query_builder
	// │   └── __init__.py
	// └── sql
	// ├── parser.py
	// └── __init__.py
	// ```
	//
	// ```text
	// src
	// ├── cloud
	// │   ├── database
	// │   │   ├── __init__.py
	// │   │   ├── query_builder
	// │   │   │   └── __init__.py
	// │   │   └── sql
	// │   │       ├── __init__.py
	// │   │       └── parser.py
	// │   └── database_pro
	// │       ├── __init__.py
	// │       └── query_builder.py
	// └── billing
	// └── modules
	// └── database_pro
	// ├── __init__.py
	// └── sql.py
	// ```.
	// Experimental.
	Namespace *bool `field:"optional" json:"namespace" yaml:"namespace"`
	// Glob expressions which files and directories to exclude from the source distribution.
	// Experimental.
	SourceExclude *[]*string `field:"optional" json:"sourceExclude" yaml:"sourceExclude"`
	// Glob expressions which files and directories to additionally include in the source distribution.
	//
	// `pyproject.toml` and the contents of the module directory are always included.
	// Experimental.
	SourceInclude *[]*string `field:"optional" json:"sourceInclude" yaml:"sourceInclude"`
	// Glob expressions which files and directories to exclude from the wheel.
	// Experimental.
	WheelExclude *[]*string `field:"optional" json:"wheelExclude" yaml:"wheelExclude"`
}

