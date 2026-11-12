package bundlerdocker


// Options for {@link DockerImage.run }.
type DockerRunOptions struct {
	// The command to run in the container, as an argv array.
	// Default: - the image's default command.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Override the image entrypoint.
	// Default: - the image's default entrypoint.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Environment variables to set in the container.
	// Default: - no extra environment.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Named docker volumes to attach to the container.
	// Default: - no named volumes.
	//
	NamedVolumes *[]*DockerNamedVolume `field:"optional" json:"namedVolumes" yaml:"namedVolumes"`
	// The container network to attach to (`--network`).
	// Default: - the default docker network.
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The platform to run the container as (`--platform`).
	// Default: - the host platform.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// A `--security-opt` value passed to the container.
	// Default: - none.
	//
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to run the container as (`-u`).
	// Default: - the image's default user.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// Bind mounts to attach to the container.
	// Default: - no bind mounts.
	//
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// The working directory inside the container (`-w`).
	// Default: - the image's default working directory.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

