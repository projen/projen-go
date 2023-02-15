package vscode


// Options for a 'VsCodeLaunchConfigurationEntry' Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type VsCodeLaunchConfigurationEntry struct {
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	Request *string `field:"required" json:"request" yaml:"request"`
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Experimental.
	Console Console `field:"optional" json:"console" yaml:"console"`
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// Experimental.
	DebugServer *float64 `field:"optional" json:"debugServer" yaml:"debugServer"`
	// Experimental.
	DisableOptimisticBPs *bool `field:"optional" json:"disableOptimisticBPs" yaml:"disableOptimisticBPs"`
	// Set value to `false` to unset an existing environment variable.
	// Experimental.
	Env *map[string]interface{} `field:"optional" json:"env" yaml:"env"`
	// Experimental.
	EnvFile *string `field:"optional" json:"envFile" yaml:"envFile"`
	// Experimental.
	InternalConsoleOptions InternalConsoleOptions `field:"optional" json:"internalConsoleOptions" yaml:"internalConsoleOptions"`
	// Experimental.
	OutFiles *[]*string `field:"optional" json:"outFiles" yaml:"outFiles"`
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Experimental.
	PostDebugTask *string `field:"optional" json:"postDebugTask" yaml:"postDebugTask"`
	// Experimental.
	PreLaunchTask *string `field:"optional" json:"preLaunchTask" yaml:"preLaunchTask"`
	// Experimental.
	Presentation *Presentation `field:"optional" json:"presentation" yaml:"presentation"`
	// Experimental.
	Program *string `field:"optional" json:"program" yaml:"program"`
	// Experimental.
	RuntimeArgs *[]*string `field:"optional" json:"runtimeArgs" yaml:"runtimeArgs"`
	// Experimental.
	ServerReadyAction *ServerReadyAction `field:"optional" json:"serverReadyAction" yaml:"serverReadyAction"`
	// Experimental.
	SkipFiles *[]*string `field:"optional" json:"skipFiles" yaml:"skipFiles"`
	// Experimental.
	StopOnEntry *bool `field:"optional" json:"stopOnEntry" yaml:"stopOnEntry"`
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Experimental.
	WebRoot *string `field:"optional" json:"webRoot" yaml:"webRoot"`
}

