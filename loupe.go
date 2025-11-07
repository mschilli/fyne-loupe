package loupe

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
)

type Loupe struct {
	widget.BaseWidget
	image        *canvas.Image
	Scroll       *container.Scroll
	scrollOffset fyne.Position
	bounds       image.Rectangle
}

func NewLoupe(img *canvas.Image) *Loupe {
	l := &Loupe{image: img}
	l.ExtendBaseWidget(l)

	// save original image dimensions
	l.bounds = img.Image.Bounds()
	img.SetMinSize(fyne.NewSize(
		float32(l.bounds.Dx()),
		float32(l.bounds.Dy()),
	))

	scrollContent := container.NewMax(l)
	scroll := container.NewScroll(scrollContent)
	l.Scroll = scroll

	return l
}

type loupeRenderer struct {
	loupe *Loupe
}

func (lr *loupeRenderer) Layout(size fyne.Size) {
	lr.loupe.image.Resize(size)
}

func (lr *loupeRenderer) MinSize() fyne.Size {
	return lr.loupe.image.MinSize()
}

func (lr *loupeRenderer) Refresh() {
	// Don't refresh here, it'll kill performance
	// canvas.Refresh(lr.loupe.image)
}

func (lr *loupeRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (lr *loupeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{lr.loupe.image}
}

func (lr *loupeRenderer) Destroy() {}

func (l *Loupe) Dragged(e *fyne.DragEvent) {
	l.Scroll.Offset = l.Scroll.Offset.SubtractXY(e.Dragged.DX, e.Dragged.DY)
	l.Scroll.Refresh()
}

func (l *Loupe) CreateRenderer() fyne.WidgetRenderer {
	return &loupeRenderer{l}
}

func (l *Loupe) Center() {
	x := float32(l.bounds.Dx()/2) - l.Scroll.Size().Width/2.0
	y := float32(l.bounds.Dy()/2) - l.Scroll.Size().Height/2.0
	fmt.Printf("Centering to %f/%f\n", x, y)
	l.Scroll.Offset = fyne.NewPos(x, y)
	l.Scroll.Refresh()
}

func (l *Loupe) DragEnd() {
}
