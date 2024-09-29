package shapes

func Perimeter(rectangle Rectangle) (p float64) {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) (a float64) {
	return rectangle.Height * rectangle.Width
}
