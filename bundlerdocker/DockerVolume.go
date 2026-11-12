package bundlerdocker


// A single docker container filesystem mount.
type DockerVolume struct {
	// Absolute path inside the container.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Absolute path on the host.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// Mount the path read-only.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

