package web


// Experimental.
type ReactComponentOptions struct {
	// Rewire webpack configuration.
	//
	// Use this property to override webpack configuration properties provided
	// by create-react-app, without needing to eject.
	//
	// This property will create a `config-overrides.js` file in your root directory,
	// which will contain the desired rewiring code.
	//
	// To **override** the configuration, you can provide simple key value pairs.
	// Keys take the form of js code directives that traverse to the desired property.
	// Values should be JSON serializable objects.
	//
	// For example, the following config:
	//
	// ```json
	// rewire: { "module.unknownContextCritical": false }
	// ```
	//
	// Will translate to the following `config-overrides.js` file:
	//
	// ```js
	// module.exports = function override(config, env) {
	//   config.module.unknownContextCritical = false;
	// }
	// ```.
	// See: https://github.com/timarney/react-app-rewired
	//
	// Experimental.
	Rewire *map[string]interface{} `field:"optional" json:"rewire" yaml:"rewire"`
	// Whether to apply options specific for TypeScript React projects.
	// Experimental.
	Typescript *bool `field:"optional" json:"typescript" yaml:"typescript"`
}

