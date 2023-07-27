package geographic

type ENU struct {
}

func (e *ENU) ToECEF(o *Origin, east, north, up float64) (x, y, z float64) {
	sinLambda := o.GetSinLambda()
	cosLambda := o.GetCosLambda()
	sinPhi := o.GetSinPhi()
	cosPhi := o.GetCosPhi()
	alt := o.GetAlt()
	N := o.GetN()

	x0 := (alt + N) * cosLambda * cosPhi
	y0 := (alt + N) * cosLambda * sinPhi
	z0 := (alt + (1-_E2)*N) * sinLambda

	t := cosLambda*up - sinLambda*north

	zd := sinLambda*up + cosLambda*north
	xd := cosPhi*t - sinPhi*east
	yd := sinPhi*t + cosPhi*east

	x = xd + x0
	y = yd + y0
	z = zd + z0
	return x, y, z
}
