package github


// Experimental.
type UploadArtifactWith struct {
	// A file, directory or wildcard pattern that describes what to upload.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The level of compression for Zlib to be applied to the artifact archive.
	//
	// The value can range from 0 to 9.
	// For large files that are not easily compressed, a value of 0 is recommended for significantly faster uploads.
	// Default: 6.
	//
	// Experimental.
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// The desired behavior if no files are found using the provided path.
	//
	// Available Options:
	//   warn: Output a warning but do not fail the action
	//   error: Fail the action with an error message
	// ignore: Do not output any warnings or errors, the action does not fail.
	// Default: "warn".
	//
	// Experimental.
	IfNoFilesFound *string `field:"optional" json:"ifNoFilesFound" yaml:"ifNoFilesFound"`
	// Name of the artifact to upload.
	// Default: "artifact".
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether action should overwrite an existing artifact with the same name (should one exist).
	//
	// Introduced in v4 and represents a breaking change from the behavior of the v3 action.
	// To maintain backwards compatibility with existing, this should be set the `true` (the default).
	// Default: true.
	//
	// Experimental.
	Overwrite *bool `field:"optional" json:"overwrite" yaml:"overwrite"`
	// Duration after which artifact will expire in days. 0 means using default repository retention.
	//
	// Minimum 1 day.
	// Maximum 90 days unless changed from the repository settings page.
	// Default: - The default repository retention.
	//
	// Experimental.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

