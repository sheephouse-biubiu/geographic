package geographic

import (
	"math"
	"strconv"
	"strings"
)

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

func ParseString(str string, sc string, ec string) string {
	s := strings.Index(str, sc)
	if s == -1 {
		s = 0
	} else {
		s = s + len(sc)
	}
	e := strings.Index(str, ec)
	if e == -1 {
		return ""
	}

	return str[s:e]
}

// ParseDegrees 如 "xxx°xx′xx.xxx″"格式数据
func ParseDegrees(str string) (string, string, string) {
	degree := ParseString(str, "\"", "°")
	minutes := ParseString(str, "°", "′")
	second := ParseString(str, "′", "″")
	return degree, minutes, second
}

func DMSToDegrees(degree, minutes, seconds string) float64 {
	d, _ := strconv.ParseFloat(degree, 64)
	m, _ := strconv.ParseFloat(minutes, 64)
	s, _ := strconv.ParseFloat(seconds, 64)
	return d + m/60.0 + s/3600.0
}
