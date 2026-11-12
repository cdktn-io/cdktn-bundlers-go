package bundlerdocker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-bundlers-go/bundlerdocker/jsii"
)

// A docker image referenced from a registry.
type DockerImage interface {
	// The image reference (name and optional tag) this instance runs.
	Image() *string
	// Run the image once with the given options and wait for it to exit.
	Run(options *DockerRunOptions)
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_DockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Reference an image by name or by name and tag from a registry.
func DockerImage_FromRegistry(image *string) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromRegistryParameters(image); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"@cdktn/bundler-docker.DockerImage",
		"fromRegistry",
		[]interface{}{image},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Run(options *DockerRunOptions) {
	if err := d.validateRunParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"run",
		[]interface{}{options},
	)
}

