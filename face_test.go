package healpix

import "testing"

func TestFaceRow(t *testing.T) {
	testCases := []struct {
		name   string
		faceId int
		row    int
		southX int
		southY int
	}{
		{"Face 0 on row 0", 0, 0, 1, 2},
		{"Face 1 on row 0", 1, 0, 3, 2},
		{"Face 2 on row 0", 2, 0, 5, 2},
		{"Face 3 on row 0", 3, 0, 7, 2},

		{"Face 4 on row 1", 4, 1, 0, 3},
		{"Face 5 on row 1", 5, 1, 2, 3},
		{"Face 6 on row 1", 6, 1, 4, 3},
		{"Face 7 on row 1", 7, 1, 6, 3},

		{"Face 8 on row 2", 8, 2, 1, 4},
		{"Face 9 on row 2", 9, 2, 3, 4},
		{"Face 10 on row 2", 10, 2, 5, 4},
		{"Face 11 on row 2", 11, 2, 7, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			face := NewFace(tc.faceId)
			if face.FaceRow() != tc.row {
				t.Errorf("Face %v row expected %v, got %v instead", tc.faceId, tc.row, face.FaceRow())
			}
			southX, southY := face.SouthernmostVertex()
			if southX != tc.southX {
				t.Errorf("Face %v south X expected %v, got %v instead", tc.faceId, tc.southX, southX)
			}
			if southY != tc.southY {
				t.Errorf("Face %v south Y expected %v, got %v instead", tc.faceId, tc.southY, southY)
			}
		})
	}
}

func TestFaceNeighbors(t *testing.T) {
	testCases := []struct {
		name     string
		faceId   int
		xn       int
		yn       int
		neighbor int
	}{
		{"Face 0 neighbor -1,-1 is face 8", 0, -1, -1, 8},
		{"Face 0 neighbor -1,0 is face 5", 0, -1, 0, 4},
		{"Face 0 neighbor -1,1 is face 1", 0, -1, 1, 3},
		{"Face 0 neighbor 0,-1 is face 4", 0, 0, -1, 5},
		{"Face 0 neighbor 0,1 is face 1", 0, 0, 1, 3},
		{"Face 0 neighbor 1,-1 is face 3", 0, 1, -1, 1},
		{"Face 0 neighbor 1,0 is face 3", 0, 1, 0, 1},

		{"Face 1 neighbor -1,-1 is face 9", 1, -1, -1, 9},
		{"Face 1 neighbor -1,0 is face 6", 1, -1, 0, 5},
		{"Face 1 neighbor -1,1 is face 2", 1, -1, 1, 0},
		{"Face 1 neighbor 0,-1 is face 5", 1, 0, -1, 6},
		{"Face 1 neighbor 0,1 is face 2", 1, 0, 1, 0},
		{"Face 1 neighbor 1,0 is face 0", 1, 1, 0, 2},
		{"Face 1 neighbor 1,-1 is face 0", 1, 1, -1, 2},

		{"Face 4 neighbor -1,0 is face 8", 4, -1, 0, 11},
		{"Face 4 neighbor -1,1 is face 5", 4, -1, 1, 7},
		{"Face 4 neighbor 0,-1 is face 11", 4, 0, -1, 8},
		{"Face 4 neighbor 0,1 is face 0", 4, 0, 1, 3},
		{"Face 4 neighbor 1,-1 is face 7", 4, 1, -1, 5},
		{"Face 4 neighbor 1,0 is face 3", 4, 1, 0, 0},

		{"Face 7 neighbor -1,0 is face 8", 7, -1, 0, 10},
		{"Face 7 neighbor -1,1 is face 5", 7, -1, 1, 6},
		{"Face 7 neighbor 0,-1 is face 11", 7, 0, -1, 11},
		{"Face 7 neighbor 0,1 is face 0", 7, 0, 1, 2},
		{"Face 7 neighbor 1,-1 is face 7", 7, 1, -1, 4},
		{"Face 7 neighbor 1,0 is face 3", 7, 1, 0, 3},

		{"Face 8 neighbor -1,0 is face 9", 8, -1, 0, 11},
		{"Face 8 neighbor -1,1 is face 9", 8, -1, 1, 11},
		{"Face 8 neighbor 0,-1 is face 11", 8, 0, -1, 9},
		{"Face 8 neighbor 0,1 is face 5", 8, 0, 1, 4},
		{"Face 8 neighbor 1,-1 is face 11", 8, 1, -1, 9},
		{"Face 8 neighbor 1,0 is face 4", 8, 1, 0, 5},
		{"Face 8 neighbor 1,1 is face 0", 8, 1, 1, 0},

		{"Face 11 neighbor -1,0 is face 8", 11, -1, 0, 10},
		{"Face 11 neighbor -1,1 is face 8", 11, -1, 1, 10},
		{"Face 11 neighbor 0,-1 is face 10", 11, 0, -1, 8},
		{"Face 11 neighbor 0,1 is face 4", 11, 0, 1, 7},
		{"Face 11 neighbor 1,-1 is face 10", 11, 1, -1, 8},
		{"Face 11 neighbor 1,0 is face 7", 11, 1, 0, 4},
		{"Face 11 neighbor 1,1 is face 3", 11, 1, 1, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			face := NewFace(tc.faceId)
			neighbor := face.Neighbor(tc.xn, tc.yn)
			if neighbor != tc.neighbor {
				t.Errorf("Face %v expected neighbor %v in direction %v,%v, got %v instead", tc.faceId, tc.neighbor, tc.xn, tc.yn, neighbor)
			}
		})
	}
}
