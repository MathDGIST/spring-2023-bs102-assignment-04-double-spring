package main

import "main/double"

func main() {
	dsp := double.NewDouble()
	dsp.HandleInputs()
	dsp.FindSolution()
	dsp.Animate("double.gif")
}
