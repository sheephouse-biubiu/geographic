package geographic

import (
	//"fmt"
	"math"
)

type ECEF struct {
}

func (e *ECEF) toENU(o *Origin, x, y, z float64) (east, north, up float64) {
	sinLambda := o.GetSinLambda()
	cosLambda := o.GetCosLambda()
	sinPhi := o.GetSinPhi()
	cosPhi := o.GetCosPhi()

	alt := o.GetAlt()
	N := o.GetN()

	x0 := (alt + N) * cosLambda * cosPhi
	y0 := (alt + N) * cosLambda * sinPhi
	z0 := (alt + (1-_E2)*N) * sinLambda

	Δx := x - x0
	Δy := y - y0
	Δz := z - z0

	t := -cosPhi*Δx - sinPhi*Δy

	xEast := -sinPhi*Δx + cosPhi*Δy
	yNorth := t*sinLambda + cosLambda*Δz
	zUp := cosLambda*cosPhi*Δx + cosLambda*sinPhi*Δy + sinLambda*Δz
	return xEast, yNorth, zUp
}

func (e *ECEF) ToWgs84(x, y, z float64) (lat0, lon0, h0 float64) {
	x2 := math.Pow(x, 2)
	y2 := math.Pow(y, 2)
	z2 := math.Pow(z, 2)

	e1 := math.Sqrt(1 - math.Pow(_RN/_R, 2))
	b2 := _RN * _RN
	e2 := math.Pow(e1, 2)
	ep := e1 * (_R / _RN)
	r := math.Sqrt(x2 + y2)
	r2 := r * r
	E2 := math.Pow(_R, 2) - math.Pow(_RN, 2)
	F := 54 * b2 * z2
	G := r2 + (1-e2)*z2 - e2*E2
	c := (e2 * e2 * F * r2) / (G * G * G)
	//fmt.Printf("%0.15f\n", 1+c+math.Sqrt(c*c+2*c))
	s := math.Pow(1+c+math.Sqrt(c*c+2*c), 1.0/3.0)

	P := F / (3 * math.Pow((s+1/s+1), 2) * G * G)
	Q := math.Sqrt(1 + 2*e2*e2*P)
	ro := -(P*e2*r)/(1+Q) + math.Sqrt((_R*_R/2)*(1+1/Q)-(P*(1-e2)*z2)/(Q*(1+Q))-P*r2/2)
	tmp := math.Pow(r-e2*ro, 2)
	U := math.Sqrt(tmp + z2)
	V := math.Sqrt(tmp + (1-e2)*z2)
	zo := (b2 * z) / (_R * V)

	height := U * (1 - b2/(_R*V))
	lat := math.Atan((z + ep*ep*zo) / r)
	temp := math.Atan(y / x)
	var long float64
	if x >= 0 {
		long = temp
	} else if x < 0 && y >= 0 {
		long = math.Pi + temp
	} else {
		long = temp - math.Pi
	}
	lat0 = lat / (math.Pi / 180)
	lon0 = long / (math.Pi / 180)
	h0 = height
	return lat0, lon0, h0
}
