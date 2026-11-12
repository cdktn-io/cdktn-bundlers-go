package bundlerdocker


// A named docker volume mounted into a container.
type DockerNamedVolume struct {
	// Absolute path inside the container.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The volume name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Mount the volume read-only.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

