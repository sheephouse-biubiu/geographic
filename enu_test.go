package geographic

import (
	"fmt"
	"math"
	"testing"
)

func Wgs84_ENU(w *Wgs84) {

	////w.SetOrigin(0, 0, 0)
	//s := time.Now()
	//for i := 0; i < 100000; i++ {
	//	w.ToENU(29.7466623337786, 106.5538601892096, 239.1483001708984)
	//}
	//e := time.Now()
	//fmt.Println(e.Sub(s))
	//x, y, h := w.ToENU(29.7466623337786, 106.5538601892096, 239.1483001708984)
	w.ToENU(29.7466623337786, 106.5538601892096, 239.1483001708984)

	// fmt.Printf("%.14f, %.14f, %.14f \n", x, y, h)
}

func BenchmarkWgs84_ToENU(b *testing.B) {
	fmt.Println(math.Pow(2, 0.5))
	w := Wgs84{}
	w.SetOrigin(29.746668757502476, 106.55388876468517, 239.23825073242188)
	x, y, z := w.ToENU(29.7466623337786, 106.5538601892096, 239.1483001708984)
	lat, lon, alt := w.FromEnu(x, y, z)
	fmt.Printf("%.15f, %.15f, %.15f\n", lat, lon, alt)
}
