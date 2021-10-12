package q3bsp

import (
	"image/color"

	"github.com/g3n/engine/math32"
)

// BSP is a binary space partition
type BSP struct {
	header     bspHeader
	EntityInfo string
	Textures   []*Texture
	Planes     []*Plane
	Nodes      []*Node
	Leaves     []*Leaf
	LeafFaces  []*LeafFace
	// LeafBrushes stores lists of brush indices, with one list per leaf. There are a total of length / sizeof(leafbrush) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	LeafBrushes []*LeafBrush
	// Models describes rigid groups of world geometry. The first model correponds to the base portion of the map while the remaining models correspond to movable portions of the map, such as the map's doors, platforms, and buttons. Each model has a list of faces and list of brushes; these are especially important for the movable parts of the map, which (unlike the base portion of the map) do not have BSP trees associated with them. There are a total of length / sizeof(models) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	Models []*Model
	// Brushes stores a set of brushes, which are in turn used for collision detection. Each brush describes a convex volume as defined by its surrounding surfaces. There are a total of length / sizeof(brushes) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	Brushes []*Brush
	// BrushSides stores descriptions of brush bounding surfaces. There are a total of length / sizeof(brushsides) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	BrushSides []*BrushSide
	// Vertexes stores lists of vertex offsets, used to describe generalized triangle meshes. There are a total of length / sizeof(meshvert) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	Vertexes []*Vertex
	// MeshVertexOffsets stores lists of vertex offsets, used to describe generalized triangle meshes. There are a total of length / sizeof(meshvert) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	MeshVertexOffsets []*MeshVertexOffset
	// Effects stores references to volumetric shaders (typically fog) which affect the rendering of a particular group of faces. There are a total of length / sizeof(effect) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	Effects []*Effect
	// Faces stores information used to render the surfaces of the map. There are a total of length / sizeof(faces) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	Faces []*Face
	// Lightmaps stores the light map textures used make surface lighting look more realistic. There are a total of length / sizeof(lightmap) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	Lightmaps []*Lightmap
	// LightVolumes stores a uniform grid of lighting information used to illuminate non-map objects. There are a total of length / sizeof(lightvol) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
	LightVolumes []*LightVolume
	// VisInfo stores bit vectors that provide cluster-to-cluster visibility information. There is exactly one visdata record, with a length equal to that specified in the lump directory.
	VisInfo []*VisData
}

func New() *BSP {
	b := &BSP{}
	b.header.Header = [4]byte{0x49, 0x42, 0x53, 0x50}
	b.header.Version = 0x2E
	return b
}

type entry struct {
	Offset int32
	Size   int32
}

type bspHeader struct {
	Header  [4]byte
	Version int32
}

// Texture stores information about surfaces and volumes, which are in turn associated with faces, brushes, and brushsides. There are a total of length / sizeof(texture) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
type Texture struct {
	RawName      [64]byte
	Flags        int32
	ContentFlags int32
}

func (t *Texture) Name() string {
	return string(t.RawName[:])
}

// Plane referenced by nodes and brushsides. There are a total of length / sizeof(plane) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
type Plane struct {
	Normal math32.Vector3
	Dist   float32
}

// Node in the map's BSP tree. The BSP tree is used primarily as a spatial subdivision scheme, dividing the world into convex regions called leafs. The first node in the lump is the tree's root node. There are a total of length / sizeof(node) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
type Node struct {
	PlaneID  int32
	Children [2]int32
	Mins     [3]int32
	Maxs     [3]int32
}

// Leaf of the map's BSP tree. Each leaf is a convex region that contains, among other things, a cluster index (for determining the other leafs potentially visible from within the leaf), a list of faces (for rendering), and a list of brushes (for collision detection). There are a total of length / sizeof(leaf) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
type Leaf struct {
	// Visdata cluster index
	ClusterID int32
	Area      int32
}

// LeafFace stores leaf references, with one list per leaf. There are a total of length / sizeof(leafface) records in the lump, where length is the size of the lump itself, as specified in the lump directory. leafface
type LeafFace struct {
	FaceID int32
}

// LeafBrush has a brush reference
type LeafBrush struct {
	BrushID int32
}

// Model describes rigid groups of world geometry. The first model correponds to the base portion of the map while the remaining models correspond to movable portions of the map, such as the map's doors, platforms, and buttons. Each model has a list of faces and list of brushes; these are especially important for the movable parts of the map, which (unlike the base portion of the map) do not have BSP trees associated with them. There are a total of length / sizeof(models) records in the lump, where length is the size of the lump itself, as specified in the lump directory.
type Model struct {
	// Bounding box min coord.
	Mins [3]int32
	// 	Bounding box max coord.
	Maxs [3]int32
	// First face for model.
	FaceID     int32
	FaceCount  int32
	BrushID    int32
	BrushCount int32
}

// Brush is used for collision detection
type Brush struct {
	BrushSide      int32
	BrushSideCount int32
	TextureID      int32
}

// BrushSide is used for brush bounding surface info
type BrushSide struct {
	PlaneID   int32
	TextureID int32
}

type Vertex struct {
	Position math32.Vector3
	// TexCoords stores texture coordinates. 0=surface, 1=lightmap.
	TexCoords [2][2]float32
	Normal    math32.Vector3
	Color     color.RGBA
}

type MeshVertexOffset struct {
	OffsetID int32
}

type Effect struct {
	RawName [64]byte
	BrushID int32
	// always 5
	Unknown int32
}

func (e *Effect) Name() string {
	return string(e.RawName[:])
}

type Face struct {
	TextureID int32
	EffectID  int32
	//1=polygon, 2=patch, 3=mesh, 4=billboard
	TypeID int32
	// index of first vertex
	VertexID        int32
	VertexCount     int32
	MeshVertexID    int32
	MeshVertexCount int32
	LightMapID      int32
	LightMapStart   [2]int32
	LightMapSize    [2]int32
	LightMapOrigin  math32.Vector3
	LightMapVectors [2]math32.Vector3
	Normal          math32.Vector3
	Size            math32.Vector2
}

type Lightmap struct {
	Colors [128][128][3]uint8
}

type LightVolume struct {
	Ambient     [3]uint8
	Directional [3]uint8
	// 0=phi, 1=theta
	Direction [2]uint8
}

type VisData struct {
	VectorCount int32
	VectorSize  int32
	Vectors     []uint8
}

const (
	//	Game-related object descriptions.
	dirEntryEntities = 0
	//	Surface descriptions.
	dirEntryTextures = 1
	//	Planes used by map geometry.
	dirEntryPlanes = 2
	//	BSP tree nodes.
	dirEntryNodes = 3
	//	BSP tree leaves.
	dirEntryLeafs = 4
	//	Lists of face indices, one list per leaf.
	dirEntryLeaffaces = 5
	//	Lists of brush indices, one list per leaf.
	dirEntryLeafbrushes = 6
	//	Descriptions of rigid world geometry in map.
	dirEntryModels = 7
	//	Convex polyhedra used to describe solid space.
	dirEntryBrushes = 8
	//	Brush surfaces.
	dirEntryBrushsides = 9
	//	Vertices used to describe faces.
	dirEntryVertexes = 10
	//	Lists of offsets, one list per mesh.
	dirEntryMeshverts = 11
	//	List of special map effects.
	dirEntryEffects = 12
	//	Surface geometry.
	dirEntryFaces = 13
	//	Packed lightmap data.
	dirEntryLightmaps = 14
	//	Local illumination data.
	dirEntryLightvols = 15
	//	Cluster-cluster visibility data.
	dirEntryVisdata = 16
)

const (
	//dirEntryEntitiesSize    = 0
	dirEntryTexturesSize    = 76
	dirEntryPlanesSize      = 16
	dirEntryNodesSize       = 36
	dirEntryLeafsSize       = 8
	dirEntryLeaffacesSize   = 4
	dirEntryLeafbrushesSize = 4
	dirEntryModelsSize      = 40
	dirEntryBrushesSize     = 12
	dirEntryBrushsidesSize  = 44
	dirEntryVertexesSize    = 44
	dirEntryMeshvertsSize   = 4
	dirEntryEffectsSize     = 72
	dirEntryFacesSize       = 108
	dirEntryLightmapsSize   = 49152
	dirEntryLightvolsSize   = 8
	dirEntryVisdataSize     = 8
)
