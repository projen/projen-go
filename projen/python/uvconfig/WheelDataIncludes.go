package uvconfig


// Data includes for wheels.
//
// See `BuildBackendSettings::data`.
// Experimental.
type WheelDataIncludes struct {
	// Experimental.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Experimental.
	Headers *string `field:"optional" json:"headers" yaml:"headers"`
	// Experimental.
	Platlib *string `field:"optional" json:"platlib" yaml:"platlib"`
	// Experimental.
	Purelib *string `field:"optional" json:"purelib" yaml:"purelib"`
	// Experimental.
	Scripts *string `field:"optional" json:"scripts" yaml:"scripts"`
}

