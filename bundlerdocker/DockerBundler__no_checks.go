//go:build no_runtime_type_checking

package bundlerdocker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerBundler) validateBundleParameters(options *cdktn.BundleOptions) error {
	return nil
}

func validateNewDockerBundlerParameters(props *DockerBundlerProps) error {
	return nil
}

