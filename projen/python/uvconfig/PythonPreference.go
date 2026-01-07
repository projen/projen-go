package uvconfig


// Experimental.
type PythonPreference string

const (
	// Only use managed Python installations;
	//
	// never use system Python installations. (only-managed)
	// Experimental.
	PythonPreference_ONLY_HYPHEN_MANAGED PythonPreference = "ONLY_HYPHEN_MANAGED"
	// Prefer managed Python installations over system Python installations.
	//
	// System Python installations are still preferred over downloading managed Python versions.
	// Use `only-managed` to always fetch a managed Python version. (managed)
	// Experimental.
	PythonPreference_MANAGED PythonPreference = "MANAGED"
	// Prefer system Python installations over managed Python installations.
	//
	// If a system Python installation cannot be found, a managed Python installation can be used. (system)
	// Experimental.
	PythonPreference_SYSTEM PythonPreference = "SYSTEM"
	// Only use system Python installations;
	//
	// never use managed Python installations. (only-system)
	// Experimental.
	PythonPreference_ONLY_HYPHEN_SYSTEM PythonPreference = "ONLY_HYPHEN_SYSTEM"
)

