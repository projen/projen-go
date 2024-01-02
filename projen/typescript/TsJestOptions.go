package typescript


// Experimental.
type TsJestOptions struct {
	// Override the default ts-jest transformer configuration.
	// Experimental.
	TransformOptions *TsJestTransformOptions `field:"optional" json:"transformOptions" yaml:"transformOptions"`
	// Which files should ts-jest act upon.
	// See: https://jestjs.io/docs/configuration#transform-objectstring-pathtotransformer--pathtotransformer-object
	//
	// Default: "^.+\\.[t]sx?$"
	//
	// Experimental.
	TransformPattern *string `field:"optional" json:"transformPattern" yaml:"transformPattern"`
}

