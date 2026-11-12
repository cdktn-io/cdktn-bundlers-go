package bundlerdocker


// How the source and output cross the host/container boundary.
type BundlingFileAccess string

const (
	// Bind-mount the host source and output directories into the container.
	//
	// The fastest option, but requires the host paths to be shareable with the
	// docker daemon.
	BundlingFileAccess_BIND_MOUNT BundlingFileAccess = "BIND_MOUNT"
	// Copy the source into a named volume, run the build, then copy the output back out.
	//
	// Works where bind mounts are unavailable (e.g. a remote docker daemon).
	BundlingFileAccess_VOLUME_COPY BundlingFileAccess = "VOLUME_COPY"
)

