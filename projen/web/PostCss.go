package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
)

// Declares a PostCSS dependency with a default config file.
// Experimental.
type PostCss interface {
	// Experimental.
	File() projen.JsonFile
	// Experimental.
	FileName() *string
	// Experimental.
	Tailwind() TailwindConfig
}

// The jsii proxy struct for PostCss
type jsiiProxy_PostCss struct {
	_ byte // padding
}

func (j *jsiiProxy_PostCss) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostCss) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostCss) Tailwind() TailwindConfig {
	var returns TailwindConfig
	_jsii_.Get(
		j,
		"tailwind",
		&returns,
	)
	return returns
}


// Experimental.
func NewPostCss(project javascript.NodeProject, options *PostCssOptions) PostCss {
	_init_.Initialize()

	if err := validateNewPostCssParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostCss{}

	_jsii_.Create(
		"projen.web.PostCss",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPostCss_Override(p PostCss, project javascript.NodeProject, options *PostCssOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.PostCss",
		[]interface{}{project, options},
		p,
	)
}

