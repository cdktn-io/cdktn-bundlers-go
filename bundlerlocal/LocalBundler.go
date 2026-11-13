package bundlerlocal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-bundlers-go/bundlerlocal/jsii"

	"github.com/cdktn-io/cdktn-bundlers-go/bundlerlocal/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Runs a build command on the host and stages its output as a CDK for Terraform asset.
//
// The command is spawned with the asset source and a caller-owned output
// directory exposed through the `ASSET_INPUT_DIR` and `ASSET_OUTPUT_DIR`
// environment variables. Unlike a container bundler, the build runs directly
// on the machine performing the synth, so the required toolchain must already
// be installed there.
type LocalBundler interface {
	cdktn.IAssetBundler
	// A value identifying the build, folded into the asset hash.
	//
	// The source tree alone cannot see the build, so swapping a `node:18` base
	// image for `node:20` would otherwise leave identity unchanged. Under
	// `SOURCE` hashing this is the only channel by which the build reaches
	// identity, so it must serialize every input that can move the output, or a
	// changed build silently reuses a stale artifact. {@link BundlerKey} builds
	// one from an ordered set of parts.
	//
	// Mirrors {@link IIgnoreStrategy.cacheKey}: omit it when the build cannot be
	// summarized as a string, and fall back to `extraHash`.
	BundlerKey() *string
	// The name a single-file artifact is staged under (e.g. `archive.zip`).
	//
	// A file-producing bundler (`BundleResult.file`) staged with
	// `AssetType.FILE` would otherwise take the source path's basename, which is
	// a directory name when the source is a directory. Declaring the name here
	// lets the artifact reflect what the bundler produces. It is static
	// configuration, needed at construction before a deferred `SOURCE` build
	// runs, not the file's runtime name.
	//
	// Valid only for a file-producing bundler: setting it with a directory
	// packaging (anything but `AssetType.FILE`) is rejected at construction.
	OutputFileName() *string
	// Produce the artifact and return a {@link BundleResult} describing it.
	//
	// Implementations write into `options.outputDir` and never write back to
	// `options.source`. The returned path must exist and match its declared
	// shape, or staging rejects it.
	//
	// A bundler that cannot run in this environment returns
	// `BundleResult.declined()` so a {@link ChainBundler} can fall through to
	// the next; a thrown error is a hard failure, not a decline.
	Bundle(options *cdktn.BundleOptions) cdktn.BundleResult
}

// The jsii proxy struct for LocalBundler
type jsiiProxy_LocalBundler struct {
	internal.Type__cdktnIAssetBundler
}

func (j *jsiiProxy_LocalBundler) BundlerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundlerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBundler) OutputFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileName",
		&returns,
	)
	return returns
}


func NewLocalBundler(props *LocalBundlerProps) LocalBundler {
	_init_.Initialize()

	if err := validateNewLocalBundlerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LocalBundler{}

	_jsii_.Create(
		"@cdktn/bundler-local.LocalBundler",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLocalBundler_Override(l LocalBundler, props *LocalBundlerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/bundler-local.LocalBundler",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LocalBundler) Bundle(options *cdktn.BundleOptions) cdktn.BundleResult {
	if err := l.validateBundleParameters(options); err != nil {
		panic(err)
	}
	var returns cdktn.BundleResult

	_jsii_.Invoke(
		l,
		"bundle",
		[]interface{}{options},
		&returns,
	)

	return returns
}

