package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

var (
	//go:embed lunch.png
	lunchPNG []byte
	//go:embed ebi_fry.png
	ebiFryPNG []byte

	PlayerImage *ebiten.Image
	EbiFryImage *ebiten.Image
)

func init() {
	PlayerImage = lo.Must(LoadImage(lunchPNG))
	EbiFryImage = lo.Must(LoadImage(ebiFryPNG))
}

func LoadImage(data []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
