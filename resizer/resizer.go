package resizer

import (
	"bytes"
	"log"

	"github.com/anthonynsimon/bild/imgio"

	"github.com/anthonynsimon/bild/transform"

	"image"
)

func Downscale(fileBytes []byte) {

	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	resized := transform.Resize(img, 400, 400, transform.Linear)

	if err := imgio.Save("resized.png", resized, imgio.PNGEncoder()); err != nil {
		log.Fatal(err)
		return
	}

}
