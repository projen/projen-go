package javascript


// Experimental.
type NodeConfigSchemaPermission struct {
	// allow use of addons when any permissions are set.
	// Experimental.
	AllowAddons *bool `field:"optional" json:"allowAddons" yaml:"allowAddons"`
	// allow use of child process when any permissions are set.
	// Experimental.
	AllowChildProcess *bool `field:"optional" json:"allowChildProcess" yaml:"allowChildProcess"`
	// allow use of FFI when any permissions are set.
	// Experimental.
	AllowFfi *bool `field:"optional" json:"allowFfi" yaml:"allowFfi"`
	// allow permissions to read the filesystem.
	// Experimental.
	AllowFsRead *[]*string `field:"optional" json:"allowFsRead" yaml:"allowFsRead"`
	// allow permissions to write in the filesystem.
	// Experimental.
	AllowFsWrite *[]*string `field:"optional" json:"allowFsWrite" yaml:"allowFsWrite"`
	// allow use of inspector when any permissions are set.
	// Experimental.
	AllowInspector *bool `field:"optional" json:"allowInspector" yaml:"allowInspector"`
	// allow use of network when any permissions are set.
	// Experimental.
	AllowNet *bool `field:"optional" json:"allowNet" yaml:"allowNet"`
	// allow use of OpenSSL STORE loaders when any permissions are set.
	// Experimental.
	AllowOpensslStore *bool `field:"optional" json:"allowOpensslStore" yaml:"allowOpensslStore"`
	// allow wasi when any permissions are set.
	// Experimental.
	AllowWasi *bool `field:"optional" json:"allowWasi" yaml:"allowWasi"`
	// allow worker threads when any permissions are set.
	// Experimental.
	AllowWorker *bool `field:"optional" json:"allowWorker" yaml:"allowWorker"`
	// enable the permission system.
	// Experimental.
	Permission *bool `field:"optional" json:"permission" yaml:"permission"`
}

