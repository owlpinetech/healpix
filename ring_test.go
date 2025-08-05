package healpix

import (
	"math"
	"testing"
)

func TestFirstIndex(t *testing.T) {
	testCases := []struct {
		name     string
		order    HealpixOrder
		ringId   int
		fstIndex uint
	}{
		{"Ring 0 @ order 0 = 0 index", 0, 0, 0},
		{"Ring 1 @ order 0 = 4 index", 0, 1, 4},
		{"Ring 2 @ order 0 = 8 index", 0, 2, 8},

		{"Ring 0 @ order 1 = 0 index", 1, 0, 0},
		{"Ring 1 @ order 1 = 4 index", 1, 1, 4},
		{"Ring 2 @ order 1 = 12 index", 1, 2, 12},
		{"Ring 3 @ order 1 = 20 index", 1, 3, 20},
		{"Ring 4 @ order 1 = 28 index", 1, 4, 28},
		{"Ring 5 @ order 1 = 36 index", 1, 5, 36},
		{"Ring 6 @ order 1 = 44 index", 1, 6, 44},

		{"Ring 0 @ order 2 = 0 index", 2, 0, 0},
		{"Ring 1 @ order 2 = 4 index", 2, 1, 4},
		{"Ring 2 @ order 2 = 8 index", 2, 2, 12},
		{"Ring 3 @ order 2 = 24 index", 2, 3, 24},
		{"Ring 4 @ order 2 = 40 index", 2, 4, 40},
		{"Ring 10 @ order 2 = 136 index", 2, 10, 136},
		{"Ring 11 @ order 2 = 152 index", 2, 11, 152},
		{"Ring 12 @ order 2 = 168 index", 2, 12, 168},
		{"Ring 13 @ order 2 = 180 index", 2, 13, 180},
		{"Ring 14 @ order 2 = 188 index", 2, 14, 188},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ring := NewRing(New(tc.order), tc.ringId)
			if ring.FirstIndex() != tc.fstIndex {
				t.Errorf("Ring %v @ order %v first index expected %v, got %v instead", tc.ringId, tc.order, tc.fstIndex, ring.FirstIndex())
			}
		})
	}
}

func TestRingOffset(t *testing.T) {
	testCases := []struct {
		name   string
		order  HealpixOrder
		ringId int
		offset bool
	}{
		{"Ring 0 @ order 0 = true", 0, 0, true},
		{"Ring 1 @ order 0 = false", 0, 1, false},
		{"Ring 2 @ order 0 = true", 0, 2, true},

		{"Ring 0 @ order 1 = true", 1, 0, true},
		{"Ring 1 @ order 1 = true", 1, 1, true},
		{"Ring 2 @ order 1 = false", 1, 2, false},
		{"Ring 3 @ order 1 = true", 1, 3, true},
		{"Ring 4 @ order 1 = false", 1, 4, false},
		{"Ring 5 @ order 1 = true", 1, 5, true},
		{"Ring 6 @ order 1 = true", 1, 6, true},

		{"Ring 0 @ order 2 = true", 2, 0, true},
		{"Ring 1 @ order 2 = true", 2, 1, true},
		{"Ring 2 @ order 2 = true", 2, 2, true},
		{"Ring 3 @ order 2 = true", 2, 3, true},
		{"Ring 4 @ order 2 = false", 2, 4, false},
		{"Ring 10 @ order 2 = false", 2, 10, false},
		{"Ring 11 @ order 2 = true", 2, 11, true},
		{"Ring 12 @ order 2 = true", 2, 12, true},
		{"Ring 13 @ order 2 = true", 2, 13, true},
		{"Ring 14 @ order 2 = true", 2, 14, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ring := NewRing(New(tc.order), tc.ringId)
			if ring.IsOffset() != tc.offset {
				t.Errorf("Ring %v @ order %v offset expected %v, got %v instead", tc.ringId, tc.order, tc.offset, ring.IsOffset())
			}
		})
	}
}

func TestRing0AndNMinus1PixelsAlways4(t *testing.T) {
	for order := HealpixOrder(0); order <= HealpixOrder(MaxOrder()); order++ {
		hp := New(order)
		ring := NewRing(hp, 0)
		if ring.Pixels() != 4 {
			t.Errorf("Ring 0 @ order %v pixels expected 4, got %v instead", order, ring.Pixels())
		}
		ring = NewRing(New(order), hp.Rings()-1)
		if ring.Pixels() != 4 {
			t.Errorf("Ring N-1 @ order %v pixels expected 4, got %v instead", order, ring.Pixels())
		}
	}
}

func TestRingPixels(t *testing.T) {
	testCases := []struct {
		name   string
		order  HealpixOrder
		ringId int
		pixels int
	}{
		{"Ring 0 @ order 0 = 4 pixels", 0, 0, 4},
		{"Ring 1 @ order 0 = 4 pixels", 0, 1, 4},
		{"Ring 2 @ order 0 = 4 pixels", 0, 2, 4},

		{"Ring 0 @ order 1 = 4 pixels", 1, 0, 4},
		{"Ring 1 @ order 1 = 8 pixels", 1, 1, 8},
		{"Ring 2 @ order 1 = 8 pixels", 1, 2, 8},
		{"Ring 3 @ order 1 = 8 pixels", 1, 3, 8},
		{"Ring 4 @ order 1 = 8 pixels", 1, 4, 8},
		{"Ring 5 @ order 1 = 8 pixels", 1, 5, 8},
		{"Ring 6 @ order 1 = 4 pixels", 1, 6, 4},

		{"Ring 0 @ order 2 = 0 pixels", 2, 0, 4},
		{"Ring 1 @ order 2 = 4 pixels", 2, 1, 8},
		{"Ring 2 @ order 2 = 8 pixels", 2, 2, 12},
		{"Ring 3 @ order 2 = 24 pixels", 2, 3, 16},
		{"Ring 4 @ order 2 = 40 pixels", 2, 4, 16},
		{"Ring 10 @ order 2 = 136 pixels", 2, 10, 16},
		{"Ring 11 @ order 2 = 152 pixels", 2, 11, 16},
		{"Ring 12 @ order 2 = 168 pixels", 2, 12, 12},
		{"Ring 13 @ order 2 = 180 pixels", 2, 13, 8},
		{"Ring 14 @ order 2 = 188 pixels", 2, 14, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ring := NewRing(New(tc.order), tc.ringId)
			if ring.Pixels() != tc.pixels {
				t.Errorf("Ring %v @ order %v pixels expected %v, got %v instead", tc.ringId, tc.order, tc.pixels, ring.Pixels())
			}
		})
	}
}

func withinTolerance(n1, n2, tolerance float64) bool {
	if n1 == n2 {
		return true
	}
	diff := math.Abs(n1 - n2)
	if n2 == 0 {
		return diff < tolerance
	}
	return (diff / math.Abs(n2)) < tolerance
}

func TestRingLatitude(t *testing.T) {
	testCases := []struct {
		name     string
		order    HealpixOrder
		ringId   int
		latitude float64
	}{
		{"Ring 0 @ order 0 = Pi/4", 0, 0, math.Pi/2 - math.Acos(float64(2)/float64(3))},
		{"Ring 1 @ order 0 = 0", 0, 1, 0},
		{"Ring 2 @ order 0 = -Pi/4", 0, 2, math.Pi/2 - math.Acos(float64(-2)/float64(3))},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ring := NewRing(New(tc.order), tc.ringId)
			if !withinTolerance(ring.Latitude(), tc.latitude, 1e-10) {
				t.Errorf("Ring %v @ order %v latitude expected %v, got %v instead", tc.ringId, tc.order, tc.latitude, ring.Latitude())
			}
		})
	}
}
