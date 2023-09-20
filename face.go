package healpix

import (
	"fmt"
)

// Access the properties of a base pixel, or 'face', of a HEALPix map.
type Face struct {
	faceId       int          // the index of the face from 0 to 11, proceeding in rings around the HEALPix map
	row          int          // the row of base pixels to which the face belongs (one of 0, 1, 2)
	southVertexY int          // the y coordinate of the southernmost vertex (one of 2, 3, 4)
	southVertexX int          // the x coordinate of the southernmost vertex (one of 0 - 7)
	neighbors    map[byte]int // the neighboring faces of this face. index is x/y directional offset packed into a byte.
}

var faces []Face

// Precompute all the stats and properties for each face.
func init() {
	// the face neighbors of each face
	neighbors := [][]int{
		{8, 4, 3, 5, 0, 3, 1, 1},
		{9, 5, 0, 6, 1, 0, 2, 2},
		{10, 6, 1, 7, 2, 1, 3, 3},
		{11, 7, 2, 8, 3, 2, 0, 0},

		{11, 7, 8, 4, 3, 5, 0},
		{8, 4, 9, 5, 0, 6, 1},
		{9, 5, 10, 6, 1, 7, 2},
		{10, 6, 11, 7, 2, 4, 3},

		{11, 11, 9, 8, 4, 9, 5, 0},
		{8, 8, 10, 9, 5, 10, 6, 1},
		{9, 9, 11, 10, 6, 11, 7, 2},
		{10, 10, 8, 11, 7, 8, 4, 3},
	}

	faces = make([]Face, 12)
	for i := 0; i < 12; i++ {
		row := i / 4
		ind := i % 4
		y := row + 2
		x := 2*ind - (row % 2) + 1
		faces[i] = Face{i, row, y, x, make(map[byte]int, 6)}

		nind := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if row < 2 && x == 1 && y == 1 {
					continue
				}
				if row > 0 && x == -1 && y == -1 {
					continue
				}

				xpack := byte(x + 1)
				ypack := byte(y+1) << 2
				n := xpack | ypack
				// make neighbors easier to index by x/y 'direction'
				faces[i].neighbors[n] = neighbors[i][nind]
				nind += 1
			}
		}
	}
}

// Return the face at the given index, from 0 - 11.
func NewFace(faceId int) Face {
	return faces[faceId]
}

// Returns the index of the face.
func (f Face) FaceId() int {
	return f.faceId
}

// Returns the row index, starting from 0, in which the face resides.
func (f Face) FaceRow() int {
	return f.row
}

// Return the x and y coordinate of the southernmost vertex of this face in abstract face division coordinates.
// x ranges from 0 - 7, y ranges from 0 - 4
func (f Face) SouthernmostVertex() (int, int) {
	return f.southVertexX, f.southVertexY
}

// Return the face id of the neighbor of the current face in the given direction. Each of x and y expects one of
// three values: -1, 0, or 1.
// In this scheme, x and y have 0 at the bottom most vertex of the face (which are arrayed as diamonds conceptually),
// and each increases outward along the diamond boundary, x along the left side and y along the right side.
// So x,y == -1,-1 is the directly southern neighbor; x,y == 1,1 is the directly northern neighbor, and x,y == -1,1
// is the northwestern neighbor.
func (f Face) Neighbor(xOffset int, yOffset int) int {
	xpack := byte(xOffset + 1)
	ypack := byte(yOffset+1) << 2
	ind := xpack | ypack
	if n, ok := f.neighbors[byte(ind)]; ok {
		return n
	}
	panic(fmt.Sprintf("healpix: tried to get a neighbor of a face that has no neighbor in direction %v,%v", xOffset, yOffset))
}
