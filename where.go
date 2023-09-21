package healpix

import (
	"fmt"
	"math"
)

// An interface for converting between different indexing schemes and accessing the desired
// pixel index given a HEALPix map with a specific indexing scheme, regardless of which scheme
// the index itself references.
type Where interface {
	ToNestPixel(Healpix) NestPixel                       // Convert the index to an equivalent pixel index in Nest scheme.
	ToRingPixel(Healpix) RingPixel                       // Convert the index to an equivalent pixel index in Ring scheme.
	ToFacePixel(Healpix) FacePixel                       // Convert the index to an equivalent face 'x & y' index.
	ToRingCoordinate(Healpix) RingCoordinate             // Convert the index to an equivalent ring & offset index.
	ToProjectionCoordinate(Healpix) ProjectionCoordinate // Convert the index to an x/y position on the planar HEALPix projection of the sphere.
	ToSphereCoordinate(Healpix) SphereCoordinate         // Convert the index to a latitude/longitude position on the sphere.

	PixelId(Healpix, HealpixScheme) int // Convert the index into an equivalent index for the given HEALPix pixel numbering scheme.
}

// The index of a pixel in a HEALPix map using a 'quad-tree' division counting scheme that makes
// nearest neighbor searches more efficient.
type NestPixel int

// Identity function on NestPixel to satisfy the Where interface.
func (p NestPixel) ToNestPixel(hp Healpix) NestPixel {
	return p
}

func (p NestPixel) ToRingPixel(hp Healpix) RingPixel {
	return p.ToFacePixel(hp).ToRingCoordinate(hp).ToRingPixel(hp)
}

func (p NestPixel) ToFacePixel(hp Healpix) FacePixel {
	// the 12 faces store pixels in linear ranges, i.e. face 0 = 0 - n, face 1 = n - 2n, etc.
	// so we know which face we have by simply dividing by n = number of pixels per face
	face := int(p) / hp.FacePixels()
	// we can get the index of the pixel within the face by simple remainder
	facePixelId := int(p) % hp.FacePixels()

	// x and y are the compressed even and odd bits of facePixelId respectively
	x := 0
	y := 0
	// thanks Google
	intBits := 32 << (^uint(0) >> 63)
	// iteratively compress
	for i := 0; i < intBits/2; i++ {
		xMask := 1 << (i * 2)                 // even shifting
		yMask := 1 << ((i * 2) + 1)           // odd shifting
		x |= (xMask & facePixelId) >> i       // compress the correct amount each iteration
		y |= (yMask & facePixelId) >> (i + 1) // need to account for the odd offset in compression
	}
	return FacePixel{x, y, face}
}

func (p NestPixel) ToRingCoordinate(hp Healpix) RingCoordinate {
	return p.ToFacePixel(hp).ToRingCoordinate(hp)
}

func (p NestPixel) ToSphereCoordinate(hp Healpix) SphereCoordinate {
	return p.ToFacePixel(hp).ToRingCoordinate(hp).ToSphereCoordinate(hp)
}

func (p NestPixel) ToProjectionCoordinate(hp Healpix) ProjectionCoordinate {
	return p.ToFacePixel(hp).ToProjectionCoordinate(hp)
}

func (p NestPixel) PixelId(hp Healpix, scheme HealpixScheme) int {
	if scheme == NestScheme {
		return int(p)
	}
	return int(p.ToRingPixel(hp))
}

// The index of a pixel in a HEALPix map counting ring-wise down from the north pole.
type RingPixel int

func (p RingPixel) RingId(hp Healpix) int {
	return int(p)
}

func (p RingPixel) ToNestPixel(hp Healpix) NestPixel {
	return p.ToRingCoordinate(hp).ToFacePixel(hp).ToNestPixel(hp)
}

func (p RingPixel) ToRingPixel(hp Healpix) RingPixel {
	return p
}

func (p RingPixel) ToFacePixel(hp Healpix) FacePixel {
	return p.ToRingCoordinate(hp).ToFacePixel(hp)
}

func (p RingPixel) ToRingCoordinate(hp Healpix) RingCoordinate {
	// three cases: north polar cap, equatorial region, south polar cap
	if int(p) < hp.PolarRegionPixels() {
		pH := float64(int(p)+1) / float64(2)
		ringInd := int(math.Sqrt(pH - math.Sqrt(math.Floor(pH))))
		return RingCoordinate{
			ringInd,
			int(p) - 2*(ringInd+1)*ringInd,
		}
	} else if int(p) < hp.Pixels()-hp.PolarRegionPixels() {
		pE := int(p) - hp.PolarRegionPixels()
		ringInd := int(float64(pE)/(4*float64(hp.FaceSidePixels()))) + hp.FaceSidePixels() - 1
		return RingCoordinate{
			ringInd,
			pE % (4 * hp.FaceSidePixels()),
		}
	} else {
		// very similar to north polar cap, but some indices are inverted to account for counting
		// backwards from south pole as if it were zero
		nP := hp.Pixels() - int(p) - 1
		pH := float64(int(nP)+1) / float64(2)
		northRingInd := int(math.Sqrt(pH-math.Sqrt(math.Floor(pH)))) + 1
		ringInd := hp.Rings() - northRingInd
		return RingCoordinate{
			ringInd,
			2*(northRingInd+1)*northRingInd - 1 - nP,
		}
	}
}

func (p RingPixel) ToSphereCoordinate(hp Healpix) SphereCoordinate {
	return p.ToRingCoordinate(hp).ToSphereCoordinate(hp)
}

func (p RingPixel) ToProjectionCoordinate(hp Healpix) ProjectionCoordinate {
	return p.ToRingCoordinate(hp).ToFacePixel(hp).ToProjectionCoordinate(hp)
}

func (p RingPixel) PixelId(hp Healpix, scheme HealpixScheme) int {
	if scheme == RingScheme {
		return int(p)
	}
	return int(p.ToNestPixel(hp))
}

// Describes a HEALPix pixel as a combination of ring number and pixel-from-start-of-ring. The ring pixel numbering
// is analgous to this indexing scheme but with both components combined into a single number.
type RingCoordinate struct {
	ring        int // The HEALPix ring the pixel is located in.
	pixelInRing int // The index within the ring (assuming first pixel in the ring == 0) that contains the pixel.
}

// The HEALPix ring the pixel is located in.
func (p RingCoordinate) Ring() int {
	return p.ring
}

// The index within the ring (assuming first pixel in the ring == 0) that contains the pixel.
func (p RingCoordinate) PixelInRing() int {
	return p.pixelInRing
}

func (p RingCoordinate) ToNestPixel(hp Healpix) NestPixel {
	return p.ToFacePixel(hp).ToNestPixel(hp)
}

func (p RingCoordinate) ToRingPixel(hp Healpix) RingPixel {
	ring := NewRing(hp, p.ring)
	return RingPixel(ring.FirstIndex() + p.pixelInRing)
}

func (p RingCoordinate) ToFacePixel(hp Healpix) FacePixel {
	ring := NewRing(hp, p.ring)
	faceInd := 0
	nr := ring.northIndex + 1
	if ring.FirstIndex() < hp.PolarRegionPixels() {
		faceInd = p.pixelInRing / nr
	} else if ring.FirstIndex() < hp.Pixels()-hp.PolarRegionPixels() {
		nr = hp.FaceSidePixels()
		ire := (p.ring + 1) - hp.FaceSidePixels() + 1
		irm := hp.FaceSidePixels()*2 + 2 - ire
		ifm := ((p.pixelInRing + 1) - ire/2 + hp.FaceSidePixels() - 1) >> hp.Order()
		ifp := ((p.pixelInRing + 1) - irm/2 + hp.FaceSidePixels() - 1) >> hp.Order()
		if ifp == ifm {
			faceInd = ifp | 4
		} else if ifp < ifm {
			faceInd = ifp
		} else {
			faceInd = ifm + 8
		}
	} else {
		faceInd = 8 + p.pixelInRing/nr
	}

	southX, southY := NewFace(faceInd).SouthernmostVertex()
	shift := 1
	if ring.IsOffset() {
		shift = 0
	}
	irt := (p.ring + 1) - southY*hp.FaceSidePixels() + 1
	ipt := 2*(p.pixelInRing+1) - southX*nr - shift - 1
	if ipt >= hp.FaceSidePixels()*2 {
		ipt -= 8 * hp.FaceSidePixels()
	}
	x := (ipt - irt) >> 1
	y := (-ipt - irt) >> 1
	return FacePixel{x, y, faceInd}
}

func (p RingCoordinate) ToRingCoordinate(hp Healpix) RingCoordinate {
	return p
}

func (p RingCoordinate) ToProjectionCoordinate(hp Healpix) ProjectionCoordinate {
	return p.ToFacePixel(hp).ToProjectionCoordinate(hp)
}

func (p RingCoordinate) ToSphereCoordinate(hp Healpix) SphereCoordinate {
	// ring abstraction does the heavy liftng for latitude
	ring := NewRing(hp, p.ring)
	longitude := float64(0)
	// longitude is the same on both north and south hemispheres, so simplify by only computing for north
	if ring.northIndex < hp.FaceSidePixels() {
		longitude = (math.Pi / (2 * float64(ring.northIndex+1))) * (float64(p.pixelInRing) + float64(1)/float64(2))
	} else {
		shift := 0.0
		if ring.IsOffset() {
			shift = 1.0
		}
		longitude = (math.Pi / (2 * float64(hp.FaceSidePixels()))) * (float64(p.pixelInRing) + shift/2.0)
	}
	return SphereCoordinate{
		ring.Latitude(),
		ring.Colatitude(),
		longitude,
	}
}

func (p RingCoordinate) PixelId(hp Healpix, scheme HealpixScheme) int {
	if scheme == RingScheme {
		return int(p.ToRingPixel(hp))
	}
	return int(p.ToNestPixel(hp))
}

// Describes a discrete HEALPix pixel using the 'face' (base pixel) in which the pixel belongs, and it's relative
// x/y offset from the southernmost vertex of the face. So 0,0 is the x/y coordinate of the southernmost vertex on
// each face, and NSide-1,NSide-1 is the northernmost vertex of the face. X increases in a north-east direction,
// y increases in a north-west direction.
type FacePixel struct {
	x    int
	y    int
	face int
}

func (p FacePixel) X() int {
	return p.x
}

func (p FacePixel) Y() int {
	return p.y
}

func (p FacePixel) Face() int {
	return p.face
}

func (p FacePixel) ToNestPixel(hp Healpix) NestPixel {
	// first convert x and y into the pixel index within the face
	// we spread the bits of x and y then bitwise or them together
	// because x and y are compressed even and odd respectively of face pixel id
	intBits := 32 << (^uint(0) >> 63)
	x := 0
	y := 0
	for i := 0; i < intBits/2; i++ {
		mask := 1 << i               // when compressed both even and odd bits are in the same indices
		x |= (mask & p.x) << i       // spread the correct amount each iteration
		y |= (mask & p.y) << (i + 1) // need to account for the odd offset in spreading
	}
	facePixelId := x | y

	// then just add the number of preceding pixels to get the actual global index
	return NestPixel(facePixelId + p.face*hp.FacePixels())
}

func (p FacePixel) ToRingPixel(hp Healpix) RingPixel {
	return p.ToRingCoordinate(hp).ToRingPixel(hp)
}

func (p FacePixel) ToFacePixel(hp Healpix) FacePixel {
	return p
}

func (p FacePixel) ToRingCoordinate(hp Healpix) RingCoordinate {
	// find vertical and horizontal pixel boundary indices/vertices
	v := p.x + p.y
	h := p.x - p.y
	southX, southY := NewFace(p.face).SouthernmostVertex()
	// can now convert to ring
	ringId := southY*hp.FaceSidePixels() - v - 2
	ring := NewRing(hp, ringId)
	// find pixel in ring using the vertex offset
	s := 1
	if ring.IsOffset() {
		s = 0
	}

	pixelInRing := (southX*(ring.Pixels()>>2) + h + s) / 2
	if pixelInRing < 0 {
		pixelInRing += ring.Pixels() - 1
	}
	return RingCoordinate{ringId, pixelInRing}
}

func (p FacePixel) ToProjectionCoordinate(hp Healpix) ProjectionCoordinate {
	// find vertical and horizontal pixel boundary indices/vertices
	v := p.x + p.y
	h := p.x - p.y
	southX, southY := NewFace(p.face).SouthernmostVertex()
	// get ring and offset pixel center
	ringId := southY*hp.FaceSidePixels() - v - 2
	k := southX*hp.FaceSidePixels() + h

	x := float64(k) / float64(hp.FaceSidePixels()) * (math.Pi / 4)
	if x < 0 {
		x += 2 * math.Pi
	}

	return ProjectionCoordinate{
		x,
		(math.Pi / 2) - float64(ringId+1)/float64(hp.FaceSidePixels())*(math.Pi/4),
	}
}

func (p FacePixel) ToSphereCoordinate(hp Healpix) SphereCoordinate {
	return p.ToRingCoordinate(hp).ToSphereCoordinate(hp)
}

func (p FacePixel) PixelId(hp Healpix, scheme HealpixScheme) int {
	if scheme == RingScheme {
		return int(p.ToRingPixel(hp))
	}
	return int(p.ToNestPixel(hp))
}

// Represents a position on a HEALPix sphere projected into the standard HEALPix projection on a 2D plane.
type ProjectionCoordinate struct {
	x float64
	y float64
}

// The horizontal component of the coordinate on the planar projection. Correlated with longitude in spherical coordinates.
func (p ProjectionCoordinate) X() float64 {
	return p.x
}

// The vertical component of the coordinate on the planar projection. Correlated with lattitude in spherical coordinates.
func (p ProjectionCoordinate) Y() float64 {
	return p.y
}

func (p ProjectionCoordinate) ToNestPixel(hp Healpix) NestPixel {
	return p.ToFacePixel(hp).ToNestPixel(hp)
}

func (p ProjectionCoordinate) ToRingPixel(hp Healpix) RingPixel {
	return p.ToFacePixel(hp).ToRingCoordinate(hp).ToRingPixel(hp)
}

func (coord ProjectionCoordinate) ToFacePixel(hp Healpix) FacePixel {
	t := (4*coord.x)/math.Pi - 4
	u := (4*coord.y)/math.Pi + 5
	pp := (u + t) / 2
	if pp < 0 {
		pp = 0
	} else if pp > 5 {
		pp = 5
	}
	flooredPP := float64(int(pp))
	qq := (u - t) / 2
	if qq < 3-flooredPP {
		qq = 3 - flooredPP
	} else if qq > 6-flooredPP {
		qq = 6 - flooredPP
	}
	v := 5 - (int(pp) + int(qq))
	if v < 0 {
		return FacePixel{hp.FaceSidePixels(), hp.FaceSidePixels(), 0}
	}
	h := int(pp) - int(qq) + 4
	f := 4*v + (h>>1)%4
	p := math.Mod(float64(pp), 1)
	q := math.Mod(float64(qq), 1)
	x := int(float64(hp.FaceSidePixels()) * p)
	y := int(float64(hp.FaceSidePixels()) * q)
	return FacePixel{x, y, f}
}

func (p ProjectionCoordinate) ToRingCoordinate(hp Healpix) RingCoordinate {
	return p.ToFacePixel(hp).ToRingCoordinate(hp)
}

func (p ProjectionCoordinate) ToProjectionCoordinate(hp Healpix) ProjectionCoordinate {
	return p
}

func (p ProjectionCoordinate) ToSphereCoordinate(hp Healpix) SphereCoordinate {
	absY := math.Abs(p.y)
	if absY >= math.Pi/2 {
		panic(fmt.Sprintf("healpix: domain error in projection coordinate y dimension, %v too big", p.y))
	}

	if absY <= math.Pi/4 {
		// equatorial region
		z := (8 / (3 * math.Pi)) * p.y
		colat := math.Acos(z)
		return SphereCoordinate{math.Pi/2 - colat, colat, p.x}
	} else {
		// polar region
		tt := math.Mod(p.x, math.Pi/2)
		lng := p.x - ((absY-math.Pi/4)/(absY-math.Pi/2))*(tt-math.Pi/4)
		zz := 2 - 4*absY/math.Pi
		z := (1 - 1.0/3.0*(zz*zz)) * (p.y / absY)
		colat := math.Acos(z)
		return SphereCoordinate{math.Pi/2 - colat, colat, lng}
	}
}

func (p ProjectionCoordinate) PixelId(hp Healpix, scheme HealpixScheme) int {
	if scheme == RingScheme {
		return int(p.ToRingPixel(hp))
	}
	return int(p.ToNestPixel(hp))
}

// A position on the sphere represented by two components, the latitude (0 at equator and +/- Pi/2 at poles)
// and the longitude (0 - 2Pi). Units in radians. ALso provides colatitude/longitude representation for
// familiarity, as colatitude is used more frequently in HEALPix applications.
type SphereCoordinate struct {
	latitude   float64
	colatitude float64
	longitude  float64
}

// The latitude component of the coordinate on the sphere, in units of radians.
func (p SphereCoordinate) Latitude() float64 {
	return p.latitude
}

// The colatitude component of the coordinate on the sphere, in units of radians.
func (p SphereCoordinate) Colatitude() float64 {
	return p.colatitude
}

// The longitude component of the coordinate on the sphere, in units of radians.
func (p SphereCoordinate) Longitude() float64 {
	return p.longitude
}

func (p SphereCoordinate) ToNestPixel(hp Healpix) NestPixel {
	return p.ToProjectionCoordinate(hp).ToFacePixel(hp).ToNestPixel(hp)
}

func (p SphereCoordinate) ToRingPixel(hp Healpix) RingPixel {
	return p.ToProjectionCoordinate(hp).ToFacePixel(hp).ToRingCoordinate(hp).ToRingPixel(hp)
}

func (p SphereCoordinate) ToFacePixel(hp Healpix) FacePixel {
	return p.ToProjectionCoordinate(hp).ToFacePixel(hp)
}

func (p SphereCoordinate) ToRingCoordinate(hp Healpix) RingCoordinate {
	return p.ToProjectionCoordinate(hp).ToFacePixel(hp).ToRingCoordinate(hp)
}

func (p SphereCoordinate) ToProjectionCoordinate(hp Healpix) ProjectionCoordinate {
	z := math.Cos(p.colatitude)
	if math.Abs(z) <= 2.0/3.0 {
		// equatiorial region
		return ProjectionCoordinate{p.longitude, 3 * (math.Pi / 8) * z}
	} else {
		// polar region
		facetX := math.Mod(p.longitude, math.Pi/2)
		sigma := 2 - math.Sqrt(3*(1-math.Abs(z)))
		if z < 0 {
			sigma = -sigma
		}
		y := (math.Pi / 4) * sigma
		x := p.longitude - (math.Abs(sigma)-1)*(facetX-math.Pi/4)
		return ProjectionCoordinate{x, y}
	}
}

func (p SphereCoordinate) ToSphereCoordinate(hp Healpix) SphereCoordinate {
	return p
}

func (p SphereCoordinate) PixelId(hp Healpix, scheme HealpixScheme) int {
	if scheme == RingScheme {
		return int(p.ToRingPixel(hp))
	}
	return int(p.ToNestPixel(hp))
}
