package healpix

import "math"

// The maximum possible order of a healpix map supported on the currently
// executing machine. Significantly smaller on 32-bit machines than on
// 64-bit machines.
func MaxOrder() int {
	return int(math.Floor(0.5 * math.Log2(math.MaxInt/12)))
}

// The maximum possible number of subpixels on the side of a base pixel
// of a healpix map on the currently executing machine. Significantly
// smaller on 32-bit machines than on 64-bit machines.
func MaxNSide() int {
	return MaxOrder() * MaxOrder()
}

// Check whether the given number is a valid order value to create a healpix map.
func IsValidOrder(test int) bool {
	return test >= 0 && test <= MaxOrder()
}

// Check whether the given number is a valid NSide value for a healpix map.
func IsValidNSide(test int) bool {
	return test > 0 && (test&(test-1) == 0)
}
