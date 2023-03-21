package typescript


// Experimental.
type ProjenrcTsOptions struct {
	// The name of the projenrc file.
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// A directory tree that may contain *.ts files that can be referenced from your projenrc typescript file.
	// Experimental.
	ProjenCodeDir *string `field:"optional" json:"projenCodeDir" yaml:"projenCodeDir"`
	// The name of the tsconfig file that will be used by ts-node when compiling projen source files.
	// Experimental.
	TsconfigFileName *string `field:"optional" json:"tsconfigFileName" yaml:"tsconfigFileName"`
}

