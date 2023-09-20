package healpix

// Given a desired coordinate on a healpix map, return the pixel index of
// of the desired neighbor pixel of in the selected HEALPix numbering scheme.
func Neighbor(hp Healpix, scheme HealpixScheme, where Where, xo int, yo int) int {
	fp := where.ToFacePixel(hp)
	maxXY := hp.FaceSidePixels() - 1

	x := fp.x + xo
	y := fp.y + yo
	// these track whether the neighbor is in a different face, and which x/y direction the face is in
	fxdir := 0
	fydir := 0

	if x < 0 {
		fxdir = -1
		x += hp.FaceSidePixels()
	} else if x >= hp.FaceSidePixels() {
		fxdir = 1
		x -= hp.FaceSidePixels()
	}

	if y < 0 {
		fydir = -1
		y += hp.FaceSidePixels()
	} else if y >= hp.FaceSidePixels() {
		fydir = 1
		y -= hp.FaceSidePixels()
	}

	if fxdir == 1 && fydir != -1 && fp.face < 4 {
		x = y
		y = maxXY
	}
	if fydir == 1 && fxdir != -1 && fp.face < 4 {
		y = x
		x = maxXY
	}

	if fydir == -1 && fxdir != 1 && fp.face > 7 {
		y = x
		x = 0
	}
	if fxdir == -1 && fydir != 1 && fp.face > 7 {
		x = y
		y = 0
	}

	face := NewFace(fp.face).Neighbor(fxdir, fydir)
	return FacePixel{x, y, face}.PixelId(hp, scheme)
}

// Given a desired coordinate on a healpix map, return the pixel indices
// of each neighbor pixel of the selected coordinate in the HEALPix index
// scheme desired.
func Neighbors(hp Healpix, where Where, scheme HealpixScheme) []int {
	fp := where.ToFacePixel(hp)
	maxXY := hp.FaceSidePixels() - 1
	var result []int

	if fp.x > 0 && fp.x < maxXY && fp.y > 0 && fp.y < maxXY {
		result = make([]int, 8)
		// pixel not on face boundary
		// highest probability branch in higher resolutions
		// TODO: nested can be even faster here as a special case, doesn't have to go through FacePixel first
		result[0] = FacePixel{fp.x - 1, fp.y - 1, fp.face}.PixelId(hp, scheme)
		result[1] = FacePixel{fp.x, fp.y - 1, fp.face}.PixelId(hp, scheme)
		result[2] = FacePixel{fp.x + 1, fp.y - 1, fp.face}.PixelId(hp, scheme)
		result[3] = FacePixel{fp.x - 1, fp.y, fp.face}.PixelId(hp, scheme)
		result[4] = FacePixel{fp.x + 1, fp.y, fp.face}.PixelId(hp, scheme)
		result[5] = FacePixel{fp.x - 1, fp.y + 1, fp.face}.PixelId(hp, scheme)
		result[6] = FacePixel{fp.x, fp.y + 1, fp.face}.PixelId(hp, scheme)
		result[7] = FacePixel{fp.x + 1, fp.y + 1, fp.face}.PixelId(hp, scheme)
	} else {
		// pixel is on an edge boundary, we need to be cognizant of
		// special corner pixels and edge index swapping

		result = []int{}
		iterLen := 9
		// account for special pixels that only have 7 neighbors
		if fp.face < 8 && fp.x == maxXY && fp.y == maxXY {
			iterLen = 8
		}
		iterStart := 0
		if fp.face > 3 && fp.x == 0 && fp.y == 0 {
			iterStart = 1
		}
		for i := iterStart; i < iterLen; i++ {
			xo := i%3 - 1
			yo := (i / 3) - 1
			// we iterate over the pixel itself but don't include it in the result set
			if xo == 0 && yo == 0 {
				continue
			}

			x := fp.x + xo
			y := fp.y + yo
			// these track whether the neighbor is in a different face, and which x/y direction the face is in
			fxdir := 0
			fydir := 0

			if x < 0 {
				fxdir = -1
				x += hp.FaceSidePixels()
			} else if x >= hp.FaceSidePixels() {
				fxdir = 1
				x -= hp.FaceSidePixels()
			}

			if y < 0 {
				fydir = -1
				y += hp.FaceSidePixels()
			} else if y >= hp.FaceSidePixels() {
				fydir = 1
				y -= hp.FaceSidePixels()
			}

			if fxdir == 1 && fydir != -1 && fp.face < 4 {
				x = y
				y = maxXY
			}
			if fydir == 1 && fxdir != -1 && fp.face < 4 {
				y = x
				x = maxXY
			}

			if fydir == -1 && fxdir != 1 && fp.face > 7 {
				y = x
				x = 0
			}
			if fxdir == -1 && fydir != 1 && fp.face > 7 {
				x = y
				y = 0
			}

			face := NewFace(fp.face).Neighbor(fxdir, fydir)
			neigh := FacePixel{x, y, face}.PixelId(hp, scheme)
			result = append(result, neigh)
		}
	}

	return result
}
