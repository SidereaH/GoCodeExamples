package interfaces
import "fmt"
func Main() {
	/*
	   Interfaces
	   An interface type is defined as a set of method signatures.
	   A value of interface type can hold any value that implements those methods.
	   Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
	*/
	printInterface()
	/*
	Errors
	Go programs express error state with error values.
	The error type is a built-in interface similar to fmt.Stringer:
	type error interface {
	    Error() string
	}
	(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)
	Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
	i, err := strconv.Atoi("42")
	if err != nil {
	    fmt.Printf("couldn't convert number: %v\n", err)
	    return
	}
	fmt.Println("Converted integer:", i)
	A nil error denotes success; a non-nil error denotes failure
	 */

	if err := run(); err != nil {
		fmt.Println(err)
	}
	/*
	Exercise: Errors
	Copy your Sqrt function from the earlier exercise and modify it to return an error value.
	Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
	*/
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	/*
	Readers
	The io package specifies the io.Reader interface, which represents the read end of a stream of data.
	The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.
	The io.Reader interface has a Read method:
	func (T) Read(b []byte) (n int, err error)
	Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.
	The example code creates a strings.Reader and consumes its output 8 bytes at a time.
	*/
	readerPrint()
	/*
	Exercise: rot13Reader
	A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
	For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
	Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
	The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
	*/
	testRot13()
	/*
	Images
	Package image defines the Image interface:
	package image
	type Image interface {
	    ColorModel() color.Model
	    Bounds() Rectangle
	    At(x, y int) color.Color
	}
	Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.
	(See the documentation for all the details.)
	The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package
	*/
	createImage()
	/*
	Exercise: Images
	Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.
	Define your own Image type, implement the necessary methods, and call pic.ShowImage.
	Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
	ColorModel should return color.RGBAModel.
	At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
	*/
	testImage()
}
