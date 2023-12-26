package geographic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDegree_DegreeZero(t *testing.T) {
	degree, minutes, seconds := ParseDegrees("12′23.1243″")
	as := assert.New(t)
	as.Equal(degree, "")
	as.Equal(minutes, "12")
	as.Equal(seconds, "23.1243")
}

func TestParseDegrees_MinutesZero(t *testing.T) {
	degree, minutes, seconds := ParseDegrees("20°′23.1243″")
	as := assert.New(t)
	as.Equal(degree, "20")
	as.Equal(minutes, "")
	as.Equal(seconds, "23.1243")
}

func TestParseDegrees_SecondZero(t *testing.T) {
	degree, minutes, seconds := ParseDegrees("20°30′″")
	as := assert.New(t)
	as.Equal(degree, "20")
	as.Equal(minutes, "30")
	as.Equal(seconds, "")
}
