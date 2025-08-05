package healpix

import "testing"

func TestMaxOrderMaxNSideSamePixelMax(t *testing.T) {
	maxOrder := New(NewHealpixOrder(MaxOrder()))
	maxNSide := New(NewHealpixSide(MaxNSide()))
	if maxOrder.Pixels() != maxNSide.Pixels() {
		t.Errorf("Max order pixels %v should match max nside pixels %v", maxOrder.Pixels(), maxNSide.Pixels())
	}
}

func TestIsValidNSide(t *testing.T) {
	cases := []struct {
		nside int
		valid bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{8, true},
		{16, true},
		{31, false},
		{32, true},
		{64, true},
		{MaxNSide() - 1, false},
		{MaxNSide(), true},
		{MaxNSide() + 1, false},
	}

	for _, c := range cases {
		if IsValidNSide(c.nside) != c.valid {
			t.Errorf("IsValidNSide(%d) = %v, want %v", c.nside, IsValidNSide(c.nside), c.valid)
		}
	}
}
