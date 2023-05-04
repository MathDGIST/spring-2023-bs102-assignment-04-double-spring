package double

import "math"

func findMax(p0, p1 [2]int) int {
	mx := absInt(p0[0] - p1[0])
	my := absInt(p0[1] - p1[1])
	if mx >= my {
		return mx
	}
	return my
}

func (dsp *Double) findBounds() {
	dsp.Xb = [2]float64{0, dsp.Val[8]}
	for k := 0; k < 2; k++ {
		ymin, ymax := math.MaxFloat64, -math.MaxFloat64
		for i := 0; i <= 30*Second; i++ {
			x := (dsp.Xb[1] - dsp.Xb[0]) / float64(30*Second) * float64(i)
			y := -dsp.Sol[k](x)
			if y < ymin {
				ymin = y
			}
			if y > ymax {
				ymax = y
			}
		}
		if ymin > 0 {
			ymin = 0
		}
		if ymax < 0 {
			ymax = 0
		}
		dsp.Yb[k] = [2]float64{ymin, ymax}
	}
}

func (dsp *Double) findTr() {
	dsp.TrX = func(x float64) int {
		return Pad + SprWidth + int(float64(GphWidth-2*Pad)*(x-dsp.Xb[0])/(dsp.Xb[1]-dsp.Xb[0]))
	}
	dsp.TrY[0] = func(y float64) int {
		return Pad + Ceil + SpL + int(float64(GphHeight-2*Pad)*(dsp.Yb[0][1]-y)/(dsp.Yb[0][1]-dsp.Yb[0][0]))
	}
	dsp.TrY[1] = func(y float64) int {
		return Pad + Ceil + 2*SpL + GphHeight + int(float64(GphHeight-2*Pad)*(dsp.Yb[1][1]-y)/(dsp.Yb[1][1]-dsp.Yb[1][0]))
	}
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
