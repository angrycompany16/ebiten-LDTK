package ebitenLDTK

type LayerType string

const (
	LayerTypeTiles    = "Tiles"
	LayerTypeEntities = "Entities"
	LayerTypeIntGrid  = "IntGrid"
)

type Layer struct {
	Name           string    `json:"__identifier"`
	Type           LayerType `json:"__type"`
	Width          int       `json:"__cWid"`
	Height         int       `json:"__cHei"`
	GridSize       float64   `json:"__gridSize"`
	TilesetUid     int       `json:"__tilesetDefUid"`
	TilesetRelPath string    `json:"__tilesetRelPath"`
	PxOffsetX      int       `json:"pxOffsetX"`
	PxOffsetY      int       `json:"pxOffsetY"`
	IntGridCsv     []int     `json:"intGridCsv"`
	GridTiles      []Tile    `json:"gridTiles"`
	Entities       []Entity  `json:"entityInstances"`
	AutoLayerTiles []Tile    `json:"autoLayerTiles"`
}

type Tile struct {
	Px              []float64   `json:"px"`
	Src             []float64   `json:"src"`
	TileOrientation Orientation `json:"f"`
	T               int         `json:"t"`
	D               []int       `json:"d"`
	A               float64     `json:"a"`
}

type Orientation int

const (
	OrientationNone = iota
	OrientationFlipX
	OrientationFlipY
	OrientationFlipXY
)

func (l *Layer) ExtractLayerCSV(ignoredIDs []int) [][]int {
	bitmap := make([][]int, l.Height)
	for i := range bitmap {
		bitmap[i] = make([]int, l.Width)
	}

	if l.Type == LayerTypeIntGrid {
		// make 2D array from 1D array
		for i := range l.Height {
			for j := range l.Width {
				shouldIgnore := false
				for _, id := range ignoredIDs {
					if l.IntGridCsv[i*l.Height+j] == id {
						shouldIgnore = true
					}
				}
				if shouldIgnore {
					continue
				}
				bitmap[i][j] = l.IntGridCsv[i*l.Height+j]
			}
		}
	} else if l.Type == LayerTypeTiles {
		for _, tile := range l.AutoLayerTiles {
			posX := tile.Px[0] / l.GridSize
			posY := tile.Px[1] / l.GridSize
			bitmap[int(posY)][int(posX)] = 1
		}
		for _, tile := range l.GridTiles {
			posX := tile.Px[0] / l.GridSize
			posY := tile.Px[1] / l.GridSize
			bitmap[int(posY)][int(posX)] = 1
		}
	}

	return bitmap
}
