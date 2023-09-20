package healpix

import (
	"math"
)

// Describes a ring of pixels in the HEALPix pixelization, and provides functionality on rings.
type Ring struct {
	base  Healpix
	index int

	// Skip similar code for southern hemisphere by temporarily converting all rings to northern hemisphere.
	// For ring indexes in the northern hemisphere of the map, index == northIndex.
	northIndex int
}

// Get a new ring description from the given HEALPix map, at the given ring index. Ring indices start at the
// the north pole at 0, and increment southward by 1 until reaching the south pole. The largest ring index possible
// for a given HEALPix map is base.Rings() - 1. Panics if the ring index is invalid.
func NewRing(base Healpix, index int) Ring {
	if index < 0 && index > base.Rings() {
		panic("healpix: ring index was invalid during ring creation")
	}

	north := index
	if index > base.FaceSidePixels()*2 {
		north = base.Rings() - index - 1
	}
	return Ring{
		base,
		index,
		north,
	}
}

// The healpix map of which this ring is a part. This determines some of the properties of the rings,
// and combined with Index() represents a 'minimal cover' of the attributes of a ring.
func (r Ring) Base() Healpix {
	return r.base
}

// The index of the ring, proceeding from 0 at the north pole and incrementing toward the south pole by one.
func (r Ring) Index() int {
	return r.index
}

// The equivalent Ring scheme pixel number of the first (0) pixel in this ring.
func (r Ring) FirstIndex() int {
	first := 0
	if r.northIndex < r.base.FaceSidePixels() {
		first = 2 * (r.northIndex + 1) * r.northIndex
	} else {
		first = r.base.PolarRegionPixels() + (r.northIndex+1-r.base.FaceSidePixels())*r.Pixels()
	}
	if r.northIndex != r.index {
		first = r.base.Pixels() - first - r.Pixels()
	}
	return first
}

// The number of pixels in this ring.
func (r Ring) Pixels() int {
	if r.northIndex < r.base.FaceSidePixels() {
		return 4 * (r.northIndex + 1)
	}
	return 4 * r.base.FaceSidePixels()
}

// The latitude where the center of each pixel in the ring lies, in radians.
func (r Ring) Colatitude() float64 {
	if r.northIndex < r.base.FaceSidePixels() {
		z := 1 - float64(r.northIndex+1)*float64(r.northIndex+1)/(float64(3)*float64(r.base.FacePixels()))
		if r.northIndex != r.index {
			return math.Pi - math.Acos(z)
		} else {
			return math.Acos(z)
		}
	} else {
		z := float64(4)/float64(3) - float64(2*(r.index+1))/float64(3*r.base.FaceSidePixels())
		return math.Acos(z)
	}
}

// The colatitutde where the center of each pixel in the ring lies, in radians.
func (r Ring) Latitude() float64 {
	return math.Pi/2 - r.Colatitude()
}

// True if the first pixel of the ring has it's center not on 0 longitude, but just slightly off it.
// False if the first pixel of the ring has it's center on 0 longitude.
func (r Ring) IsOffset() bool {
	if r.northIndex < r.base.FaceSidePixels() {
		return true
	}
	// only odd rings have their first pixel center on 0 longitude
	return ((r.northIndex - r.base.FaceSidePixels()) & 1) != 0
}
