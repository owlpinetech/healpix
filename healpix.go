package healpix

import (
	"math"
	"math/bits"
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

// Technically these numbers below can be adjusted for different HEALPix configurations, but the standard
// 12 face configuration has nice properties and is easy to think about conceptually.

const (
	BasePixelsPerRow int = 4
	BasePixelRows    int = 3
)

// The base interface describing the core aspects of a HEALPix projection. All other HEALPix information can be
// derived from this core interface, though it may be useful in calculations to cache some of it rather than
// derive it on every access.
type HealpixBase interface {
	Order() int          // Returns the exponent describing how many pixels are in the HEALPix map.
	FaceSidePixels() int // Returns the number of pixels on the side of each base pixel (NSide) of the HEALPix map.
	FacePixels() int     // Returns the number of pixels in each base pixel (face) of the HEALPix map.
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

// Returns the number of pixels on the side of each base pixel (NSide) of the HEALPix map.
func (o HealpixOrder) FaceSidePixels() int {
	return 1 << o // 2 raised to the power of the order
}

// Returns the number of pixels in each base pixel (face) of the HEALPix map.
func (o HealpixOrder) FacePixels() int {
	nside := o.FaceSidePixels()
	return nside * nside
}

// A description of a healpix map based on the number of each pixels on each side of
// the base pixels of the map. Referred to as NSide in most HEALPix literature. All other
// values are derived from this core attributed (and computed on each access).
type HealpixSide int

func NewHealpixSide(nside int) HealpixSide {
	if !IsValidNSide(nside) {
		//panic("healpix: attempt to create HealpixSide with invalid nside argument")
	}
	return HealpixSide(nside)
}

// Returns the exponent describing how many pixels are in the HEALPix map.
func (o HealpixSide) Order() int {
	// floor(log2()) equivalent for integers
	return bits.Len(uint(o)) - 1
}

// Returns the number of pixels on the side of each base pixel (NSide) of the HEALPix map.
func (o HealpixSide) FaceSidePixels() int {
	return int(o)
}

// Returns the number of pixels in each base pixel (face) of the HEALPix map.
func (o HealpixSide) FacePixels() int {
	return int(o) * int(o)
}

// Lightweight abstraction over a HEALPix base configuration, to provide more derived information about the HEALPix map.
// Generally recommended to use this instead of HealpixOrder or HealpixSide directly, as it provides a more complete
// description of the HEALPix map.
type Healpix struct {
	HealpixBase
}

func New(base HealpixBase) Healpix {
	return Healpix{base}
}

// Returns the total number of pixels in the HEALPix map.
func (o Healpix) Pixels() uint {
	return 12 * uint(o.FacePixels())
}

// Returns the number of pixels in each (north or south) polar region.
func (o Healpix) PolarRegionPixels() int {
	return 2 * o.FaceSidePixels() * (o.FaceSidePixels() - 1)
}

// Returns the number of rings in the HEALPix map.
func (o Healpix) Rings() int {
	return 4*o.FaceSidePixels() - 1
}

// Returns the index of the ring that sits on the equator in the HEALPix map.
func (o Healpix) EquatorRing() int {
	return 2*o.FaceSidePixels() - 1
}

// Returns the area of a pixel in the HEALPix map in steradians.
func (o Healpix) PixelArea() float64 {
	return math.Pi / float64(3*o.FacePixels())
}

// Returns an approximation of the surface area of a pixel in the HEALPix map in meters squared, given
// a radius of the sphere in meters. This approximation has higher accuracy when the PixelArea() is less
// than 0.03 steradians, which is the case for HEALPix maps with orders >= 6.
func (o Healpix) PixelSurfaceArea(radius float64) float64 {
	return o.PixelArea() * radius * radius
}

// Returns the angular resolution of the HEALPix map in radians.
func (o Healpix) AngularResolution() float64 {
	return math.Sqrt(o.PixelArea())
}
