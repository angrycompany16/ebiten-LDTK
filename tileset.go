package ebitenLDTK

type Tileset struct {
	Name         string  `json:"identifier"`
	Uid          int     `json:"uid"`
	RelPath      string  `json:"relPath"`
	PxWid        int     `json:"pxWid"`
	PxHei        int     `json:"pxHei"`
	TileGridSize float64 `json:"tileGridSize"`
	Spacing      int     `json:"spacing"`
	Padding      int     `json:"padding"`
}
