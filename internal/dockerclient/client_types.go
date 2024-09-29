package dockerclient

// ContainerInfo struct to contain info returned by listContainers
type ContainerInfo struct {
	ID      string
	Names   []string
	Image   string
	ImageID string
	Command string
	Created int64
	Ports   []string
	IP      string
	Labels  map[string]string
	State   string
	Status  string
}

// ImageInfo struct to contain info returned by listImages
type ImageInfo struct {
	Id        string
	RepoTags  []string
	Size      int64
	CreatedAt int64
}
