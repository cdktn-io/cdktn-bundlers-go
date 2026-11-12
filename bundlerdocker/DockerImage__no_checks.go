//go:build no_runtime_type_checking

package bundlerdocker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerImage) validateRunParameters(options *DockerRunOptions) error {
	return nil
}

func validateDockerImage_FromRegistryParameters(image *string) error {
	return nil
}

