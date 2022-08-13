package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
)

// Declares a Tailwind CSS configuration file.
//
// There are multiple ways to add Tailwind CSS in your node project - see:
// https://tailwindcss.com/docs/installation
// See: PostCss.
//
// Experimental.
type TailwindConfig interface {
	// Experimental.
	File() projen.JsonFile
	// Experimental.
	FileName() *string
}

// The jsii proxy struct for TailwindConfig
type jsiiProxy_TailwindConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_TailwindConfig) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TailwindConfig) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTailwindConfig(project javascript.NodeProject, options *TailwindConfigOptions) TailwindConfig {
	_init_.Initialize()

	j := jsiiProxy_TailwindConfig{}

	_jsii_.Create(
		"projen.web.TailwindConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTailwindConfig_Override(t TailwindConfig, project javascript.NodeProject, options *TailwindConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.TailwindConfig",
		[]interface{}{project, options},
		t,
	)
}

