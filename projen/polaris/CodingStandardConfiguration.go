package polaris


// Experimental.
type CodingStandardConfiguration struct {
	// Enables AUTOSAR code compliance checking according to the given configuration.
	// Experimental.
	Autosarcpp14 *SpecificCodingStandardConfiguration `field:"optional" json:"autosarcpp14" yaml:"autosarcpp14"`
	// Enables CERT-C code compliance checking according to the given configuration.
	// Experimental.
	CertC *SpecificCodingStandardConfiguration `field:"optional" json:"certC" yaml:"certC"`
	// Enables CERT-CPP code compliance checking according to the given configuration.
	// Experimental.
	CertCpp *SpecificCodingStandardConfiguration `field:"optional" json:"certCpp" yaml:"certCpp"`
	// Enables CERT-C Recommendation code compliance checking according to the given configuration.
	// Experimental.
	CertCRecommendation *SpecificCodingStandardConfiguration `field:"optional" json:"certCRecommendation" yaml:"certCRecommendation"`
	// Enables CERT-Java code compliance checking according to the given configuration.
	// Experimental.
	CertJava *SpecificCodingStandardConfiguration `field:"optional" json:"certJava" yaml:"certJava"`
	// Enables HYUNDAI-C code compliance checking according to the given configuration.
	// Experimental.
	HyundaiC *SpecificCodingStandardConfiguration `field:"optional" json:"hyundaiC" yaml:"hyundaiC"`
	// Enables HYUNDAI-CPP code compliance checking according to the given configuration.
	// Experimental.
	HyundaiCpp *SpecificCodingStandardConfiguration `field:"optional" json:"hyundaiCpp" yaml:"hyundaiCpp"`
	// Enables HYUNDAI-Java code compliance checking according to the given configuration.
	// Experimental.
	HyundaiJava *SpecificCodingStandardConfiguration `field:"optional" json:"hyundaiJava" yaml:"hyundaiJava"`
	// If set to true, any defects found in code annotated using the #pragma Coverity compliance directive will not be reported in Coverity Connect.
	//
	// Information about the defects that were suppressed can then be found in two files: deviations.txt deviations-warnings.txt
	// Experimental.
	IgnoreDeviatedFindings *bool `field:"optional" json:"ignoreDeviatedFindings" yaml:"ignoreDeviatedFindings"`
	// Enables ISO TS 17961 code compliance checking according to the given configuration.
	// Experimental.
	IsoTs17961 *SpecificCodingStandardConfiguration `field:"optional" json:"isoTs17961" yaml:"isoTs17961"`
	// Enables MISRA C 2004 code compliance checking according to the given configuration.
	// Experimental.
	Misrac2004 *SpecificCodingStandardConfiguration `field:"optional" json:"misrac2004" yaml:"misrac2004"`
	// Enables MISRA C 2012 code compliance checking according to the given configuration.
	// Experimental.
	Misrac2012 *SpecificCodingStandardConfiguration `field:"optional" json:"misrac2012" yaml:"misrac2012"`
	// Enables MISRA C 2023 code compliance checking according to the given configuration.
	// Experimental.
	Misrac2023 *SpecificCodingStandardConfiguration `field:"optional" json:"misrac2023" yaml:"misrac2023"`
	// Enables MISRA C++ 2008 code compliance checking according to the given configuration.
	// Experimental.
	Misracpp2008 *SpecificCodingStandardConfiguration `field:"optional" json:"misracpp2008" yaml:"misracpp2008"`
	// Enables MISRA C++ 2023 code compliance checking according to the given configuration.
	// Experimental.
	Misracpp2023 *SpecificCodingStandardConfiguration `field:"optional" json:"misracpp2023" yaml:"misracpp2023"`
}

