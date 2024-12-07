package interfaces
import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

func createImage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
/*
Exercise: Images
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.
Define your own Image type, implement the necessary methods, and call pic.ShowImage.
Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
ColorModel should return color.RGBAModel.
At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
 */
type Image struct{
	width, height int
}
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}
func (img *Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v,v,255,255}
}

func testImage() {
	m := &Image{10,10}
	pic.ShowImage(m) //works only at tour of go
}