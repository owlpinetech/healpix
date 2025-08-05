package healpix

import (
	"testing"
)

func TestHealpixOrderNSide(t *testing.T) {
	testCases := []struct {
		name  string
		order int
		nside int
	}{
		{"0 order = 1 nside", 0, 1},
		{"1 order = 2 nside", 1, 2},
		{"2 order = 4 nside", 2, 4},
		{"3 order = 8 nside", 3, 8},
		{"4 order = 16 nside", 4, 16},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hpo := NewHealpixOrder(tc.order)
			hpn := NewHealpixSide(tc.nside)
			if hpo.Order() != hpn.Order() {
				t.Errorf("Side config expected order %v, got %v instead", hpo.Order(), hpn.Order())
			}
			if hpn.FaceSidePixels() != hpo.FaceSidePixels() {
				t.Errorf("Order config expected side pixels %v, got %v instead", hpn.FaceSidePixels(), hpo.FaceSidePixels())
			}
		})
	}
}

func TestHealpixPixels(t *testing.T) {
	testCases := []struct {
		name   string
		order  int
		nside  int
		pixels uint
	}{
		{"0 order = 1 nside = 12 pixels", 0, 1, 12},
		{"1 order = 2 nside = 48 pixels", 1, 2, 48},
		{"2 order = 4 nside = 192 pixels", 2, 4, 192},
		{"3 order = 8 nside = 768 pixels", 3, 8, 768},
		{"4 order = 16 nside = 3072 pixels", 4, 16, 3072},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hpo := New(NewHealpixOrder(tc.order))
			hpn := New(NewHealpixSide(tc.nside))
			if hpo.Pixels() != tc.pixels {
				t.Errorf("Order config expected pixels %v, got %v instead", tc.pixels, hpo.Pixels())
			}
			if hpn.Pixels() != tc.pixels {
				t.Errorf("Side config expected pixels %v, got %v instead", tc.pixels, hpn.Pixels())
			}
		})
	}
}

func TestHealpixRings(t *testing.T) {
	testCases := []struct {
		name  string
		order int
		nside int
		rings int
	}{
		{"0 order = 1 nside = 3 rings", 0, 1, 3},
		{"1 order = 2 nside = 7 rings", 1, 2, 7},
		{"2 order = 4 nside = 15 rings", 2, 4, 15},
		{"3 order = 8 nside = 31 rings", 3, 8, 31},
		{"4 order = 16 nside = 63 rings", 4, 16, 63},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hpo := New(NewHealpixOrder(tc.order))
			hpn := New(NewHealpixSide(tc.nside))
			if hpo.Rings() != tc.rings {
				t.Errorf("Order config expected rings %v, got %v instead", tc.rings, hpo.Rings())
			}
			if hpn.Rings() != tc.rings {
				t.Errorf("Side config expected rings %v, got %v instead", tc.rings, hpn.Rings())
			}
		})
	}
}

func TestHealpixPolarRegionPixels(t *testing.T) {
	testCases := []struct {
		name  string
		order int
		nside int
		polar int
	}{
		{"0 order = 1 nside = 0 polar", 0, 1, 0},
		{"1 order = 2 nside = 4 polar", 1, 2, 4},
		{"2 order = 4 nside = 24 polar", 2, 4, 24},
		{"3 order = 8 nside = 112 polar", 3, 8, 112},
		{"4 order = 16 nside = 480 polar", 4, 16, 480},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hpo := New(NewHealpixOrder(tc.order))
			hpn := New(NewHealpixSide(tc.nside))
			if hpo.PolarRegionPixels() != tc.polar {
				t.Errorf("Order config expected polar pixels %v, got %v instead", tc.polar, hpo.PolarRegionPixels())
			}
			if hpn.PolarRegionPixels() != tc.polar {
				t.Errorf("Side config expected polar pixels %v, got %v instead", tc.polar, hpn.PolarRegionPixels())
			}
		})
	}
}
