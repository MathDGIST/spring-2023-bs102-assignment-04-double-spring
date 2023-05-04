package double

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

type Double struct {
	Val [9]float64               // m1, m2, k1, k2, x0, x1, x2, x3, T
	Sol [2]func(float64) float64 // spring motion solution
	Xb  [2]float64
	Yb  [2][2]float64
	TrX func(float64) int
	TrY [2]func(float64) int
}

const (
	Second    int = 2   // Second number (100th of seconds)
	Height    int = 600 // height (pixel) of image
	GphHeight int = 175
	Width     int = 600 // width (pixel) of image
	SprWidth  int = 200 // spring canvas width (pixel) -- on the left
	GphWidth  int = 400 // graph canvas width (pixel) -- on the right
	Pad       int = 10  // padding size (pixel)
	Ceil      int = 50  // ceiling height (pixel)
	SpL       int = 100 // spring minimum length (pixel)
	Mr        int = 10  // circle radius (pixel)
	CoilN     int = 10  // coil number (int)
	CoilR     int = 10  // coil radius (pixel)
	CoilM     int = 5   // coil margin (pixel)
)

func NewDouble() *Double {
	return &Double{}
}

func (dsp *Double) HandleInputs() {
	msg := []string{
		"====================================\n",
		"  @       This program simulates\n",
		"  @ k1    the double spring motion\n",
		"  @       modeled by two displacement\n",
		"  m1      functions x1(t) and x2(t)\n",
		"  @       over the interval [0, T]\n",
		"  @ k2    Intial values are\n",
		"  @       x1(0)=x0, x1'(0)=x1\n",
		"  m2      x2(0)=x2, x2'(0)=x3\n",
		"=====================================\n",
	}
	for _, m := range msg {
		fmt.Printf("%s", m)
	}
	msg = []string{
		"m1: ",
		"m2: ",
		"k1: ",
		"k2: ",
		"x0: ",
		"x1: ",
		"x2: ",
		"x3: ",
		"T: ",
	}
	for i, m := range msg {
		fmt.Printf("Enter the value of %s", m)
		fmt.Scanf("%f\n", &dsp.Val[i])
	}
}

// FindSolution finds the solutions for the differential equation
// and store the results in the field Sol in Double struct.
// The function of displacement of mass 1 is stored in Sol[0]
// and the function of displacement of mass 2 is stored in Sol[1].
func (dsp *Double) FindSolution() {
	m1, m2, k1, k2 := dsp.Val[0], dsp.Val[1], dsp.Val[2], dsp.Val[3]
	x0, x1, x2, x3 := dsp.Val[4], dsp.Val[5], dsp.Val[6], dsp.Val[7]
	// The differential equation is
	// m1*x1'' = -k1*x1 + k2*(x2-x1)
	// m2*x2'' = -k2*(x2-x1)
	// Laplace transform
	// (m1*s^2 + (k1+k2))*X1 - k2*X2 = m1*x0*s + m1*x1
	// -k2*X1 + (m2*s^2+k2)*X2 = m2*x2*s + m2*x3
	// a^2, b^2 are the (real) roots of t^2 + (k2/m2 + (k1+k2)/m1)*t + k2^2/(m1*m2)
	/* Use this space if necessary ------------


	------------------------------------------- */
	// X1 = A*(s/(s^2+a^2) + B*(1/(s^2+a^2) + C*(s/(s^2+b^2) + D*(1/(s^2+b^2)
	// A + C = x0
	// b^2*A + a^2*C = k2*(x2/m1+x0/m2)
	// B + D = x1
	// b^2*B + a^2*D = k2(x3/m1+x1/m2)
	// X2 = A*(s/(s^2+a^2) + B*(1/(s^2+a^2) + C*(s/(s^2+b^2) + D*(1/(s^2+b^2)
	// A + C = x2
	// b^2*A + a^2*C = (k1+k2)/m1*x2 + k2/m2*x0
	// B + D = x3
	// b^2*B + a^2*D = (k1+k2)/m1*x3 + k2/m2*x1
	/* Use this space if necessary ------------


	------------------------------------------- */

	dsp.Sol[0] = func(t float64) float64 {
		/* Fill here ------------

		------------------------- */
	}

	dsp.Sol[1] = func(t float64) float64 {
		/* Fill here ------------

		------------------------- */
	}
}

func (dsp *Double) Animate(fn string) {
	dsp.findBounds()
	dsp.findTr()
	g := &gif.GIF{
		LoopCount: 1,
	}
	for i := 1; i <= 30*Second; i++ {
		p := []color.Color{
			color.White,
			color.Black,
			color.NRGBA{255, 0, 0, 255},
			color.NRGBA{0, 0, 255, 255}}
		c := image.NewPaletted(image.Rect(0, 0, Width, Height), p)
		dsp.drawAxis(c)
		dsp.drawGraph(c, i)
		dsp.drawSpring(c, i)
		dsp.drawLabel(c)
		g.Image = append(g.Image, c)
		g.Delay = append(g.Delay, 1)
	}
	fp, err := os.Create(fn)
	defer fp.Close()
	if err != nil {
		panic(err)
	}
	err = gif.EncodeAll(fp, g)
	if err != nil {
		panic(err)
	}
}
