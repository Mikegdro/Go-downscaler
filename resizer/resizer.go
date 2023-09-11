package resizer

import (
	"bytes"

	"github.com/anthonynsimon/bild/imgio"

	"github.com/anthonynsimon/bild/transform"

	"image"
)

func Downscale(fileBytes []byte)(imageResized []byte, err error) {

	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, err
	}

	resized := transform.Resize(img, 400, 400, transform.Linear)

	if err := imgio.Save("resized.png", resized, imgio.PNGEncoder()); err != nil {
		return nil,  err
	}

	return imageResized, nil

}
