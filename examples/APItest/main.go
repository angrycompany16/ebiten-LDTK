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
				testFields(entity, level)
			}

			if entity.Name == "EntityFieldsTest" {
				field, _ := entity.GetFieldByName("Point")
				fmt.Println(field.Point)
				field, _ = entity.GetFieldByName("String_singleLine")
				fmt.Println(field.String)
			}

			if entity.Name == "EntityRefTest" {
				entityField, err := entity.GetFieldByName("target")
				if err != nil {
					panic(err)
				}

				if entityField.EntityRef.EntityIid == "" {
					fmt.Println("I have no entity ref, skipping...")
					continue
				}
				entity, err := level.GetEntityByIid(entityField.EntityRef.EntityIid)
				if err != nil {
					panic(err)
				}
				fmt.Println("This is my friend", entity.Name)
			}
		}
	}

	// Test int grid value
	// var ignoreValue int
	// for _, layerDef := range world.Defs.LayerDefs {
	// 	if layerDef.Name != "IntGrid_with_rules" {
	// 		continue
	// 	}

	// 	ignoreValue = layerDef.GetIntGridValue("walls")
	// }

	// var intGridCSV [][]int
	// for _, layer := range level.Layers {
	// 	if layer.Name != "IntGrid_with_rules" {
	// 		continue
	// 	}

	// 	intGridCSV = layer.ExtractLayerCSV([]int{ignoreValue})
	// }

	// for i := range intGridCSV {
	// 	fmt.Println(intGridCSV[i])
	// }
	fmt.Println("All tests passed successfully :)")
}

func testFields(entity ebitenLDTK.Entity, level ebitenLDTK.Level) {
	entityField, err := entity.GetFieldByName("Entity_refs")
	if err != nil {
		panic(err)
	}
	for _, entityRef := range entityField.EntityRefArray {
		fmt.Printf("I am entity %s, this is my friend entity %s\n", entity.Iid, entityRef.EntityIid)
	}

	levelField, err := level.GetFieldByName("Biome")
	if err != nil {
		panic(err)
	}
	fmt.Println("This level is in biome", levelField.Biome)
}
