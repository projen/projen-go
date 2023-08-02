package projen


// Configure options for a task to be run when opening a Gitpod workspace (e.g. running tests, or starting a dev server).
//
// Start Mode         | Execution
// Fresh Workspace    | before && init && command
// Restart Workspace  | before && command
// Snapshot           | before && command
// Prebuild           | before && init && prebuild.
// Experimental.
type GitpodTask struct {
	// Required.
	//
	// The shell command to run.
	// Experimental.
	Command *string `field:"required" json:"command" yaml:"command"`
	// In case you need to run something even before init, that is a requirement for both init and command, you can use the before property.
	// Experimental.
	Before *string `field:"optional" json:"before" yaml:"before"`
	// The init property can be used to specify shell commands that should only be executed after a workspace was freshly cloned and needs to be initialized somehow.
	//
	// Such tasks are usually builds or downloading
	// dependencies. Anything you only want to do once but not when you restart a workspace or start a snapshot.
	// Experimental.
	Init *string `field:"optional" json:"init" yaml:"init"`
	// A name for this task.
	// Default: - task names are omitted when blank.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// You can configure where in the IDE the terminal should be opened.
	// Default: GitpodOpenIn.BOTTOM
	//
	// Experimental.
	OpenIn GitpodOpenIn `field:"optional" json:"openIn" yaml:"openIn"`
	// You can configure how the terminal should be opened relative to the previous task.
	// Default: GitpodOpenMode.TAB_AFTER
	//
	// Experimental.
	OpenMode GitpodOpenMode `field:"optional" json:"openMode" yaml:"openMode"`
	// The optional prebuild command will be executed during prebuilds.
	//
	// It is meant to run additional long running
	// processes that could be useful, e.g. running test suites.
	// Experimental.
	Prebuild *string `field:"optional" json:"prebuild" yaml:"prebuild"`
}

