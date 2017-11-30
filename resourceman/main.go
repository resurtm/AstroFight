package resourceman

type ResourceType byte

const (
	MeshResource    ResourceType = iota
	TextureResource
)

type Resource interface {
	Type() ResourceType
}

type Mesh struct {
}

func (m *Mesh) Type() ResourceType {
	return MeshResource
}

type Texture struct {
}

func (t *Texture) Type() ResourceType {
	return TextureResource
}

type ResourceMan struct {
}

func New() *ResourceMan {
	return &ResourceMan{}
}

func (rm *ResourceMan) Load(fileName string, resourceType ResourceType) (res Resource) {
	switch resourceType {
	case MeshResource:
		res = rm.loadMesh(fileName)
	case TextureResource:
		res = rm.loadTexxture(fileName)
	}
	return
}

func (rm *ResourceMan) loadMesh(fileName string) Resource {
	return &Mesh{}
}

func (rm *ResourceMan) loadTexxture(fileName string) Resource {
	return &Texture{}
}
