package python


// Every tool that is used by the project can have users specify configuration data as long as they use a sub-table within `[tool]`.
//
// Generally a project can use the subtable `tool.$NAME` if, and only if, they own the entry for `$NAME` in the Cheeseshop/PyPI.
// Experimental.
type PyprojectTomlTool struct {
	// The uncompromising Python code formatter.
	// Experimental.
	Black interface{} `field:"optional" json:"black" yaml:"black"`
	// Build Python wheels for all platforms.
	// Experimental.
	Cibuildwheel interface{} `field:"optional" json:"cibuildwheel" yaml:"cibuildwheel"`
	// A CLI tool to check and validate Python docstring formatting and completeness.
	// Experimental.
	Dfc interface{} `field:"optional" json:"dfc" yaml:"dfc"`
	// A CLI tool to check and validate Python docstring formatting and completeness.
	// Experimental.
	DocstringFormatChecker interface{} `field:"optional" json:"docstringFormatChecker" yaml:"docstringFormatChecker"`
	// FastAPI web framework configuration.
	// Experimental.
	Fastapi interface{} `field:"optional" json:"fastapi" yaml:"fastapi"`
	// Modern, extensible Python project management.
	// Experimental.
	Hatch interface{} `field:"optional" json:"hatch" yaml:"hatch"`
	// Build and publish crates with pyo3, cffi and uniffi bindings as well as rust binaries as python packages.
	// Experimental.
	Maturin interface{} `field:"optional" json:"maturin" yaml:"maturin"`
	// Optional static typing for Python.
	// Experimental.
	Mypy interface{} `field:"optional" json:"mypy" yaml:"mypy"`
	// A modern Python package manager with PEP 621 support.
	// Experimental.
	Pdm interface{} `field:"optional" json:"pdm" yaml:"pdm"`
	// A package manager and task runner.
	// Experimental.
	Pixi interface{} `field:"optional" json:"pixi" yaml:"pixi"`
	// A task runner that works well with `pyproject.toml` files.
	// Experimental.
	Poe interface{} `field:"optional" json:"poe" yaml:"poe"`
	// Python dependency management and packaging made easy.
	// Experimental.
	Poetry interface{} `field:"optional" json:"poetry" yaml:"poetry"`
	// Static type checker for Python.
	// Experimental.
	Pyright interface{} `field:"optional" json:"pyright" yaml:"pyright"`
	// Standardized automated testing of Python packages.
	// Experimental.
	Pytest interface{} `field:"optional" json:"pytest" yaml:"pytest"`
	// A CLI tool to run code files instantly without typing complex commands in terminal.
	// Experimental.
	Quikrun interface{} `field:"optional" json:"quikrun" yaml:"quikrun"`
	// Review a repository for best practices.
	// Experimental.
	RepoReview interface{} `field:"optional" json:"repoReview" yaml:"repoReview"`
	// An extremely fast Python linter and formatter, written in Rust.
	// Experimental.
	Ruff interface{} `field:"optional" json:"ruff" yaml:"ruff"`
	// Scheduled jobs in Python's `pyproject.toml`.
	//
	// This is a specification for declaring recurring scheduled jobs in Python projects, in `pyproject.toml`.
	//
	// It defines how jobs are declared and how providers would run them.
	//
	// It does not provide a specific implementation for running scheduled jobs, because that is provider specific.
	//
	// For example, a file at `app/jobs.py` could define:
	//
	// ```python
	// def clean_files():
	// print("Running cleanup...")
	// ```
	//
	// You could define a scheduled job to run that function once per day with:
	//
	// ```toml
	// [tool.scheduled.clean-files]
	// every = "day"
	// entrypoint = "app.jobs:clean_files"
	// ```.
	// Experimental.
	Scheduled interface{} `field:"optional" json:"scheduled" yaml:"scheduled"`
	// Improved build system generator for Python C/C++/Fortran extensions.
	// Experimental.
	ScikitBuild interface{} `field:"optional" json:"scikitBuild" yaml:"scikitBuild"`
	// Easily download, build, install, upgrade, and uninstall Python packages.
	// Experimental.
	Setuptools interface{} `field:"optional" json:"setuptools" yaml:"setuptools"`
	// Manage Python package versions using SCM (e.g. Git).
	// Experimental.
	SetuptoolsScm interface{} `field:"optional" json:"setuptoolsScm" yaml:"setuptoolsScm"`
	// The complementary task runner for python.
	// Experimental.
	Taskipy interface{} `field:"optional" json:"taskipy" yaml:"taskipy"`
	// Tombi is a toolkit for TOML;
	//
	// providing a formatter/linter and language server.
	// Experimental.
	Tombi interface{} `field:"optional" json:"tombi" yaml:"tombi"`
	// Standardized automated testing of Python packages.
	// Experimental.
	Tox interface{} `field:"optional" json:"tox" yaml:"tox"`
	// An extremely fast Python type checker, written in Rust.
	// Experimental.
	Ty interface{} `field:"optional" json:"ty" yaml:"ty"`
	// An extremely fast Python package installer and resolver, written in Rust.
	// Experimental.
	Uv interface{} `field:"optional" json:"uv" yaml:"uv"`
}

