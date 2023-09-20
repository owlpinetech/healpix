package healpix

import (
	"math"

	"github.com/owlpinetech/healpix/internal/intmath"
)

// A description of how pixels are accessed and stored on the machine. HEALPix provides two common
// indexing schemes for pixels: Ring and Nested. Ring starts pixel index 0 at the north pole and continues
// around the ring incrementing by one. When the first pixel is reach, the pixel numbering continues
// on the next ring below the north pole, again incrementing by one, until the south pole is reached.
// The Nested index is more complicated to describe textually, but allows for more efficient access
// of neighboring pixel indices mathemtically than the Ring scheme. Both are useful depending on the
// use case however.
type HealpixScheme int

const (
	// The HEALPix pixel numbering scheme in which pixel indexing starts 0 at the north pole and
	// increments laterally (around the ring of the same latitude) before moving down the globe
	// to the next 'ring'.
	RingScheme HealpixScheme = iota
	// The HEALPix pixel numbering scheme in which faces are split into a 'quadtree' like structure,
	// with the pixels numbered in a 'bottom-up' like fashion from the southernmost vertex of
	// of each face.
	NestScheme
)

const (
	BasePixelsPerRow int = 4
	BasePixelRows    int = 3
)

// The base interface describing the core aspects of a HEALPix projection.
type Healpix interface {
	Order() int                 // Returns the exponent describing how many pixels are in the HEALPix map.
	Pixels() int                // Returns the total number of pixels in the HEALPix map.
	FaceSidePixels() int        // Returns the number of pixels on the side of each base pixel (NSide) of the HEALPix map.
	FacePixels() int            // Returns the number of pixels in each base pixel (face) of the HEALPix map.
	PolarRegionPixels() int     // Returns the number of pixels in each (north or south) polar region.
	Rings() int                 // Returns the number of rings in the HEALPix map.
	EquatorRing() int           // Returns the index of the ring that sits on the equator in the HEALPix map.
	PixelArea() float64         // Returns the area of a pixel in the HEALPix map in steradians.
	AngularResolution() float64 // Returns the angular resolution of the HEALPix map in degrees.
}

// A description of a healpix map based on the order of the map. All other values
// are derived from this core descriptor (and computed each time they are accessed).
type HealpixOrder int

func NewHealpixOrder(order int) HealpixOrder {
	if !IsValidOrder(order) {
		panic("healpix: attempt to create HealpixOrder with invalid order argument")
	}
	return HealpixOrder(order)
}

// Returns the exponent describing how many pixels are in the HEALPix map.
func (o HealpixOrder) Order() int {
	return int(o)
}

func (o HealpixOrder) FaceSidePixels() int {
	return intmath.Exp2(int(o))
}

func (o HealpixOrder) FacePixels() int {
	nside := o.FaceSidePixels()
	return nside * nside
}

func (o HealpixOrder) Pixels() int {
	return 12 * o.FacePixels()
}

func (o HealpixOrder) PolarRegionPixels() int {
	return 2 * o.FaceSidePixels() * (o.FaceSidePixels() - 1)
}

func (o HealpixOrder) Rings() int {
	return 4*o.FaceSidePixels() - 1
}

func (o HealpixOrder) EquatorRing() int {
	return 2*o.FaceSidePixels() - 1
}

func (o HealpixOrder) PixelArea() float64 {
	return math.Pi / float64(3*o.FacePixels())
}

func (o HealpixOrder) AngularResolution() float64 {
	return math.Sqrt(o.PixelArea())
}

// A description of a healpix map based on the number of each pixels on each side of
// the base pixels of the map. Referred to as NSide in most HEALPix literature. All other
// values are derived from this core attributed (and computed on each access).
type HealpixSide int

func NewHealpixSide(nside int) HealpixSide {
	if !IsValidNSide(nside) {
		panic("healpix: attempt to create HealpixSide with invalid nside argument")
	}
	return HealpixSide(nside)
}

func (o HealpixSide) Order() int {
	return intmath.Log2(int(o))
}

func (o HealpixSide) FaceSidePixels() int {
	return int(o)
}

func (o HealpixSide) FacePixels() int {
	return int(o) * int(o)
}

func (o HealpixSide) Pixels() int {
	return 12 * o.FacePixels()
}

func (o HealpixSide) PolarRegionPixels() int {
	return 2 * o.FaceSidePixels() * (o.FaceSidePixels() - 1)
}

func (o HealpixSide) Rings() int {
	return 4*o.FaceSidePixels() - 1
}

func (o HealpixSide) EquatorRing() int {
	return 2*o.FaceSidePixels() - 1
}

func (o HealpixSide) PixelArea() float64 {
	return math.Pi / float64(3*o.FacePixels())
}

func (o HealpixSide) AngularResolution() float64 {
	return math.Sqrt(o.PixelArea())
}
