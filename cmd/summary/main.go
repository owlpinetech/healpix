package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/owlpinetech/healpix"
)

func main() {
	order := flag.Int("order", -1, "Healpix order for the map")
	nside := flag.Int("nside", 0, "Healpix nside for the map")
	flag.Parse()

	if *order < 0 && *nside <= 0 {
		fmt.Println("HEALPix Library Limits:")
		fmt.Printf("\tMax Order: %d\n", healpix.MaxOrder())
		fmt.Printf("\tMax NSide: %d\n", healpix.MaxNSide())
		flag.Usage()
		return
	}

	var hp healpix.Healpix
	if *order >= 0 {
		if *nside > 0 {
			fmt.Println("Both order and nside specified. Using order value.")
		}
		if !healpix.IsValidOrder(*order) {
			fmt.Println("Invalid order value. Must be between 0 and", healpix.MaxOrder())
			return
		}
		hp = healpix.New(healpix.NewHealpixOrder(*order))
	} else if *nside > 0 {
		if !healpix.IsValidNSide(*nside) {
			fmt.Println("Invalid nside value. Must be a power of 2 and between 1 and", healpix.MaxNSide())
			return
		}
		hp = healpix.New(healpix.NewHealpixSide(*nside))
	}

	fmt.Println("HEALPix Summary:")
	fmt.Printf("\tOrder: %d\n", hp.Order())
	fmt.Printf("\tNSide: %d\n", hp.FaceSidePixels())
	fmt.Printf("\tFace Pixels: %d\n", hp.FacePixels())
	fmt.Printf("\tTotal Pixels: %d\n", hp.Pixels())
	fmt.Printf("\tPolar Region Pixels: %d\n", hp.PolarRegionPixels())
	fmt.Printf("\tRings: %d\n", hp.Rings())
	fmt.Printf("\tAngular Resolution: %.18f radians\n", hp.AngularResolution())

	degrees := hp.AngularResolution() * 180 / math.Pi
	fmt.Printf("\tAngular Resolution (Earth): %.18f m\n", (degrees*3600)*30.86)
	fmt.Printf("\tPixel Area: %.18f steradians\n", hp.PixelArea())
	fmt.Printf("\tPixel Surface Area (Earth): %.6f m^2\n", hp.PixelSurfaceArea(6371000))
	fmt.Printf("\tTotal Data Size (uint32): ~%d MB\n", hp.Pixels()/1e6*4)
}
