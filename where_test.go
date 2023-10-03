package healpix

import (
	"math"
	"testing"
	"testing/quick"
)

func TestRingPixelRingCoord(t *testing.T) {
	testCases := []struct {
		name  string
		order HealpixOrder
		pixel RingPixel
		coord RingCoordinate
	}{
		{"0 order: 0 ring pixel = 0,0 ring coordinate", 0, 0, RingCoordinate{0, 0}},
		{"0 order: 1 ring pixel = 0,1 ring coordinate", 0, 1, RingCoordinate{0, 1}},
		{"0 order: 2 ring pixel = 0,2 ring coordinate", 0, 2, RingCoordinate{0, 2}},
		{"0 order: 4 ring pixel = 1,0 ring coordinate", 0, 4, RingCoordinate{1, 0}},
		{"0 order: 5 ring pixel = 1,1 ring coordinate", 0, 5, RingCoordinate{1, 1}},
		{"0 order: 8 ring pixel = 2,0 ring coordinate", 0, 8, RingCoordinate{2, 0}},
		{"0 order: 9 ring pixel = 2,1 ring coordinate", 0, 9, RingCoordinate{2, 1}},
		{"0 order: 10 ring pixel = 2,2 ring coordinate", 0, 10, RingCoordinate{2, 2}},
		{"0 order: 11 ring pixel = 2,3 ring coordinate", 0, 11, RingCoordinate{2, 3}},

		{"1 order: 0 ring pixel = 0,0 ring coordinate", 1, 0, RingCoordinate{0, 0}},
		{"1 order: 1 ring pixel = 0,1 ring coordinate", 1, 1, RingCoordinate{0, 1}},
		{"1 order: 4 ring pixel = 1,0 ring coordinate", 1, 4, RingCoordinate{1, 0}},
		{"1 order: 9 ring pixel = 1,5 ring coordinate", 1, 9, RingCoordinate{1, 5}},
		{"1 order: 12 ring pixel = 2,0 ring coordinate", 1, 12, RingCoordinate{2, 0}},
		{"1 order: 20 ring pixel = 3,0 ring coordinate", 1, 20, RingCoordinate{3, 0}},
		{"1 order: 28 ring pixel = 4,0 ring coordinate", 1, 28, RingCoordinate{4, 0}},
		{"1 order: 36 ring pixel = 5,0 ring coordinate", 1, 36, RingCoordinate{5, 0}},
		{"1 order: 44 ring pixel = 6,0 ring coordinate", 1, 44, RingCoordinate{6, 0}},
		{"1 order: 47 ring pixel = 6,3 ring coordinate", 1, 47, RingCoordinate{6, 3}},

		{"2 order: 0 ring pixel = 0,0 ring coordinate", 2, 0, RingCoordinate{0, 0}},
		{"2 order: 1 ring pixel = 0,1 ring coordinate", 2, 1, RingCoordinate{0, 1}},
		{"2 order: 4 ring pixel = 1,0 ring coordinate", 2, 4, RingCoordinate{1, 0}},
		{"2 order: 9 ring pixel = 1,5 ring coordinate", 2, 9, RingCoordinate{1, 5}},
		{"2 order: 12 ring pixel = 2,0 ring coordinate", 2, 12, RingCoordinate{2, 0}},
		{"2 order: 24 ring pixel = 3,0 ring coordinate", 2, 24, RingCoordinate{3, 0}},
		{"2 order: 40 ring pixel = 4,0 ring coordinate", 2, 40, RingCoordinate{4, 0}},
		{"2 order: 136 ring pixel = 10,0 ring coordinate", 2, 136, RingCoordinate{10, 0}},
		{"2 order: 152 ring pixel = 11,0 ring coordinate", 2, 152, RingCoordinate{11, 0}},
		{"2 order: 168 ring pixel = 12,0 ring coordinate", 2, 168, RingCoordinate{12, 0}},
	}

	for ind, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rCoord := testCases[ind].pixel.ToRingCoordinate(tc.order)
			rRing := testCases[ind].coord.ToRingPixel(tc.order)
			if rCoord != tc.coord {
				t.Errorf("Ring to coordinate expected %v, got %v instead", tc.coord, rCoord)
			}
			if rRing != tc.pixel {
				t.Errorf("Coordinate to ring expected %v, got %v instead", tc.pixel, rRing)
			}
		})
	}
}

func TestFacePixelRingCoord(t *testing.T) {
	testCases := []struct {
		name  string
		order HealpixOrder
		pixel FacePixel
		coord RingCoordinate
	}{
		{"0 order: 0,0,0 face pixel = 0,0 ring coordinate", 0, FacePixel{0, 0, 0}, RingCoordinate{0, 0}},
		{"0 order: 0,0,1 face pixel = 0,1 ring coordinate", 0, FacePixel{0, 0, 1}, RingCoordinate{0, 1}},
		{"0 order: 0,0,2 face pixel = 0,2 ring coordinate", 0, FacePixel{0, 0, 2}, RingCoordinate{0, 2}},
		{"0 order: 0,0,4 face pixel = 1,0 ring coordinate", 0, FacePixel{0, 0, 4}, RingCoordinate{1, 0}},
		{"0 order: 0,0,5 face pixel = 1,1 ring coordinate", 0, FacePixel{0, 0, 5}, RingCoordinate{1, 1}},
		{"0 order: 0,0,8 face pixel = 2,0 ring coordinate", 0, FacePixel{0, 0, 8}, RingCoordinate{2, 0}},
		{"0 order: 0,0,9 face pixel = 2,1 ring coordinate", 0, FacePixel{0, 0, 9}, RingCoordinate{2, 1}},
		{"0 order: 0,0,10 face pixel = 2,2 ring coordinate", 0, FacePixel{0, 0, 10}, RingCoordinate{2, 2}},
		{"0 order: 0,0,11 face pixel = 2,3 ring coordinate", 0, FacePixel{0, 0, 11}, RingCoordinate{2, 3}},

		{"1 order: 1,1,0 face pixel = 0,0 ring coordinate", 1, FacePixel{1, 1, 0}, RingCoordinate{0, 0}},
		{"1 order: 1,1,1 face pixel = 0,1 ring coordinate", 1, FacePixel{1, 1, 1}, RingCoordinate{0, 1}},
		{"1 order: 1,1,4 face pixel = 2,0 ring coordinate", 1, FacePixel{1, 1, 4}, RingCoordinate{2, 0}},
		{"1 order: 1,1,5 face pixel = 2,2 ring coordinate", 1, FacePixel{1, 1, 5}, RingCoordinate{2, 2}},
		{"1 order: 0,0,5 face pixel = 5,2 ring coordinate", 1, FacePixel{0, 0, 5}, RingCoordinate{4, 2}},
		{"1 order: 1,0,4 face pixel = 3,0 ring coordinate", 1, FacePixel{1, 0, 4}, RingCoordinate{3, 0}},
		{"1 order: 1,1,8 face pixel = 4,1 ring coordinate", 1, FacePixel{1, 1, 8}, RingCoordinate{4, 1}},
		{"1 order: 0,0,8 face pixel = 6,0 ring coordinate", 1, FacePixel{0, 0, 8}, RingCoordinate{6, 0}},
		{"1 order: 0,0,11 face pixel = 6,3 ring coordinate", 1, FacePixel{0, 0, 11}, RingCoordinate{6, 3}},
	}

	for ind, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rCoord := testCases[ind].pixel.ToRingCoordinate(tc.order)
			rFace := testCases[ind].coord.ToFacePixel(tc.order)
			if rCoord != tc.coord {
				t.Errorf("Face to coordinate expected %v, got %v instead", tc.coord, rCoord)
			}
			if rFace != tc.pixel {
				t.Errorf("Coordinate to face expected %v, got %v instead", tc.pixel, rFace)
			}
		})
	}
}

func TestNestPixelFacePixel(t *testing.T) {
	testCases := []struct {
		name  string
		order HealpixOrder
		nest  NestPixel
		face  FacePixel
	}{
		{"0 order: 0 nest pixel = 0,0,0 face pixel", 0, 0, FacePixel{0, 0, 0}},
		{"0 order: 1 nest pixel = 0,0,1 face pixel", 0, 1, FacePixel{0, 0, 1}},
		{"0 order: 2 nest pixel = 0,0,2 face pixel", 0, 2, FacePixel{0, 0, 2}},
		{"0 order: 4 nest pixel = 0,0,4 face pixel", 0, 4, FacePixel{0, 0, 4}},
		{"0 order: 5 nest pixel = 0,0,5 face pixel", 0, 5, FacePixel{0, 0, 5}},
		{"0 order: 8 nest pixel = 0,0,8 face pixel", 0, 8, FacePixel{0, 0, 8}},
		{"0 order: 9 nest pixel = 0,0,9 face pixel", 0, 9, FacePixel{0, 0, 9}},
		{"0 order: 10 nest pixel = 0,0,10 face pixel", 0, 10, FacePixel{0, 0, 10}},
		{"0 order: 11 nest pixel = 0,0,11 face pixel", 0, 11, FacePixel{0, 0, 11}},

		{"1 order: 0 nest pixel = 0,0,0 face pixel", 1, 0, FacePixel{0, 0, 0}},
		{"1 order: 1 nest pixel = 1,0,0 face pixel", 1, 1, FacePixel{1, 0, 0}},
		{"1 order: 2 nest pixel = 0,1,0 face pixel", 1, 2, FacePixel{0, 1, 0}},
		{"1 order: 4 nest pixel = 0,0,1 face pixel", 1, 4, FacePixel{0, 0, 1}},
		{"1 order: 5 nest pixel = 1,0,1 face pixel", 1, 5, FacePixel{1, 0, 1}},
		{"1 order: 8 nest pixel = 0,0,2 face pixel", 1, 8, FacePixel{0, 0, 2}},
		{"1 order: 9 nest pixel = 1,0,2 face pixel", 1, 9, FacePixel{1, 0, 2}},
		{"1 order: 10 nest pixel = 0,1,2 face pixel", 1, 10, FacePixel{0, 1, 2}},
		{"1 order: 11 nest pixel = 1,1,2 face pixel", 1, 11, FacePixel{1, 1, 2}},

		{"2 order: 0 nest pixel = 0,0,0 face pixel", 2, 0, FacePixel{0, 0, 0}},
		{"2 order: 1 nest pixel = 1,0,0 face pixel", 2, 1, FacePixel{1, 0, 0}},
		{"2 order: 4 nest pixel = 2,0,0 face pixel", 2, 4, FacePixel{2, 0, 0}},
		{"2 order: 9 nest pixel = 1,2,0 face pixel", 2, 9, FacePixel{1, 2, 0}},

		{"Max order: 1314064518130923784 nest pixel = ?,?,? face pixel", HealpixOrder(MaxOrder()), 1314064518130923784, FacePixel{115814864, 377337186, 4}},
	}

	for ind, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rFace := testCases[ind].nest.ToFacePixel(tc.order)
			rNest := testCases[ind].face.ToNestPixel(tc.order)
			if rFace != tc.face {
				t.Errorf("Nest to face expected %v, got %v instead", tc.face, rFace)
			}
			if rNest != tc.nest {
				t.Errorf("Face to nest expected %v, got %v instead", tc.nest, rNest)
			}
		})
	}
}

func TestNestPixelUniquePixel(t *testing.T) {
	testCases := []struct {
		name  string
		order HealpixOrder
		nest  NestPixel
		uniq  UniquePixel
	}{
		{"0 order: 0 nest pixel = 4 unique pixel", 0, 0, 4},
		{"0 order: 1 nest pixel = 5 unique pixel", 0, 1, 5},
		{"0 order: 2 nest pixel = 6 unique pixel", 0, 2, 6},
		{"0 order: 11 nest pixel = 15 unique pixel", 0, 11, 15},

		{"1 order: 0 nest pixel = 16 unique pixel", 1, 0, 16},
		{"1 order: 1 nest pixel = 17 unique pixel", 1, 1, 17},
		{"1 order: 2 nest pixel = 18 unique pixel", 1, 2, 18},

		{"2 order: 0 nest pixel = 64 unique pixel", 2, 0, 64},
		{"2 order: 1 nest pixel = 65 unique pixel", 2, 1, 65},
		{"2 order: 2 nest pixel = 66 unique pixel", 2, 2, 66},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rUniq := tc.nest.ToUniquePixel(tc.order)
			rNest := tc.uniq.ToNestPixel(tc.order)
			if rUniq != tc.uniq {
				t.Errorf("Nest to unique expected %v, got %v instead", tc.uniq, rUniq)
			}
			if rNest != tc.nest {
				t.Errorf("Unique to nest expected %v, got %v instead", tc.nest, rNest)
			}
		})
	}
}

func TestPositionToNestPixel(t *testing.T) {
	testCases := []struct {
		name       string
		order      HealpixOrder
		nest       NestPixel
		colatitude float64
		longitude  float64
	}{
		{"0 order: 0 nest pixel = Pi/4, Pi/4", 0, 0, 0.841068670567930, math.Pi / 4},
		{"0 order: 1 nest pixel = Pi/4, 3Pi/4", 0, 1, 0.841068670567930, 3 * math.Pi / 4},
		{"0 order: 2 nest pixel = Pi/4, 5Pi/4", 0, 2, 0.841068670567930, 5 * math.Pi / 4},
		{"0 order: 4 nest pixel = Pi/2, 0", 0, 4, math.Pi / 2, 0},
		{"0 order: 5 nest pixel = Pi/2, Pi/2", 0, 5, math.Pi / 2, math.Pi / 2},
		{"0 order: 8 nest pixel = 3Pi/2, Pi/4", 0, 8, 2.300523983021862982, math.Pi / 4},
		{"0 order: 9 nest pixel = 3Pi/2, 3Pi/4", 0, 9, 2.300523983021862982, 3 * math.Pi / 4},
		{"0 order: 11 nest pixel = 3Pi/2, 7Pi/4", 0, 11, 2.300523983021862982, 7 * math.Pi / 4},

		{"1 order: 0 nest pixel = Pi/3, Pi/4", 1, 0, 1.2309594173407746, math.Pi / 4},
		{"1 order: 1 nest pixel = Pi/4, 3Pi/8", 1, 1, 0.84106867056793, 3 * math.Pi / 8},
		{"1 order: 2 nest pixel = Pi/4, Pi/8", 1, 2, 0.84106867056793, math.Pi / 8},
		{"1 order: 4 nest pixel = 0,0,1 face pixel", 1, 4, 1.2309594173407746, 3 * math.Pi / 4},
		{"1 order: 16 nest pixel = arcos(-1/3), 0", 1, 16, math.Acos(-1.0 / 3.0), 0},
		{"1 order: 17 nest pixel = Pi / 2, Pi/8", 1, 17, math.Pi / 2, math.Pi / 8},

		{"Max order: 1314064518130923784 nest pixel = ?,?,? face pixel", HealpixOrder(MaxOrder()), 1314064518130923784, 1.6251115119976574, 5.900599574193858},
	}

	for ind, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rPos := testCases[ind].nest.ToSphereCoordinate(tc.order)
			pos := SphereCoordinate{math.Pi/2 - tc.colatitude, tc.colatitude, tc.longitude}
			rNest := pos.ToNestPixel(tc.order)
			if rNest != tc.nest {
				t.Errorf("Position to nest expected %v, got %v instead", tc.nest, rNest)
			}
			if !withinTolerance(rPos.colatitude, tc.colatitude, 0.000000001) || !withinTolerance(rPos.longitude, tc.longitude, 0.000000001) {
				t.Errorf("Nest to position expected %v,%v but got %v,%v instead", tc.colatitude, tc.longitude, rPos.colatitude, rPos.longitude)
			}
		})
	}
}

func TestPositionToProjection(t *testing.T) {
	testCases := []struct {
		name       string
		colatitude float64
		longitude  float64
		xproj      float64
		yproj      float64
	}{
		{"lat/lng Pi/2,0 = x/y 0,0", math.Pi / 2, 0, 0, 0},
		{"lat/lng Pi/2,2Pi = x/y 2Pi,0", math.Pi / 2, 2 * math.Pi, 2 * math.Pi, 0},
		{"lat/lng arccos(1/3),Pi = x/y Pi,Pi/8", math.Acos(1.0 / 3.0), math.Pi, math.Pi, math.Pi / 8},
		{"lat/lng Pi/4,Pi = x/y Pi,Pi/4", math.Acos(2.0 / 3.0), math.Pi, math.Pi, math.Pi / 4},
	}

	for _, tc := range testCases {
		hp := HealpixOrder(0)
		t.Run(tc.name, func(t *testing.T) {
			pos := SphereCoordinate{math.Pi/2 - tc.colatitude, tc.colatitude, tc.longitude}
			proj := ProjectionCoordinate{tc.xproj, tc.yproj}
			rProj := pos.ToProjectionCoordinate(hp)
			rPos := proj.ToSphereCoordinate(hp)
			if !withinTolerance(rPos.colatitude, tc.colatitude, 0.000000001) || !withinTolerance(rPos.longitude, tc.longitude, 0.000000001) {
				t.Errorf("Projection to position expected %v,%v but got %v,%v instead", tc.colatitude, tc.longitude, rPos.colatitude, rPos.longitude)
			}
			if !withinTolerance(rProj.x, tc.xproj, 0.000000001) || !withinTolerance(rProj.y, tc.yproj, 0.000000001) {
				t.Errorf("Position to projection expected %v,%v but got %v,%v instead", tc.xproj, tc.yproj, rProj.x, rProj.y)
			}
		})
	}
}

func TestProjectionToFacePixel(t *testing.T) {
	testCases := []struct {
		name  string
		order HealpixOrder
		xproj float64
		yproj float64
		face  FacePixel
	}{
		{"0 order: Pi/4,Pi/4 = {0,0,0}", 0, math.Pi / 4, math.Pi / 4, FacePixel{0, 0, 0}},
		{"0 order: 3Pi/4,Pi/4 = {0,0,1}", 0, 3 * math.Pi / 4, math.Pi / 4, FacePixel{0, 0, 1}},
		{"0 order: 5Pi/4,Pi/4 = {0,0,2}", 0, 5 * math.Pi / 4, math.Pi / 4, FacePixel{0, 0, 2}},
		{"0 order: 7Pi/4,Pi/4 = {0,0,3}", 0, 7 * math.Pi / 4, math.Pi / 4, FacePixel{0, 0, 3}},
		{"0 order: 0,0 = {0,0,4}", 0, 0, 0, FacePixel{0, 0, 4}},
		{"0 order: Pi/2,0 = {0,0,5}", 0, math.Pi / 2, 0, FacePixel{0, 0, 5}},
		{"0 order: Pi,0 = {0,0,6}", 0, math.Pi, 0, FacePixel{0, 0, 6}},
		{"0 order: 3Pi/2,0 = {0,0,7}", 0, 3 * math.Pi / 2, 0, FacePixel{0, 0, 7}},
		{"0 order: Pi/4,-Pi/4 = {0,0,8}", 0, math.Pi / 4, -math.Pi / 4, FacePixel{0, 0, 8}},
		{"0 order: 3Pi/4,-Pi/4 = {0,0,9}", 0, 3 * math.Pi / 4, -math.Pi / 4, FacePixel{0, 0, 9}},
		{"0 order: 5Pi/4,-Pi/4 = {0,0,10}", 0, 5 * math.Pi / 4, -math.Pi / 4, FacePixel{0, 0, 10}},
		{"0 order: 7Pi/4,-Pi/4 = {0,0,11}", 0, 7 * math.Pi / 4, -math.Pi / 4, FacePixel{0, 0, 11}},

		{"1 order: Pi/4,3Pi/8 = {1,1,0}", 1, math.Pi / 4, 3 * math.Pi / 8, FacePixel{1, 1, 0}},
		{"1 order: Pi/4,Pi/8 = {0,0,0}", 1, math.Pi / 4, math.Pi / 8, FacePixel{0, 0, 0}},
		{"1 order: Pi/4,-Pi/8 = {1,1,8}", 1, math.Pi / 4, math.Pi / -8, FacePixel{1, 1, 8}},
		{"1 order: Pi/4,-3Pi/8 = {0,0,8}", 1, math.Pi / 4, -3 * math.Pi / 8, FacePixel{0, 0, 8}},
		{"1 order: 0,Pi/8 = {1,1,4}", 1, 0, math.Pi / 8, FacePixel{1, 1, 4}},
		{"1 order: 0,-Pi/8 = {0,0,4}", 1, 0, -math.Pi / 8, FacePixel{0, 0, 4}},
		{"1 order: 15Pi/8,0 = {0,1,4}", 1, 15 * math.Pi / 8, 0, FacePixel{0, 1, 4}},
		{"1 order: Pi/8,0 = {1,0,4}", 1, math.Pi / 8, 0, FacePixel{1, 0, 4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			proj := ProjectionCoordinate{tc.xproj, tc.yproj}
			rFace := proj.ToFacePixel(tc.order)
			rProj := tc.face.ToProjectionCoordinate(tc.order)
			if !withinTolerance(rProj.x, tc.xproj, 0.000000001) || !withinTolerance(rProj.y, tc.yproj, 0.000000001) {
				t.Errorf("Face pixel to projection expected %v,%v but got %v,%v instead", tc.xproj, tc.yproj, rProj.x, rProj.y)
			}
			if rFace != tc.face {
				t.Errorf("Projection to face pixel expected %v, but got %v instead", tc.face, rFace)
			}
		})
	}
}

func TestSphereToProjectionInvertible(t *testing.T) {
	testCases := []struct {
		name       string
		colatitude float64
		longitude  float64
	}{
		{"colat/lng Pi/2,0", math.Pi / 2, 0},
		{"colat/lng Pi/4,Pi", math.Pi / 4, math.Pi},
		{"colat/lng Pi/3,3Pi/8", math.Pi / 3, 3 * math.Pi / 8},
		{"colat/lng Pi/6,7Pi/4", math.Pi / 6, 7 * math.Pi / 4},
		{"colat/lng 1,1", 1, 1},
		{"colat/lng 3.14,6.28", 3.14, 6.28},
	}

	for _, tc := range testCases {
		// the healpix argument is not actually used in sphere-projection conversions
		hp := HealpixOrder(0)
		t.Run(tc.name, func(t *testing.T) {
			pos := SphereCoordinate{math.Pi/2 - tc.colatitude, tc.colatitude, tc.longitude}
			rPos := pos.ToProjectionCoordinate(hp).ToSphereCoordinate(hp)
			if !withinTolerance(rPos.colatitude, tc.colatitude, 0.000000001) || !withinTolerance(rPos.longitude, tc.longitude, 0.000000001) {
				t.Errorf("Position inverted expected %v,%v but got %v,%v instead", tc.colatitude, tc.longitude, rPos.colatitude, rPos.longitude)
			}
		})
	}
}

func TestRingCoordToFacePixelInvertible(t *testing.T) {
	testCases := []struct {
		name        string
		order       int
		ring        int
		pixelInRing int
	}{
		{"order 0, ring 0, pixel 0", 0, 0, 0},
		{"order max, ring 0, pixel 0", MaxOrder(), 0, 0},
		{"order max, ring max, pixel 0", MaxOrder(), NewHealpixOrder(MaxOrder()).Rings() - 1, 0},
		{"order max, ring 1251133056, pixel 2095318657", MaxOrder(), 1251133056, 2095318657},
	}

	for _, tc := range testCases {
		// the healpix argument is not actually used in sphere-projection conversions
		hp := HealpixOrder(tc.order)
		t.Run(tc.name, func(t *testing.T) {
			coord := RingCoordinate{tc.ring, tc.pixelInRing}
			face := coord.ToFacePixel(hp)
			rPos := face.ToRingCoordinate(hp)
			if coord.ring != rPos.ring || coord.pixelInRing != rPos.pixelInRing {
				t.Errorf("Position inverted expected %v,%v but got %v,%v instead (through %v,%v,%v)", coord.ring, coord.pixelInRing, rPos.ring, rPos.pixelInRing, face.face, face.x, face.y)
			}
		})
	}
}

func TestSphereConstructorsSame(t *testing.T) {
	latToColatSame := func(latitude float64, longitude float64) bool {
		latLon := NewLatLonCoordinate(latitude, longitude)
		colatLon := NewColatLonCoordinate(latLon.Colatitude(), latLon.Longitude())
		return withinTolerance(latLon.Latitude(), colatLon.Latitude(), 0.000000001)
	}

	colatToLatSame := func(colatitude float64, longitude float64) bool {
		colatLon := NewColatLonCoordinate(colatitude, longitude)
		latLon := NewLatLonCoordinate(colatLon.Latitude(), colatLon.Longitude())
		return withinTolerance(latLon.Latitude(), colatLon.Latitude(), 0.000000001)
	}

	if err := quick.Check(latToColatSame, nil); err != nil {
		t.Errorf("Latitude was different after converted to colatitude and back: %v", err)
	}
	if err := quick.Check(colatToLatSame, nil); err != nil {
		t.Errorf("Colatitude was different after converted to latitude and back: %v", err)
	}
}

func TestConversionInverses(t *testing.T) {
	hp := NewHealpixOrder(MaxOrder())

	// ring pixels to ring coordinate is perfectly invertible
	ringPixelToRingCoordInvertible := func(ringPixel RingPixel) bool {
		if ringPixel >= RingPixel(hp.Pixels()) || ringPixel < 0 {
			return true
		}
		return ringPixel == ringPixel.ToRingCoordinate(hp).ToRingPixel(hp)
	}

	// ring coordinate to face pixel is perfectly invertible
	ringCoordToFacePixelInvertible := func(ringPixel RingPixel) bool {
		if ringPixel >= RingPixel(hp.Pixels()) || ringPixel < 0 {
			return true
		}
		ringCoord := ringPixel.ToRingCoordinate(hp)
		return ringCoord == ringCoord.ToFacePixel(hp).ToRingCoordinate(hp)
	}

	// ring pixels to face pixels is perfectly invertible
	ringPixelToFacePixelInvertible := func(ringPixel RingPixel) bool {
		if ringPixel >= RingPixel(hp.Pixels()) || ringPixel < 0 {
			return true
		}
		return ringPixel == ringPixel.ToFacePixel(hp).ToRingPixel(hp)
	}

	// nested pixels to face pixels is perfectly invertible
	nestPixelToFacePixelInvertible := func(nestPixel NestPixel) bool {
		if nestPixel >= NestPixel(hp.Pixels()) || nestPixel < 0 {
			return true
		}
		return nestPixel == nestPixel.ToFacePixel(hp).ToNestPixel(hp)
	}

	// ring pixels to nest pixels is perfectly invertible
	nestPixelToRingPixelInvertible := func(nest NestPixel) bool {
		if nest >= NestPixel(hp.Pixels()) || nest < 0 {
			return true
		}
		return nest == nest.ToRingPixel(hp).ToNestPixel(hp)
	}

	// nest pixels to position is perfectly invertible
	nestPixelToSphereCoordinateInvertible := func(nest NestPixel) bool {
		if nest >= NestPixel(hp.Pixels()) || nest < 0 {
			return true
		}
		return nest == nest.ToSphereCoordinate(hp).ToNestPixel(hp)
	}

	if err := quick.Check(ringPixelToRingCoordInvertible, nil); err != nil {
		t.Errorf("Ring pixel was different after converted to ring coordinate and back: %v", err)
	}
	if err := quick.Check(ringCoordToFacePixelInvertible, nil); err != nil {
		t.Errorf("Ring coordinate was different after converted to face pixel and back: %v", err)
	}
	if err := quick.Check(ringPixelToFacePixelInvertible, nil); err != nil {
		t.Errorf("Ring pixel was different after converted to face pixel and back: %v", err)
	}
	if err := quick.Check(nestPixelToFacePixelInvertible, nil); err != nil {
		t.Errorf("Nest pixel was different after converted to face pixel and back: %v", err)
	}
	if err := quick.Check(nestPixelToRingPixelInvertible, nil); err != nil {
		t.Errorf("Nest pixel was different after converted to ring pixel and back: %v", err)
	}
	if err := quick.Check(nestPixelToSphereCoordinateInvertible, nil); err != nil {
		t.Errorf("Nest pixel was different after converted to position and back: %v", err)
	}
}

func TestNestRingSpherePositionsSame(t *testing.T) {
	hp := NewHealpixOrder(MaxOrder())

	nestRingSpherePositionsSame := func(nest NestPixel) bool {
		if nest >= NestPixel(hp.Pixels()) || nest < 0 {
			return true
		}
		ring := nest.ToRingPixel(hp)
		nPos := nest.ToSphereCoordinate(hp)
		rPos := ring.ToSphereCoordinate(hp)
		return withinTolerance(nPos.colatitude, rPos.colatitude, 0.000000001) &&
			withinTolerance(nPos.longitude, rPos.longitude, 0.000000001)
	}

	if err := quick.Check(nestRingSpherePositionsSame, nil); err != nil {
		t.Errorf("Nest pixel and equivalent ring pixel do not have the same sphere position: %v", err)
	}
}
