package main

import (
	"fmt"
	"log"
	"path/filepath"

	ebitenLDTK "github.com/angrycompany16/ebiten-LDTK"
)

func main() {
	path := filepath.Join("examples", "APItest", "APItest.ldtk")
	world, err := ebitenLDTK.LoadWorld(path)

	if err != nil {
		log.Fatal(err)
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
}

func testFields(entity ebitenLDTK.Entity) {
	field, err := entity.GetFieldByName("Entity_refs")
	if err != nil {
		log.Fatal(err)
	}
	for _, entityRef := range field.EntityRefArray {
		fmt.Printf("I am entity %s, this is my friend entity %s\n", entity.Iid, entityRef.EntityIid)
	}
}
