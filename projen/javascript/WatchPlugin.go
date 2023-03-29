package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Experimental.
type WatchPlugin interface {
}

// The jsii proxy struct for WatchPlugin
type jsiiProxy_WatchPlugin struct {
	_ byte // padding
}

// Experimental.
func NewWatchPlugin(name *string, options interface{}) WatchPlugin {
	_init_.Initialize()

	if err := validateNewWatchPluginParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_WatchPlugin{}

	_jsii_.Create(
		"projen.javascript.WatchPlugin",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewWatchPlugin_Override(w WatchPlugin, name *string, options interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.WatchPlugin",
		[]interface{}{name, options},
		w,
	)
}

