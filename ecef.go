package geographic

type Ecef struct {
}

func (e *Ecef) toENU(o *Origin, x, y, z float64) (east, north, up float64) {
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
