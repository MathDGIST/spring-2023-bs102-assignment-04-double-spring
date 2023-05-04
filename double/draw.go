package double

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"math"
)

func (dsp *Double) drawGraph(c *image.Paletted, count int) {
	for i := 0; i < count; i++ {
		x0, x1 := dsp.Xb[1]/float64(30*Second)*float64(i), dsp.Xb[1]/float64(30*Second)*float64(i+1)
		for j := 0; j < 2; j++ {
			y0, y1 := dsp.Sol[j](x0), dsp.Sol[j](x1)
			p0 := [2]int{dsp.TrX(x0), dsp.TrY[j](y0)}
			p1 := [2]int{dsp.TrX(x1), dsp.TrY[j](y1)}
			drawLine(c, p0, p1, j+2)
		}
	}
}

func (dsp *Double) drawAxis(c *image.Paletted) {
	for j := 0; j < 2; j++ {
		for i := dsp.TrY[j](dsp.Yb[j][1]); i <= dsp.TrY[j](dsp.Yb[j][0]); i++ {
			c.Set(dsp.TrX(0), i, c.Palette[1])
		}
		for i := dsp.TrX(dsp.Xb[0]); i <= dsp.TrX(dsp.Xb[1]); i++ {
			c.Set(i, dsp.TrY[j](0), c.Palette[1])
		}

	}
}

func (dsp *Double) drawSpring(c *image.Paletted, count int) {
	// Draw ceiling
	for i := Pad; i < SprWidth-Pad; i++ {
		c.Set(i, Ceil, c.Palette[1])
	}
	// Draw spring
	for i := 0; i < 100; i++ {
		for j := 0; j < 2; j++ {
			p0 := dsp.springXY(j, i, count)
			p1 := dsp.springXY(j, i+1, count)
			drawLine(c, p0, p1, 1)
		}
	}
	// Draw mass
	for j := 0; j < 2; j++ {
		y := dsp.Sol[j](dsp.Xb[1] / float64(30*Second) * float64(count))
		for i := 0; i < 100; i++ {
			p0 := dsp.circleXY(dsp.TrY[j](y), i)
			p1 := dsp.circleXY(dsp.TrY[j](y), i+1)
			drawLine(c, p0, p1, 1)
		}
	}
}

func (dsp *Double) springXY(j, i, k int) [2]int {
	n := CoilN
	r := CoilR
	m := CoilM
	y0 := dsp.Sol[0](dsp.Xb[1] / float64(30*Second) * float64(k))
	y1 := dsp.Sol[1](dsp.Xb[1] / float64(30*Second) * float64(k))
	st := -math.Pi / 2
	xi := int(float64(r) * math.Cos(st+(2*math.Pi*float64(n)-2*st)*float64(i)/100))
	xi += Pad + SprWidth/2
	var yi int
	switch j {
	case 0:
		l := dsp.TrY[0](y0) - Ceil - Mr - 2*r - 2*m
		yi = int(float64(r) * math.Sin(st+(2*math.Pi*float64(n)-2*st)*float64(i)/100))
		yi += int(float64(l) / float64(100) * float64(i))
		yi += m + r + Ceil
	case 1:
		l := dsp.TrY[1](y1) - dsp.TrY[0](y0) - 2*Mr - 2*r - 2*m
		yi = int(float64(r) * math.Sin(st+(2*math.Pi*float64(n)-2*st)*float64(i)/100))
		yi += dsp.TrY[0](y0) + Mr + m + r + int(float64(l)/float64(100)*float64(i))
	}
	return [2]int{xi, yi}
}

func (dsp *Double) circleXY(l, i int) [2]int {
	r := Mr
	xi := int(float64(r) * math.Cos(2*math.Pi*float64(i)/100))
	xi += SprWidth/2 + Pad
	yi := int(float64(r) * math.Sin(2*math.Pi*float64(i)/100))
	yi += l
	return [2]int{xi, yi}
}

func drawLine(c *image.Paletted, p0, p1 [2]int, color int) {
	x, y := 0, 0
	n := findMax(p0, p1)
	for i := 0; i <= n; i++ {
		nx := p0[0] + int(float64((p1[0]-p0[0])*i)/float64(n))
		ny := p0[1] + int(float64((p1[1]-p0[1])*i)/float64(n))
		for nx == x && ny == y {
			continue
		}
		x, y := nx, ny
		c.Set(x, y, c.Palette[color])
	}
}

func (dsp *Double) drawLabel(c *image.Paletted) {
	d := &font.Drawer{
		Dst:  c,
		Src:  image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: basicfont.Face7x13,
	}
	label := "Double Spring Simulator v1.0 "
	d.Dot = fixed.Point26_6{fixed.I(30), fixed.I(30)}
	d.DrawString(label)

	label = fmt.Sprintf("m1 = %.1f(kg), m2 = %.1f(kg), k1 = %.1f(N/m), k2 = %.1f", dsp.Val[0], dsp.Val[1], dsp.Val[2], dsp.Val[3])
	d.Dot = fixed.Point26_6{fixed.I(SprWidth + 30), fixed.I(60)}
	d.DrawString(label)

	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			label = fmt.Sprintf("%.1f", dsp.Yb[j][i])
			d.Dot = fixed.Point26_6{fixed.I(SprWidth + 20), fixed.I(dsp.TrY[j](dsp.Yb[j][i]))}
			d.DrawString(label)
		}
		label = fmt.Sprintf("%.1f", dsp.Xb[1])
		d.Dot = fixed.Point26_6{fixed.I(Width - 50), fixed.I(dsp.TrY[j](0) - 5)}
		d.DrawString(label)
	}
}
