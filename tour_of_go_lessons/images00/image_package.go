package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {

	var img *image.RGBA

	// initialize a new RGBA image with the given bounds (the associated Rectangle is the bounds of an RGBA image)

	img = image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(img.Bounds())

	fmt.Println(img.At(0, 0).RGBA())

}

// The Image Intreface as defined in the image package
type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// Contains minimum and maximum points as defined in the image package
type Rectangle struct {
	Min, Max Point
}

// Point is an X, Y coordinate pair, as defined in the image package
type Point struct {
	X, Y int
}

// image.Rect | image package exported function

//   - a shorthand to instantiate a Rectangle object

// Package image | Go Standard Library

//   - Defines the image interface
//   - Images contain colors, which are described in the image/color package

// Values of an Image Interface

//   - can be created in two ways:
//       - calling functions such as:
//  	     ~ image.NewRGBA or image.NewPalleted

//       - calling Decode on an io.Reader containing image data of the formats such as but not limited to
//         GIF, JPEG, or PNG
