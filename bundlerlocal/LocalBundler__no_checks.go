//go:build no_runtime_type_checking

package bundlerlocal

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocalBundler) validateBundleParameters(options *cdktn.BundleOptions) error {
	return nil
}

func validateNewLocalBundlerParameters(props *LocalBundlerProps) error {
	return nil
}

