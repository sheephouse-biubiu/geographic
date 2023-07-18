package geographic

import "math"

type Origin struct {
	oLat       float64
	oLon       float64
	oAlt       float64
	oSinLambda float64
	oCosLambda float64
	oSinPhi    float64
	oCosPhi    float64
	oN         float64
}

func (o *Origin) SetOrigin(lat, lon, alt float64) {
	o.oLat = lat
	o.oLon = lon

	lamb := radians(lat)
	phi := radians(lon)

	s := math.Sin(lamb)
	o.oAlt = alt
	o.oN = _R / math.Sqrt(1-_E2*s*s)
	o.oSinLambda = math.Sin(lamb)
	o.oCosLambda = math.Cos(lamb)
	o.oSinPhi = math.Sin(phi)
	o.oCosPhi = math.Cos(phi)
}

func (o *Origin) GetLat() float64 {
	return o.oLat
}

func (o *Origin) GetLon() float64 {
	return o.oLon
}

func (o *Origin) GetAlt() float64 {
	return o.oAlt
}

func (o *Origin) GetN() float64 {
	return o.oN
}

func (o *Origin) GetSinLambda() float64 {
	return o.oSinLambda
}

func (o *Origin) GetCosLambda() float64 {
	return o.oCosLambda
}

func (o *Origin) GetSinPhi() float64 {
	return o.oSinPhi
}

func (o *Origin) GetCosPhi() float64 {
	return o.oCosPhi
}
