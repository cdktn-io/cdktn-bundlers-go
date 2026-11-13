package bundlerlocal

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Configuration for {@link LocalBundler}.
type LocalBundlerProps struct {
	// The command to run on the host, as an argv array.
	//
	// The first element is the program; the rest are its arguments. The command
	// runs with `ASSET_INPUT_DIR` set to the source and `ASSET_OUTPUT_DIR` set
	// to the output directory, and it is expected to write the built artifact
	// into `ASSET_OUTPUT_DIR`.
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// Environment variables to set for the command.
	// Default: - only the caller's environment plus `ASSET_INPUT_DIR` and
	// `ASSET_OUTPUT_DIR`.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Extra content folded into the bundler key so a change to it invalidates the asset hash even when the source and command are unchanged.
	// Default: - nothing extra is folded in.
	//
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// The file name to stage a single-file artifact under.
	//
	// Only meaningful when the result is a single file staged as
	// `AssetType.FILE`. Without it, a file produced from a directory source is
	// staged under the source directory's basename.
	// Default: - the basename of the asset source path.
	//
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// The shape of the artifact the command writes to `ASSET_OUTPUT_DIR`.
	//
	// `DIRECTORY` stages the output directory. `FILE` requires the command to
	// write exactly one file and stages that file.
	// Default: - `FILE` when `outputFileName` is set, otherwise `DIRECTORY`.
	//
	OutputType cdktn.BundleOutputType `field:"optional" json:"outputType" yaml:"outputType"`
	// The working directory the command runs in.
	// Default: - the asset source directory, or its parent for a file source.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

