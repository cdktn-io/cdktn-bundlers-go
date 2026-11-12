package bundlerdocker

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Configuration for {@link DockerBundler}.
type DockerBundlerProps struct {
	// The command to run in the container, as an argv array.
	//
	// The command reads the source from `/asset-input` and writes the built
	// artifact to `/asset-output`.
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// The image to run the build in, as a registry reference.
	//
	// Locally built images are not supported; build and push the image first.
	Image *string `field:"required" json:"image" yaml:"image"`
	// How the source and output cross the host/container boundary.
	// Default: BundlingFileAccess.BIND_MOUNT
	//
	BundlingFileAccess BundlingFileAccess `field:"optional" json:"bundlingFileAccess" yaml:"bundlingFileAccess"`
	// Override the image entrypoint.
	// Default: - the image's default entrypoint.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Environment variables to set in the container.
	// Default: - no extra environment.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The container network to attach to.
	// Default: - the default docker network.
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The file name to stage a single-file artifact under.
	//
	// Only meaningful when the result is a single file staged as
	// `AssetType.FILE`.
	// Default: - the basename of the asset source path.
	//
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// The shape of the artifact the command writes to `/asset-output`.
	//
	// `DIRECTORY` stages the output directory. `FILE` requires the command to
	// write exactly one file and stages that file.
	// Default: - `FILE` when `outputFileName` is set, otherwise `DIRECTORY`.
	//
	OutputType cdktn.BundleOutputType `field:"optional" json:"outputType" yaml:"outputType"`
	// The platform to run the container as.
	// Default: - the host platform.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// A `--security-opt` value passed to the container.
	// Default: - none.
	//
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to run the container as.
	// Default: - the current host uid:gid.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// The image used for the short-lived helper container that seeds the input volume and fixes volume ownership under `VOLUME_COPY`.
	//
	// The build image is reused by default, so no second image is pulled and the
	// caller controls every reference the bundler touches. Override this only
	// when the build image has no POSIX shell or `chown` (e.g. a `scratch`- or
	// distroless-based image), pointing it at any small image that does.
	//
	// Ignored under `BIND_MOUNT`.
	// Default: - the build image.
	//
	VolumeCopyHelperImage *string `field:"optional" json:"volumeCopyHelperImage" yaml:"volumeCopyHelperImage"`
	// The working directory inside the container.
	// Default: "/asset-input".
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

