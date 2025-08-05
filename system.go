package healpix

import (
	"math"
	"math/bits"
)

// The maximum possible order of a healpix map supported on the currently
// executing machine. Significantly smaller on 32-bit machines than on
// 64-bit machines.
func MaxOrder() int {
	// this is basically the inverse of computing the maximum number of pixels given an nside value
	maxPixels := uint(math.MaxUint)
	maxPixelsPerFace := maxPixels / uint(BasePixelsPerRow*BasePixelRows)
	pixelsPerSide := math.Sqrt(float64(maxPixelsPerFace))
	order := bits.Len(uint(pixelsPerSide)) - 1 // floor(log2()) equivalent for integers
	return order - 1                           // compensates for Unique Pixels taking up more numeric space than the base pixels
}

// The maximum possible number of subpixels on the side of a base pixel
// of a healpix map on the currently executing machine. Significantly
// smaller on 32-bit machines than on 64-bit machines.
func MaxNSide() int {
	return 1 << MaxOrder() // find the largest power of 2 that is less than or equal to pixelsPerSide
}

// Check whether the given number is a valid order value to create a healpix map.
func IsValidOrder(test int) bool {
	return test >= 0 && test <= MaxOrder()
}

// Check whether the given number is a valid NSide value for a healpix map.
func IsValidNSide(test int) bool {
	// middle test checks if test is a power of 2
	return test > 0 && (test&(test-1) == 0) && test <= MaxNSide()
}
