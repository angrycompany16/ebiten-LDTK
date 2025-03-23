package main

import (
	"fmt"
	"path/filepath"

	ebitenLDTK "github.com/angrycompany16/ebiten-LDTK"
)

func main() {
	path := filepath.Join("examples", "APItest", "APItest.ldtk")
	world, err := ebitenLDTK.LoadWorld(path)

	if err != nil {
		panic(err)
	}

	// Test entity refs
	level := world.Levels[0]
	for _, layer := range level.Layers {
		for _, entity := range layer.Entities {
			if entity.Name == "EntityRefArrayTest" {
				testFields(entity)
			}
		}
	}

	// Test int grid value
	for _, layerDef := range world.Defs.LayerDefs {
		if layerDef.Name != "IntGrid_with_rules" {
			continue
		}

		value := layerDef.GetIntGridValue("walls")
		fmt.Println(value)
	}
}

func testFields(entity ebitenLDTK.Entity) {
	field, err := entity.GetFieldByName("Entity_refs")
	if err != nil {
		panic(err)
	}
	for _, entityRef := range field.EntityRefArray {
		fmt.Printf("I am entity %s, this is my friend entity %s\n", entity.Iid, entityRef.EntityIid)
	}
}

func getIntGridValueIdentifier(value int, world ebitenLDTK.World) string {
	for _, layerDef := range world.Defs.LayerDefs {
		for _, intGridValue := range layerDef.IntGridValues {
			if intGridValue.Value == value {
				return intGridValue.Identifier
			}
		}
	}
	return ""
}
