package typescript


// Experimental.
type TsJestOptions struct {
	// Which files should ts-jest act upon.
	// See: https://jestjs.io/docs/configuration#transform-objectstring-pathtotransformer--pathtotransformer-object
	//
	// Default: "^.+\\.[t]sx?$"
	//
	// Experimental.
	TranformPattern *string `field:"optional" json:"tranformPattern" yaml:"tranformPattern"`
	// Override the default ts-jest transformer configuration.
	// Experimental.
	TransformOptions *TsJestTransformOptions `field:"optional" json:"transformOptions" yaml:"transformOptions"`
}

