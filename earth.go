package geographic

import "math"

const (
	_R  = 6378137
	_RN = 6356752.3142
	_F  = (_R - _RN) / _R
	_E2 = _F * (2 - _F)
	_PI = 3.14159265359
)

func radians(angle float64) float64 {
	return angle / 180 * math.Pi
}
