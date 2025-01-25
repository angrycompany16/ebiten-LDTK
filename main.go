package ebitenLDTK

import (
	"encoding/json"
	"os"
)

type LDTKWorld struct {
	GridWidth  int         `json:"worldGridWidth"`
	GridHeight int         `json:"worldGridHeight"`
	Defs       LDTKDefs    `json:"defs"`
	Levels     []LDTKLevel `json:"levels"`
}

type LDTKDefs struct {
	Layers        []LDTKLayer        `json:"layers"`
	Entities      []LDTKEntity       `json:"entities"`
	Tilesets      []LDTKTileset      `json:"tilesets"`
	Enums         []LDTKEnum         `json:"enums"`
	ExternalEnums []LDTKExternalEnum `json:"externalEnums"`
	LevelFields   []LDTKLevelField   `json:"levelFields"`
}

type LDTKLayer struct {
	GridSize        int  `json:"gridSize"`
	PxOffsetX       int  `json:"pxOffsetX"`
	PxOffsetY       int  `json:"pxOffsetY"`
	ParallaxFactorX int  `json:"parallaxFactorX"`
	ParallaxFactorY int  `json:"parallaxFactorY"`
	ParallaxScaling bool `json:"parallaxScaling"`
}

type LDTKEntity struct {
	// Fields...
}

type LDTKTileset struct {
	Identifier   string `json:"identifier"`
	RelPath      string `json:"relPath"`
	PxWid        int    `json:"pxWid"`
	PxHei        int    `json:"pxHei"`
	TileGridSize int    `json:"tileGridSize"`
	Spacing      int    `json:"spacing"`
	Padding      int    `json:"padding"`
}

type LDTKEnum struct {
	// Fields...
}

type LDTKExternalEnum struct {
	// Fields...
}

type LDTKLevelField struct {
	// Fields...
}

type LDTKLevel struct {
	Identifier     string              `json:"identifier"`
	WorldX         int                 `json:"worldX"`
	WorldY         int                 `json:"worldY"`
	WorldDepth     int                 `json:"worldDepth"`
	PxWid          int                 `json:"pxWid"`
	PxHei          int                 `json:"pxHei"`
	LayerInstances []LDTKLayerInstance `json:"layerInstances"`
}

type LDTKLayerInstance struct {
	Identifier string             `json:"__identifier"`
	PxOffsetX  int                `json:"pxOffsetX"`
	PxOffsetY  int                `json:"pxOffsetY"`
	GridTiles  []LDTKTileInstance `json:"gridTiles"`
}

type LDTKTileInstance struct {
	Px  []int   `json:"px"`
	Src []int   `json:"src"`
	F   int     `json:"f"`
	T   int     `json:"t"`
	D   []int   `json:"d"`
	A   float64 `json:"a"`
}

type LDTKEntityInstance struct {
	// Fields...
}

func LoadLDTK(path string) (LDTKWorld, error) {
	world := LDTKWorld{}
	file, err := os.Open(path)
	if err != nil {
		return world, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&world)
	if err != nil {
		return world, err
	}
	return world, nil
}

func (l *LDTKLevel) MakeBitmap(layer *LDTKLayer, layerInstance *LDTKLayerInstance) [][]int {
	// size = width / tilesize, height / tilesize
	numTilesX := l.PxWid / layer.GridSize
	numTilesY := l.PxHei / layer.GridSize
	bitmap := make([][]int, numTilesY)
	for i := range bitmap {
		bitmap[i] = make([]int, numTilesX)
	}
	return bitmap
}

// func main() {
// 	// testLDTKJSON("../mask_of_the_tomb/assets/LDTK/test.ldtk", "defs")
// 	world, err := LoadLDTK("../mask_of_the_tomb/assets/LDTK/test.ldtk")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// world, err := LoadLDTK("../mask_of_the_tomb/assets/LDTK/test.ldtk")
// 	// fmt.Printf("%+v\n", world)
// 	out, _ := json.MarshalIndent(world, "", "\t")
// 	fmt.Println(string(out))
// }
