package geographic

import "math"

type Wgs84 struct {
	o *Origin
}

func (w *Wgs84) SetOrigin(lat, lon, alt float64) {
	w.o = new(Origin)
	w.o.SetOrigin(lat, lon, alt)
}

func (w *Wgs84) ToECEF(lat, lon, alt float64) (x, y, z float64) {
	lamb := radians(lat)
	phi := radians(lon)
	s := math.Sin(lamb)

	N := _R / math.Sqrt(1-_E2*s*s)

	sinLambda := math.Sin(lamb)
	cosLambda := math.Cos(lamb)
	sinPhi := math.Sin(phi)
	cosPhi := math.Cos(phi)

	x = (alt + N) * cosLambda * cosPhi
	y = (alt + N) * cosLambda * sinPhi
	z = (alt + (1-_E2)*N) * sinLambda
	return x, y, z
}

func (w *Wgs84) ToENU(lat, lon, h float64) (east, north, up float64) {
	e := Ecef{}
	x, y, h := w.ToECEF(lat, lon, h)
	return e.toENU(w.o, x, y, h)
}
