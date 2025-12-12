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
	//go:embed ebi_fry_3.png
	ebiFryPNG []byte
	//go:embed ebi_fry_3_rich.png
	ebiFryRichPNG []byte
	//go:embed big_mouse.png
	bigMousePNG []byte
	//go:embed sauce.png
	saucePNG []byte
	//go:embed virus.png
	virusPNG []byte
	//go:embed virus_computer.png
	virusComputerPNG []byte

	PlayerImage        *ebiten.Image
	EbiFryImage        *ebiten.Image
	EbiFryRichImage    *ebiten.Image
	SauceImage         *ebiten.Image
	VirusImage         *ebiten.Image
	VirusComputerImage *ebiten.Image
)

func init() {
	PlayerImage = lo.Must(LoadImage(bigMousePNG))
	EbiFryImage = lo.Must(LoadImage(ebiFryPNG))
	EbiFryRichImage = lo.Must(LoadImage(ebiFryRichPNG))
	SauceImage = lo.Must(LoadImage(saucePNG))
	VirusImage = lo.Must(LoadImage(virusPNG))
	VirusComputerImage = lo.Must(LoadImage(virusComputerPNG))
}

func LoadImage(data []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
