package main

import "fmt"

type Geometry2Interface interface {
	CalculateArea(length int, width int) (area int)
}

type Geometry2 struct{}

func (g2 *Geometry2) CalculateArea(length int, width int) (area int) {
	return length * width
}

type Geometry3 struct{
	G2 Geometry2Interface
}

func (g3 *Geometry3) CalculateVolume(length int, width int, height int) (volume int) {
	return g3.G2.CalculateArea(length, width) * height
}

func main() {
	// initiate struct Geometry2
	g2 := Geometry2{}
	// initiate struct Geometry3
	//with g2 as Geometry2 depedency component
	g3 := Geometry3{
		G2: &g2,
	}

	v := g3.CalculateVolume(3, 5, 10)
	fmt.Println(v)
}