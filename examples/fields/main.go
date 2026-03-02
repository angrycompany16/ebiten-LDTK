package main

import (
	"fmt"
	"path/filepath"

	ebitenLDTK "github.com/angrycompany16/ebiten-LDTK"
)

const TestEntityName = "FieldsTest"

// Field names
var fieldNames = []string{
	"IntArray",
	"FloatArray",
	"BoolArray",
	"StringArray",
	"MultilinesArray",
	"ColorArray",
	"LocalEnumArray",
	"ExternEnumArray",
	"FilePathArray",
	"TileArray",
	"EntityRefArray",
	"PointArray",
}

func main() {
	path := filepath.Join("examples", "fields", "assets", "fields.ldtk")
	world, err := ebitenLDTK.LoadWorld(path)
	level, err := world.GetLevelByName("Level_0")
	entityLayer, err := level.GetLayerByName("Entities")

	if err != nil {
		panic(err)
	}

	for _, entity := range entityLayer.Entities {
		if entity.Name != TestEntityName {
			continue
		}
		field, _ := entity.GetFieldByName(fieldNames[11])
		// Nice. Looks like everything works
		fmt.Println(ebitenLDTK.AsArray[ebitenLDTK.Point](field))
	}
}
