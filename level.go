package ebitenLDTK

import (
	"fmt"
)

type Level struct {
	Name       string  `json:"identifier"`
	Iid        string  `json:"iid"`
	Uid        int     `json:"uid"`
	WorldX     int     `json:"worldX"`
	WorldY     int     `json:"worldY"`
	WorldDepth int     `json:"worldDepth"`
	PxWid      float64 `json:"pxWid"`
	PxHei      float64 `json:"pxHei"`
	Layers     []Layer `json:"layerInstances"`
}

// TODO: use intgrid data to make the bitmap more flexible
// Generate a 2D-array of ints indicating whether
// There is no point to this...
// Unless we just have a simple tilemap
// func (l *Level) MakeBitmapFromLayer(layerName string, ignoreIDs []int) [][]int {
// 	layer, err := l.GetLayerByName(layerName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	numTilesX := int(l.PxWid / layer.GridSize)
// 	numTilesY := int(l.PxHei / layer.GridSize)
// 	bitmap := make([][]int, numTilesY)
// 	for i := range bitmap {
// 		bitmap[i] = make([]int, numTilesX)
// 	}

// 	if layer.Type == LayerTypeIntGrid {
// 		// make 2D array from 1D array
// 		for i := range
// 		return layer.IntGridCsv
// 		// for _, tile := range layer.IntGridCsv {
// 		// 	posX := tile.Px[0] / layer.GridSize
// 		// 	posY := tile.Px[1] / layer.GridSize
// 		// 	bitmap[int(posY)][int(posX)] = 1
// 		// }
// 	} else if layer.Type == LayerTypeTiles {
// 		for _, tile := range layer.AutoLayerTiles {
// 			posX := tile.Px[0] / layer.GridSize
// 			posY := tile.Px[1] / layer.GridSize
// 			bitmap[int(posY)][int(posX)] = 1
// 		}
// 	}

// 	return bitmap
// }

func (l *Level) GetLayerByName(name string) (Layer, error) {
	for _, layer := range l.Layers {
		if layer.Name == name {
			return layer, nil
		}
	}
	return Layer{}, fmt.Errorf("layer with name [%s] was not found", name)
}

func (l *Level) GetEntityByIid(iid string) (Entity, error) {
	for _, layer := range l.Layers {
		for _, entity := range layer.Entities {
			if entity.Iid == iid {
				return entity, nil
			}
		}

	}
	return Entity{}, fmt.Errorf("entity with iid [%s] was not found", iid)
}
