# fyne-louple

A Fyne widget to show a photo in a scrolled container with
mouse panning enabled.

## Example
```
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/mschilli/fyne-loupe"
	"image"
	_ "image/jpeg"

	"fyne.io/fyne/v2/canvas"
	"log"
	"os"
)

const ViewSize = float32(1000)

func main() {
	a := app.New()
	w := a.NewWindow("Loupe Example")

	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fullImage := canvas.NewImageFromImage(img)

	l := loupe.NewLoupe(fullImage)
	w.SetContent(l.Scroll)
	w.Resize(fyne.NewSize(ViewSize, ViewSize))
	w.Show() // make sure Center() will reflect the actual window size later

	l.Center()

	w.ShowAndRun()
}
```

## Author

Mike Schilli, m@perlmeister.com 2025

## License

Released under the [Apache 2.0](LICENSE)
